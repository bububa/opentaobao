package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductUpdateFieldAPIRequest 商品按字段更新 API请求
// alibaba.icbu.product.update.field
//
// 按字段修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
type AlibabaIcbuProductUpdateFieldAPIRequest struct {
	model.Params
	// 商品属性和属性值
	_attributes []ProductAttribute
	// 根据数量设置的折扣价
	_bulkDiscountPrices []BulkDiscountPrice
	// 类目ID
	_categoryId int64
	// 商品详情描述，可包含图片中心的图片URL
	_description string
	// 补充信息
	_extraContext string
	// 分组ID
	_groupId int64
	// 关键词，不要包含特殊符号（如,;），最多三个
	_keywords []string
	// 语种，当前只有english
	_language string
	// 商品主图
	_mainImage *MainImage
	// 商品SKU定义
	_productSku *ProductSku
	// 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)
	_productType string
	// 询盘商品交易信息
	_sourcingTrade *SourcingTrade
	// 商品名称，最多128个字符
	_subject string
	// 在线批发商品交易信息
	_wholesaleTrade *WholesaleTrade
	// 发布的市场，支持main/onesite，默认main发到主市场，填onesite发布为商机通产品
	_market string
	// 定制信息
	_customInfo *CustomInfo
	// 商品详情种类，true表示智能编辑，不填默认取商品原来的详情种类
	_isSmartEdit bool
	// 使用SKU价的时候需要传入这个参数
	_useSkuPrice bool
	// 混淆商品ID
	_productId string
}

// NewAlibabaIcbuProductUpdateFieldRequest 初始化AlibabaIcbuProductUpdateFieldAPIRequest对象
func NewAlibabaIcbuProductUpdateFieldRequest() *AlibabaIcbuProductUpdateFieldAPIRequest {
	return &AlibabaIcbuProductUpdateFieldAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.update.field"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Attributes Setter
// 商品属性和属性值
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetAttributes(_attributes []ProductAttribute) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// Get Attributes Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetAttributes() []ProductAttribute {
	return r._attributes
}

// Set is BulkDiscountPrices Setter
// 根据数量设置的折扣价
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetBulkDiscountPrices(_bulkDiscountPrices []BulkDiscountPrice) error {
	r._bulkDiscountPrices = _bulkDiscountPrices
	r.Set("bulk_discount_prices", _bulkDiscountPrices)
	return nil
}

// Get BulkDiscountPrices Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetBulkDiscountPrices() []BulkDiscountPrice {
	return r._bulkDiscountPrices
}

// Set is CategoryId Setter
// 类目ID
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// Get CategoryId Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// Set is Description Setter
// 商品详情描述，可包含图片中心的图片URL
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// Get Description Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetDescription() string {
	return r._description
}

// Set is ExtraContext Setter
// 补充信息
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// Get ExtraContext Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetExtraContext() string {
	return r._extraContext
}

// Set is GroupId Setter
// 分组ID
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// Get GroupId Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetGroupId() int64 {
	return r._groupId
}

// Set is Keywords Setter
// 关键词，不要包含特殊符号（如,;），最多三个
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetKeywords(_keywords []string) error {
	r._keywords = _keywords
	r.Set("keywords", _keywords)
	return nil
}

// Get Keywords Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetKeywords() []string {
	return r._keywords
}

// Set is Language Setter
// 语种，当前只有english
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// Get Language Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetLanguage() string {
	return r._language
}

// Set is MainImage Setter
// 商品主图
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetMainImage(_mainImage *MainImage) error {
	r._mainImage = _mainImage
	r.Set("main_image", _mainImage)
	return nil
}

// Get MainImage Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetMainImage() *MainImage {
	return r._mainImage
}

// Set is ProductSku Setter
// 商品SKU定义
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetProductSku(_productSku *ProductSku) error {
	r._productSku = _productSku
	r.Set("product_sku", _productSku)
	return nil
}

// Get ProductSku Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetProductSku() *ProductSku {
	return r._productSku
}

// Set is ProductType Setter
// 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetProductType(_productType string) error {
	r._productType = _productType
	r.Set("product_type", _productType)
	return nil
}

// Get ProductType Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetProductType() string {
	return r._productType
}

// Set is SourcingTrade Setter
// 询盘商品交易信息
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetSourcingTrade(_sourcingTrade *SourcingTrade) error {
	r._sourcingTrade = _sourcingTrade
	r.Set("sourcing_trade", _sourcingTrade)
	return nil
}

// Get SourcingTrade Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetSourcingTrade() *SourcingTrade {
	return r._sourcingTrade
}

// Set is Subject Setter
// 商品名称，最多128个字符
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetSubject(_subject string) error {
	r._subject = _subject
	r.Set("subject", _subject)
	return nil
}

// Get Subject Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetSubject() string {
	return r._subject
}

// Set is WholesaleTrade Setter
// 在线批发商品交易信息
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetWholesaleTrade(_wholesaleTrade *WholesaleTrade) error {
	r._wholesaleTrade = _wholesaleTrade
	r.Set("wholesale_trade", _wholesaleTrade)
	return nil
}

// Get WholesaleTrade Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetWholesaleTrade() *WholesaleTrade {
	return r._wholesaleTrade
}

// Set is Market Setter
// 发布的市场，支持main/onesite，默认main发到主市场，填onesite发布为商机通产品
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetMarket(_market string) error {
	r._market = _market
	r.Set("market", _market)
	return nil
}

// Get Market Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetMarket() string {
	return r._market
}

// Set is CustomInfo Setter
// 定制信息
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetCustomInfo(_customInfo *CustomInfo) error {
	r._customInfo = _customInfo
	r.Set("custom_info", _customInfo)
	return nil
}

// Get CustomInfo Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetCustomInfo() *CustomInfo {
	return r._customInfo
}

// Set is IsSmartEdit Setter
// 商品详情种类，true表示智能编辑，不填默认取商品原来的详情种类
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetIsSmartEdit(_isSmartEdit bool) error {
	r._isSmartEdit = _isSmartEdit
	r.Set("is_smart_edit", _isSmartEdit)
	return nil
}

// Get IsSmartEdit Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetIsSmartEdit() bool {
	return r._isSmartEdit
}

// Set is UseSkuPrice Setter
// 使用SKU价的时候需要传入这个参数
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetUseSkuPrice(_useSkuPrice bool) error {
	r._useSkuPrice = _useSkuPrice
	r.Set("use_sku_price", _useSkuPrice)
	return nil
}

// Get UseSkuPrice Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetUseSkuPrice() bool {
	return r._useSkuPrice
}

// Set is ProductId Setter
// 混淆商品ID
func (r *AlibabaIcbuProductUpdateFieldAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r AlibabaIcbuProductUpdateFieldAPIRequest) GetProductId() string {
	return r._productId
}
