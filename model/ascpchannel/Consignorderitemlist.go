package ascpchannel

import (
	"sync"
)

// Consignorderitemlist 结构体
type Consignorderitemlist struct {
	// 包件明细( 包件为包裹信息)
	PackageDetailList []Packagedetaillist `json:"package_detail_list,omitempty" xml:"package_detail_list>packagedetaillist,omitempty"`
	// 交易主单号
	OrderSourceCode string `json:"order_source_code,omitempty" xml:"order_source_code,omitempty"`
	// 交易子单号
	SubSourceCode string `json:"sub_source_code,omitempty" xml:"sub_source_code,omitempty"`
	// 货品名称(货品为商家实际发货货品信息)
	GoodsName string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// 货品体积(m3)
	GoodsVolume string `json:"goods_volume,omitempty" xml:"goods_volume,omitempty"`
	// 货品重量(kg)
	GoodsWeight string `json:"goods_weight,omitempty" xml:"goods_weight,omitempty"`
	// 安装类型 1:整装 2:拆装
	InstallType string `json:"install_type,omitempty" xml:"install_type,omitempty"`
	// 商品名称(商品为店铺销售宝贝信息)
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 扩展字段JSON串
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 包件数量
	PackageQty int64 `json:"package_qty,omitempty" xml:"package_qty,omitempty"`
	// 货品数量
	GoodsQty int64 `json:"goods_qty,omitempty" xml:"goods_qty,omitempty"`
}

var poolConsignorderitemlist = sync.Pool{
	New: func() any {
		return new(Consignorderitemlist)
	},
}

// GetConsignorderitemlist() 从对象池中获取Consignorderitemlist
func GetConsignorderitemlist() *Consignorderitemlist {
	return poolConsignorderitemlist.Get().(*Consignorderitemlist)
}

// ReleaseConsignorderitemlist 释放Consignorderitemlist
func ReleaseConsignorderitemlist(v *Consignorderitemlist) {
	v.PackageDetailList = v.PackageDetailList[:0]
	v.OrderSourceCode = ""
	v.SubSourceCode = ""
	v.GoodsName = ""
	v.GoodsVolume = ""
	v.GoodsWeight = ""
	v.InstallType = ""
	v.ItemName = ""
	v.ItemCode = ""
	v.Feature = ""
	v.PackageQty = 0
	v.GoodsQty = 0
	poolConsignorderitemlist.Put(v)
}
