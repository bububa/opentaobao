package openmall

import (
	"sync"
)

// TopItemSkuVo 结构体
type TopItemSkuVo struct {
	// 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// sku所对应的销售属性串，由 属性名ID(pid)、属性值ID(vid)组成。  格式为：pid1:vid1;pid2:vid2 ;pid3:vid3…    pid和vid对应的中文名称，可以从item_properties字段中获取。
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// sku描述，以分号分隔描述项，### 分隔翻译内容；注意，当为别名时，翻译项中将没有冒号分隔
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 属于这个sku的商品的数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// sku的id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolTopItemSkuVo = sync.Pool{
	New: func() any {
		return new(TopItemSkuVo)
	},
}

// GetTopItemSkuVo() 从对象池中获取TopItemSkuVo
func GetTopItemSkuVo() *TopItemSkuVo {
	return poolTopItemSkuVo.Get().(*TopItemSkuVo)
}

// ReleaseTopItemSkuVo 释放TopItemSkuVo
func ReleaseTopItemSkuVo(v *TopItemSkuVo) {
	v.Price = ""
	v.Properties = ""
	v.Description = ""
	v.Quantity = 0
	v.SkuId = 0
	poolTopItemSkuVo.Put(v)
}
