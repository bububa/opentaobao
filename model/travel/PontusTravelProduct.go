package travel

import (
	"sync"
)

// PontusTravelProduct 结构体
type PontusTravelProduct struct {
	// 资源元素的外部商家编码
	ElementId string `json:"element_id,omitempty" xml:"element_id,omitempty"`
	// 元素份数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}

var poolPontusTravelProduct = sync.Pool{
	New: func() any {
		return new(PontusTravelProduct)
	},
}

// GetPontusTravelProduct() 从对象池中获取PontusTravelProduct
func GetPontusTravelProduct() *PontusTravelProduct {
	return poolPontusTravelProduct.Get().(*PontusTravelProduct)
}

// ReleasePontusTravelProduct 释放PontusTravelProduct
func ReleasePontusTravelProduct(v *PontusTravelProduct) {
	v.ElementId = ""
	v.Num = 0
	poolPontusTravelProduct.Put(v)
}
