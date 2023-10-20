package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductAddDraftAPIRequest ICBU商品发布草稿接口 API请求
// alibaba.icbu.product.add.draft
//
// 发布商品草稿,支持sourcing/一口价商品，支持英文和多种语言原发商品
type AlibabaIcbuProductAddDraftAPIRequest struct {
	model.Params
	// 商品属性和属性值
	_attributes []ProductAttribute
	// 根据数量设置的折扣价
	_bulkDiscountPrices []BulkDiscountPrice
	// 关键词，不要包含特殊符号（如,;），最多三个
	_keywords []string
	// 商品详情描述，可包含图片中心的图片URL
	_description string
	// 补充信息
	_extraContext string
	// 语种，参见FAQ 语种枚举值
	_language string
	// 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)，值为wholesale时，必须填写wholesale_trade
	_productType string
	// 商品名称，最多128个字符
	_subject string
	// 发布的市场，支持main/onesite，默认main发到主市场，填onesite发布为商机通产品
	_market string
	// 类目ID
	_categoryId int64
	// 分组ID
	_groupId int64
	// 商品主图
	_mainImage *MainImage
	// 商品SKU定义
	_productSku *ProductSku
	// 询盘商品交易信息
	_sourcingTrade *SourcingTrade
	// 在线批发商品交易信息
	_wholesaleTrade *WholesaleTrade
	// 定制信息
	_customInfo *CustomInfo
	// 是否智能编辑，如果不传，默认为false
	_isSmartEdit bool
}

// NewAlibabaIcbuProductAddDraftRequest 初始化AlibabaIcbuProductAddDraftAPIRequest对象
func NewAlibabaIcbuProductAddDraftRequest() *AlibabaIcbuProductAddDraftAPIRequest {
	return &AlibabaIcbuProductAddDraftAPIRequest{
		Params: model.NewParams(17),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductAddDraftAPIRequest) Reset() {
	r._attributes = r._attributes[:0]
	r._bulkDiscountPrices = r._bulkDiscountPrices[:0]
	r._keywords = r._keywords[:0]
	r._description = ""
	r._extraContext = ""
	r._language = ""
	r._productType = ""
	r._subject = ""
	r._market = ""
	r._categoryId = 0
	r._groupId = 0
	r._mainImage = nil
	r._productSku = nil
	r._sourcingTrade = nil
	r._wholesaleTrade = nil
	r._customInfo = nil
	r._isSmartEdit = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductAddDraftAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.add.draft"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductAddDraftAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductAddDraftAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAttributes is Attributes Setter
// 商品属性和属性值
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetAttributes(_attributes []ProductAttribute) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetAttributes() []ProductAttribute {
	return r._attributes
}

// SetBulkDiscountPrices is BulkDiscountPrices Setter
// 根据数量设置的折扣价
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetBulkDiscountPrices(_bulkDiscountPrices []BulkDiscountPrice) error {
	r._bulkDiscountPrices = _bulkDiscountPrices
	r.Set("bulk_discount_prices", _bulkDiscountPrices)
	return nil
}

// GetBulkDiscountPrices BulkDiscountPrices Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetBulkDiscountPrices() []BulkDiscountPrice {
	return r._bulkDiscountPrices
}

// SetKeywords is Keywords Setter
// 关键词，不要包含特殊符号（如,;），最多三个
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetKeywords(_keywords []string) error {
	r._keywords = _keywords
	r.Set("keywords", _keywords)
	return nil
}

// GetKeywords Keywords Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetKeywords() []string {
	return r._keywords
}

// SetDescription is Description Setter
// 商品详情描述，可包含图片中心的图片URL
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetDescription() string {
	return r._description
}

// SetExtraContext is ExtraContext Setter
// 补充信息
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// GetExtraContext ExtraContext Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetExtraContext() string {
	return r._extraContext
}

// SetLanguage is Language Setter
// 语种，参见FAQ 语种枚举值
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetLanguage() string {
	return r._language
}

// SetProductType is ProductType Setter
// 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)，值为wholesale时，必须填写wholesale_trade
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetProductType(_productType string) error {
	r._productType = _productType
	r.Set("product_type", _productType)
	return nil
}

// GetProductType ProductType Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetProductType() string {
	return r._productType
}

// SetSubject is Subject Setter
// 商品名称，最多128个字符
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetSubject(_subject string) error {
	r._subject = _subject
	r.Set("subject", _subject)
	return nil
}

// GetSubject Subject Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetSubject() string {
	return r._subject
}

// SetMarket is Market Setter
// 发布的市场，支持main/onesite，默认main发到主市场，填onesite发布为商机通产品
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetMarket(_market string) error {
	r._market = _market
	r.Set("market", _market)
	return nil
}

// GetMarket Market Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetMarket() string {
	return r._market
}

// SetCategoryId is CategoryId Setter
// 类目ID
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetGroupId is GroupId Setter
// 分组ID
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetGroupId() int64 {
	return r._groupId
}

// SetMainImage is MainImage Setter
// 商品主图
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetMainImage(_mainImage *MainImage) error {
	r._mainImage = _mainImage
	r.Set("main_image", _mainImage)
	return nil
}

// GetMainImage MainImage Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetMainImage() *MainImage {
	return r._mainImage
}

// SetProductSku is ProductSku Setter
// 商品SKU定义
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetProductSku(_productSku *ProductSku) error {
	r._productSku = _productSku
	r.Set("product_sku", _productSku)
	return nil
}

// GetProductSku ProductSku Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetProductSku() *ProductSku {
	return r._productSku
}

// SetSourcingTrade is SourcingTrade Setter
// 询盘商品交易信息
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetSourcingTrade(_sourcingTrade *SourcingTrade) error {
	r._sourcingTrade = _sourcingTrade
	r.Set("sourcing_trade", _sourcingTrade)
	return nil
}

// GetSourcingTrade SourcingTrade Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetSourcingTrade() *SourcingTrade {
	return r._sourcingTrade
}

// SetWholesaleTrade is WholesaleTrade Setter
// 在线批发商品交易信息
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetWholesaleTrade(_wholesaleTrade *WholesaleTrade) error {
	r._wholesaleTrade = _wholesaleTrade
	r.Set("wholesale_trade", _wholesaleTrade)
	return nil
}

// GetWholesaleTrade WholesaleTrade Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetWholesaleTrade() *WholesaleTrade {
	return r._wholesaleTrade
}

// SetCustomInfo is CustomInfo Setter
// 定制信息
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetCustomInfo(_customInfo *CustomInfo) error {
	r._customInfo = _customInfo
	r.Set("custom_info", _customInfo)
	return nil
}

// GetCustomInfo CustomInfo Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetCustomInfo() *CustomInfo {
	return r._customInfo
}

// SetIsSmartEdit is IsSmartEdit Setter
// 是否智能编辑，如果不传，默认为false
func (r *AlibabaIcbuProductAddDraftAPIRequest) SetIsSmartEdit(_isSmartEdit bool) error {
	r._isSmartEdit = _isSmartEdit
	r.Set("is_smart_edit", _isSmartEdit)
	return nil
}

// GetIsSmartEdit IsSmartEdit Getter
func (r AlibabaIcbuProductAddDraftAPIRequest) GetIsSmartEdit() bool {
	return r._isSmartEdit
}

var poolAlibabaIcbuProductAddDraftAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductAddDraftRequest()
	},
}

// GetAlibabaIcbuProductAddDraftRequest 从 sync.Pool 获取 AlibabaIcbuProductAddDraftAPIRequest
func GetAlibabaIcbuProductAddDraftAPIRequest() *AlibabaIcbuProductAddDraftAPIRequest {
	return poolAlibabaIcbuProductAddDraftAPIRequest.Get().(*AlibabaIcbuProductAddDraftAPIRequest)
}

// ReleaseAlibabaIcbuProductAddDraftAPIRequest 将 AlibabaIcbuProductAddDraftAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductAddDraftAPIRequest(v *AlibabaIcbuProductAddDraftAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductAddDraftAPIRequest.Put(v)
}
