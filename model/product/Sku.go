package product

import (
	"sync"
)

// Sku 结构体
type Sku struct {
	// sku的销售属性组合字符串（颜色，大小，等等，可通过类目API获取某类目下的销售属性）,格式是p1:v1;p2:v2
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// sku所对应的销售属性的中文名字串，格式如：pid1:vid1:pid_name1:vid_name1;pid2:vid2:pid_name2:vid_name2&amp;hellip;&amp;hellip;
	PropertiesName string `json:"properties_name,omitempty" xml:"properties_name,omitempty"`
	// 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商家设置的外部id。天猫和集市的卖家，需要登录才能获取到自己的商家编码，不能获取到他人的商家编码。
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 商品级别的条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 商品SKU优惠价格
	PromotedPrice string `json:"promoted_price,omitempty" xml:"promoted_price,omitempty"`
	// SKU图片
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// sku的id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolSku = sync.Pool{
	New: func() any {
		return new(Sku)
	},
}

// GetSku() 从对象池中获取Sku
func GetSku() *Sku {
	return poolSku.Get().(*Sku)
}

// ReleaseSku 释放Sku
func ReleaseSku(v *Sku) {
	v.Properties = ""
	v.PropertiesName = ""
	v.Price = ""
	v.OuterId = ""
	v.Barcode = ""
	v.PromotedPrice = ""
	v.PicUrl = ""
	v.SkuId = 0
	poolSku.Put(v)
}
