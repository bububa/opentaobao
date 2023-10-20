package icburfq

import (
	"sync"
)

// BuyRequestSearchDetailRemoteDto 结构体
type BuyRequestSearchDetailRemoteDto struct {
	// 语种
	LangSrc string `json:"lang_src,omitempty" xml:"lang_src,omitempty"`
	// 供应商国家
	SupplierCountrys string `json:"supplier_countrys,omitempty" xml:"supplier_countrys,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 附加名称
	AnnexNames string `json:"annex_names,omitempty" xml:"annex_names,omitempty"`
	// 付款条件
	PaymentTerms string `json:"payment_terms,omitempty" xml:"payment_terms,omitempty"`
	// 目的港
	DestinationPort string `json:"destination_port,omitempty" xml:"destination_port,omitempty"`
	// 价格单位
	FobPriceUnit string `json:"fob_price_unit,omitempty" xml:"fob_price_unit,omitempty"`
	// 价格
	FobPrice string `json:"fob_price,omitempty" xml:"fob_price,omitempty"`
	// 发运条件
	ShippingTerms string `json:"shipping_terms,omitempty" xml:"shipping_terms,omitempty"`
	// 国家简称
	CountrySimple string `json:"country_simple,omitempty" xml:"country_simple,omitempty"`
	// 数量单位
	QuantityUnit string `json:"quantity_unit,omitempty" xml:"quantity_unit,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// RFQ ID
	RfqId string `json:"rfq_id,omitempty" xml:"rfq_id,omitempty"`
	// 类目ID
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 剩余报价数量
	LeftCount int64 `json:"left_count,omitempty" xml:"left_count,omitempty"`
	// 开放时间
	OpenTime int64 `json:"open_time,omitempty" xml:"open_time,omitempty"`
	// 过期值
	ExpirateTime int64 `json:"expirate_time,omitempty" xml:"expirate_time,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolBuyRequestSearchDetailRemoteDto = sync.Pool{
	New: func() any {
		return new(BuyRequestSearchDetailRemoteDto)
	},
}

// GetBuyRequestSearchDetailRemoteDto() 从对象池中获取BuyRequestSearchDetailRemoteDto
func GetBuyRequestSearchDetailRemoteDto() *BuyRequestSearchDetailRemoteDto {
	return poolBuyRequestSearchDetailRemoteDto.Get().(*BuyRequestSearchDetailRemoteDto)
}

// ReleaseBuyRequestSearchDetailRemoteDto 释放BuyRequestSearchDetailRemoteDto
func ReleaseBuyRequestSearchDetailRemoteDto(v *BuyRequestSearchDetailRemoteDto) {
	v.LangSrc = ""
	v.SupplierCountrys = ""
	v.CategoryName = ""
	v.AnnexNames = ""
	v.PaymentTerms = ""
	v.DestinationPort = ""
	v.FobPriceUnit = ""
	v.FobPrice = ""
	v.ShippingTerms = ""
	v.CountrySimple = ""
	v.QuantityUnit = ""
	v.Status = ""
	v.Description = ""
	v.Subject = ""
	v.RfqId = ""
	v.CategoryId = 0
	v.LeftCount = 0
	v.OpenTime = 0
	v.ExpirateTime = 0
	v.Quantity = 0
	poolBuyRequestSearchDetailRemoteDto.Put(v)
}
