package drugtrace

import (
	"sync"
)

// JSONObject 结构体
type JSONObject struct {
	// 文件下载地址。下载文件的字段顺序的含义为：药品ID，平台药品类型，生产厂，通用名，剂型，制剂规格，包装规格，批准文号，商品名，类别码，包装比例，包装级别，子类编码
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolJSONObject = sync.Pool{
	New: func() any {
		return new(JSONObject)
	},
}

// GetJSONObject() 从对象池中获取JSONObject
func GetJSONObject() *JSONObject {
	return poolJSONObject.Get().(*JSONObject)
}

// ReleaseJSONObject 释放JSONObject
func ReleaseJSONObject(v *JSONObject) {
	v.Url = ""
	poolJSONObject.Put(v)
}
