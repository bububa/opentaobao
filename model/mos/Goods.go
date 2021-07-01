package mos

// Goods 结构体
type Goods struct {
	// 商品编码
	N0803 string `json:"n0803,omitempty" xml:"n0803,omitempty"`
	// 商品专柜
	N0804 string `json:"n0804,omitempty" xml:"n0804,omitempty"`
	// 商品单价
	N0805 string `json:"n0805,omitempty" xml:"n0805,omitempty"`
	// 商品数量
	N0806 string `json:"n0806,omitempty" xml:"n0806,omitempty"`
	// 商品折扣
	N0807 string `json:"n0807,omitempty" xml:"n0807,omitempty"`
	// 商品实销金额
	N0808 string `json:"n0808,omitempty" xml:"n0808,omitempty"`
	// 营业员号
	N0812 string `json:"n0812,omitempty" xml:"n0812,omitempty"`
	// 商品行号
	N0814 string `json:"n0814,omitempty" xml:"n0814,omitempty"`
	// pos满送金额
	N0817 string `json:"n0817,omitempty" xml:"n0817,omitempty"`
	// 后台折扣额
	N0850 string `json:"n0850,omitempty" xml:"n0850,omitempty"`
	// 前台折扣额
	N0851 string `json:"n0851,omitempty" xml:"n0851,omitempty"`
	// 会员折扣额
	N0852 string `json:"n0852,omitempty" xml:"n0852,omitempty"`
	// 满减金额
	N0856 string `json:"n0856,omitempty" xml:"n0856,omitempty"`
	// 单品商品唯一码
	Productid string `json:"productid,omitempty" xml:"productid,omitempty"`
	// 单品商品唯一码（字符串）
	Productstr string `json:"productstr,omitempty" xml:"productstr,omitempty"`
	// Sku
	Productsku string `json:"productsku,omitempty" xml:"productsku,omitempty"`
	// 订单号
	Productticket string `json:"productticket,omitempty" xml:"productticket,omitempty"`
	// 单品类别
	Producttype string `json:"producttype,omitempty" xml:"producttype,omitempty"`
	// 商品名称
	Comname string `json:"comname,omitempty" xml:"comname,omitempty"`
	// 商品店号
	Comstoreno string `json:"comstoreno,omitempty" xml:"comstoreno,omitempty"`
	// 商品开票应付金额
	N0809 string `json:"n0809,omitempty" xml:"n0809,omitempty"`
	// 满减活动id
	PromtionId string `json:"promtion_id,omitempty" xml:"promtion_id,omitempty"`
}
