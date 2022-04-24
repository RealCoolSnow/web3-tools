package datacopy

import (
	"encoding/json"
	"reflect"
)

// DataCopy 本来想使用reflect包来完成这个功能，json.Marshal()其实也是借助reflect包来实现的.
func DataCopy(data, res interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, res)
}

// DataReflact 传入*reflect.Value 且返回 *reflect.Value ，将类型交给调用者来决定。
func DataReflact(src, dst *reflect.Value) *reflect.Value {
	for i := 0; i < src.NumField(); i++ {
		filed := src.Type().Field(i).Name
		cfiled := dst.FieldByName(filed)
		if cfiled.IsValid() && cfiled.CanSet() { // 可用的属性且可 赋值
			dst.Field(i).Set(src.Field(i))
		}
	}
	return dst
}
