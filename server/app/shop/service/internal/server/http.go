package server

import (
	"context"
	"fmt"
	v1 "web3/api/shop/service/v1"
	"web3/app/shop/service/internal/conf"
	"web3/app/shop/service/internal/service"
	"web3/pkg/util/resencoder"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/gorilla/handlers"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, tp *tracesdk.TracerProvider, s *service.ShopService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			selector.Server(
				recovery.Recovery(),
				tracing.Server(tracing.WithTracerProvider(tp)),
				logging.Server(logger),
				AuthMiddleware,
			).Path("/api.shop.service.v1.Shop/GetUser").Build(),
		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	// 服务内跨域处理
	// TODO: 引入网关后在网关处理跨域时，需要删除以下处理跨域的代码
	opts = append(opts, http.Filter(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"DNT", "X-Mx-ReqToken", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Authorization", "udid", "appkey", "version", "authenticated", "cookie", "token"}),
		handlers.ExposedHeaders([]string{"DNT", "X-Mx-ReqToken", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Authorization", "udid", "appkey", "version", "authenticated", "cookie", "token"}),
		handlers.OptionStatusCode(204),
	)))

	// 指定为json编码格式
	opts = append(opts, http.ResponseEncoder(resencoder.ResponeJsonDeco()))
	srv := http.NewServer(opts...)
	v1.RegisterShopHTTPServer(srv, s)

	return srv
}

// AuthMiddleware 网关服务会将userid添加到查询参数打到本服务。次中间件将userid 添加到上下文中
func AuthMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		reply, err = handler(ctx, req)
		if tr, ok := transport.FromServerContext(ctx); ok {
			ht, _ := tr.(*http.Transport)
			if len(ht.Request().URL.Query()["userid"]) > 0 {
				fmt.Println("登陆uid:", ht.Request().URL.Query()["userid"][0])
				etxs := context.WithValue(ctx, "userid", ht.Request().URL.Query()["userid"][0])
				reply, err = handler(etxs, req)
			}
		}
		return
	}
}
