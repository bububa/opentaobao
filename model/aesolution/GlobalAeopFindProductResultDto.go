package aesolution

import (
	"sync"
)

// GlobalAeopFindProductResultDto 结构体
type GlobalAeopFindProductResultDto struct {
	// Product properties
	AeopAeProductPropertys []GlobalAeopAeProductProperty `json:"aeop_ae_product_propertys,omitempty" xml:"aeop_ae_product_propertys>global_aeop_ae_product_property,omitempty"`
	// List for multiple skus of the product, expressed in json format.
	AeopAeProductSKUs []GlobalAeopAeProductSku `json:"aeop_ae_product_s_k_us,omitempty" xml:"aeop_ae_product_s_k_us>global_aeop_ae_product_sku,omitempty"`
	// All the groups that the product belongs to.
	GroupIds []int64 `json:"group_ids,omitempty" xml:"group_ids>int64,omitempty"`
	// multi language subject list
	MultiLanguageSubjectList []GlobalSubject `json:"multi_language_subject_list,omitempty" xml:"multi_language_subject_list>global_subject,omitempty"`
	// multo language description list
	MultiLanguageDescriptionList []GlobalDescription `json:"multi_language_description_list,omitempty" xml:"multi_language_description_list>global_description,omitempty"`
	// Required when is_pack_sell equals to true. It means weight to be correspondingly added. Range value: 0.001-500.000, reserving three decimal places and applying scale; Unit: kilogram(s). Please refer to the field is_pack_sell for details.
	AddWeight string `json:"add_weight,omitempty" xml:"add_weight,omitempty"`
	// the Currency code. &#34;USD&#34; will be used as the default value if this information is not provided; Currency code is mandatory for Russian sellers and Spanish sellers. For Russian sellers, RUB should be filled in while EUR for Spanish sellers.
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// Deprecated, please use multi_language_description_list
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// created time
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// modified time
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// Product gross weight. Range value: 0.001-500.000, reserving three decimal places and applying scale; Unit: kilogram(s).
	GrossWeight string `json:"gross_weight,omitempty" xml:"gross_weight,omitempty"`
	// image urls for the product
	ImageURLs string `json:"image_u_r_ls,omitempty" xml:"image_u_r_ls,omitempty"`
	// mobile detail
	MobileDetail string `json:"mobile_detail,omitempty" xml:"mobile_detail,omitempty"`
	// Out of date, please ignore
	OwnerMemberId string `json:"owner_member_id,omitempty" xml:"owner_member_id,omitempty"`
	// Price for product
	ProductPrice string `json:"product_price,omitempty" xml:"product_price,omitempty"`
	// product status type
	ProductStatusType string `json:"product_status_type,omitempty" xml:"product_status_type,omitempty"`
	// Stock reduction strategy. It is divided into 2 types: stock reduction after placing order (place_order_withhold) or stock reduction after payment (payment_success_deduct).
	ReduceStrategy string `json:"reduce_strategy,omitempty" xml:"reduce_strategy,omitempty"`
	// Deprecated, please use multi_language_subject_list
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// The offline date of the product
	WsOfflineDate string `json:"ws_offline_date,omitempty" xml:"ws_offline_date,omitempty"`
	// Required when is_pack_sell equals to true. Value range for pieces to be added: 1-1000. Please refer to the field is_pack_sell for details.
	AddUnit int64 `json:"add_unit,omitempty" xml:"add_unit,omitempty"`
	// Multimedia information
	AeopAEMultimedia *GlobalAeopAeMultimedia `json:"aeop_a_e_multimedia,omitempty" xml:"aeop_a_e_multimedia,omitempty"`
	// Required when is_pack_sell equals to true. It means no additional freight will be charged when the number of pieces to be purchased is under the base unit. Value range: 1-1000.
	BaseUnit int64 `json:"base_unit,omitempty" xml:"base_unit,omitempty"`
	// Bulk discount for wholesale. It must be multiplied by 100 and selected and saved as integer. Value range: 1-99. Note: It is discount, other than discount rate. For example, if the discount is 68, it should be selected and saved as 32. bulk_order and bulk_discount must have value or have no value simultaneously.
	BulkDiscount int64 `json:"bulk_discount,omitempty" xml:"bulk_discount,omitempty"`
	// Minimum bulk for wholesale. Value range: 2-100000. bulk_order and bulk_discount must have value or have no value simultaneously.
	BulkOrder int64 `json:"bulk_order,omitempty" xml:"bulk_order,omitempty"`
	// Product category ID. It must be leaf category to be obtained from the category interface.
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// Stocking period. Value range: 1-60; Unit: day(s). Referring to the preparation time before the order could be dispatched.
	DeliveryTime int64 `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// shipping template id
	FreightTemplateId int64 `json:"freight_template_id,omitempty" xml:"freight_template_id,omitempty"`
	// Group ID that the product belongs to.
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// Number of piece(s) in each pack. In case of packing sale,lotNum&gt;1, and in case of unpacking sale, lotNum=1.
	LotNum int64 `json:"lot_num,omitempty" xml:"lot_num,omitempty"`
	// Out of date, please ignore.
	OwnerMemberSeq int64 `json:"owner_member_seq,omitempty" xml:"owner_member_seq,omitempty"`
	// package height
	PackageHeight int64 `json:"package_height,omitempty" xml:"package_height,omitempty"`
	// package length
	PackageLength int64 `json:"package_length,omitempty" xml:"package_length,omitempty"`
	// package width
	PackageWidth int64 `json:"package_width,omitempty" xml:"package_width,omitempty"`
	// product ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// Product unit
	ProductUnit int64 `json:"product_unit,omitempty" xml:"product_unit,omitempty"`
	// Service template ID which the product is associated with
	PromiseTemplateId int64 `json:"promise_template_id,omitempty" xml:"promise_template_id,omitempty"`
	// Size chart template ID that the product is associated with.
	SizechartId int64 `json:"sizechart_id,omitempty" xml:"sizechart_id,omitempty"`
	// multi country price configuration
	MultiCountryPriceConfiguration *MultiCountryPriceConfigurationDto `json:"multi_country_price_configuration,omitempty" xml:"multi_country_price_configuration,omitempty"`
	// Whether customized weighting is enabled or not. True means customized weighting enabled. When is_pack_sell equals to true, add_unit, add_weight and base_unit must be filled in. For example, base_unit equals to 5, add_unit equals to 2 and add_weight equals to 1.2. It means that Basic shipping cost will cover the first 5 pieces, freight calculating of which will base on a single product. For every additional 2 pieces, the shipping cost will be added to the total shipping cost for 1.2kg
	IsPackSell bool `json:"is_pack_sell,omitempty" xml:"is_pack_sell,omitempty"`
	// Packing sale: true; Unpacking sale: false.
	PackageType bool `json:"package_type,omitempty" xml:"package_type,omitempty"`
}

var poolGlobalAeopFindProductResultDto = sync.Pool{
	New: func() any {
		return new(GlobalAeopFindProductResultDto)
	},
}

// GetGlobalAeopFindProductResultDto() 从对象池中获取GlobalAeopFindProductResultDto
func GetGlobalAeopFindProductResultDto() *GlobalAeopFindProductResultDto {
	return poolGlobalAeopFindProductResultDto.Get().(*GlobalAeopFindProductResultDto)
}

// ReleaseGlobalAeopFindProductResultDto 释放GlobalAeopFindProductResultDto
func ReleaseGlobalAeopFindProductResultDto(v *GlobalAeopFindProductResultDto) {
	v.AeopAeProductPropertys = v.AeopAeProductPropertys[:0]
	v.AeopAeProductSKUs = v.AeopAeProductSKUs[:0]
	v.GroupIds = v.GroupIds[:0]
	v.MultiLanguageSubjectList = v.MultiLanguageSubjectList[:0]
	v.MultiLanguageDescriptionList = v.MultiLanguageDescriptionList[:0]
	v.AddWeight = ""
	v.CurrencyCode = ""
	v.Detail = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.GrossWeight = ""
	v.ImageURLs = ""
	v.MobileDetail = ""
	v.OwnerMemberId = ""
	v.ProductPrice = ""
	v.ProductStatusType = ""
	v.ReduceStrategy = ""
	v.Subject = ""
	v.WsOfflineDate = ""
	v.AddUnit = 0
	v.AeopAEMultimedia = nil
	v.BaseUnit = 0
	v.BulkDiscount = 0
	v.BulkOrder = 0
	v.CategoryId = 0
	v.DeliveryTime = 0
	v.FreightTemplateId = 0
	v.GroupId = 0
	v.LotNum = 0
	v.OwnerMemberSeq = 0
	v.PackageHeight = 0
	v.PackageLength = 0
	v.PackageWidth = 0
	v.ProductId = 0
	v.ProductUnit = 0
	v.PromiseTemplateId = 0
	v.SizechartId = 0
	v.MultiCountryPriceConfiguration = nil
	v.IsPackSell = false
	v.PackageType = false
	poolGlobalAeopFindProductResultDto.Put(v)
}
