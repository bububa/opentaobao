package tbitem

import (
	"sync"
)

// Sku 结构体
type Sku struct {
	// sku最后修改日期 时间格式：yyyy-MM-dd HH:mm:ss
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 商品级别的条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// sku创建日期 时间格式：yyyy-MM-dd HH:mm:ss
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 商家设置的外部id。天猫和集市的卖家，需要登录才能获取到自己的商家编码，不能获取到他人的商家编码。
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// sku的销售属性组合字符串（颜色，大小，等等，可通过类目API获取某类目下的销售属性）,格式是p1:v1;p2:v2
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// sku所对应的销售属性的中文名字串，格式如：pid1:vid1:pid_name1:vid_name1;pid2:vid2:pid_name2:vid_name2……
	PropertiesName string `json:"properties_name,omitempty" xml:"properties_name,omitempty"`
	// sku所属商品id(注意：iid近期即将废弃，请用num_iid参数)
	Iid string `json:"iid,omitempty" xml:"iid,omitempty"`
	// sku状态。 normal:正常 ；delete:删除
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// sku的id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 属于这个sku的商品的数量，
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// sku所属商品数字id
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 表示SKu上的产品规格信息
	SkuSpecId int64 `json:"sku_spec_id,omitempty" xml:"sku_spec_id,omitempty"`
	// 商品在付款减库存的状态下，该sku上未付款的订单数量
	WithHoldQuantity int64 `json:"with_hold_quantity,omitempty" xml:"with_hold_quantity,omitempty"`
	// skuFeature
	SkuFeature *SkuFeature `json:"sku_feature,omitempty" xml:"sku_feature,omitempty"`
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
	v.Modified = ""
	v.Barcode = ""
	v.Created = ""
	v.OuterId = ""
	v.Price = ""
	v.Properties = ""
	v.PropertiesName = ""
	v.Iid = ""
	v.Status = ""
	v.SkuId = 0
	v.Quantity = 0
	v.NumIid = 0
	v.SkuSpecId = 0
	v.WithHoldQuantity = 0
	v.SkuFeature = nil
	poolSku.Put(v)
}
