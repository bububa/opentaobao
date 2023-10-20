package icbu

import (
	"sync"
)

// ProductTopPublishRequest 结构体
type ProductTopPublishRequest struct {
	// 返回文案的语种，支持en_US,zh,zh_TW
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 商品的具体数据信息
	Xml string `json:"xml,omitempty" xml:"xml,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 商品明文id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolProductTopPublishRequest = sync.Pool{
	New: func() any {
		return new(ProductTopPublishRequest)
	},
}

// GetProductTopPublishRequest() 从对象池中获取ProductTopPublishRequest
func GetProductTopPublishRequest() *ProductTopPublishRequest {
	return poolProductTopPublishRequest.Get().(*ProductTopPublishRequest)
}

// ReleaseProductTopPublishRequest 释放ProductTopPublishRequest
func ReleaseProductTopPublishRequest(v *ProductTopPublishRequest) {
	v.Language = ""
	v.Xml = ""
	v.CatId = 0
	v.ProductId = 0
	poolProductTopPublishRequest.Put(v)
}
