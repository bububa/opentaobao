package bus

import (
	"sync"
)

// TopInsProduct 结构体
type TopInsProduct struct {
	// 保险模块标题
	InsTitle string `json:"ins_title,omitempty" xml:"ins_title,omitempty"`
	// 保险名称
	InsName string `json:"ins_name,omitempty" xml:"ins_name,omitempty"`
	// 保险利益点信息
	InterestInfo string `json:"interest_info,omitempty" xml:"interest_info,omitempty"`
	// 保险产品码
	ProCode string `json:"pro_code,omitempty" xml:"pro_code,omitempty"`
	// 资源项，图文信息等
	ResourceMap string `json:"resource_map,omitempty" xml:"resource_map,omitempty"`
	// 保险金额
	InsPrice int64 `json:"ins_price,omitempty" xml:"ins_price,omitempty"`
}

var poolTopInsProduct = sync.Pool{
	New: func() any {
		return new(TopInsProduct)
	},
}

// GetTopInsProduct() 从对象池中获取TopInsProduct
func GetTopInsProduct() *TopInsProduct {
	return poolTopInsProduct.Get().(*TopInsProduct)
}

// ReleaseTopInsProduct 释放TopInsProduct
func ReleaseTopInsProduct(v *TopInsProduct) {
	v.InsTitle = ""
	v.InsName = ""
	v.InterestInfo = ""
	v.ProCode = ""
	v.ResourceMap = ""
	v.InsPrice = 0
	poolTopInsProduct.Put(v)
}
