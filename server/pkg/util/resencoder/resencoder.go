package resencoder

import (
	"encoding/json"
	ht "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
)

func ResponeJsonDeco() http.EncodeResponseFunc {
	return func(w ht.ResponseWriter, r *ht.Request, v interface{}) error {
		// codec, _ := http.CodecForRequest(r, "Accept")
		// data, err := codec.Marshal(v)
		data, err := json.Marshal(v) // 指定json 序列化方式
		if err != nil {
			return err
		}

		_, err = w.Write(data)
		if err != nil {
			return err
		}

		return nil
	}
}
