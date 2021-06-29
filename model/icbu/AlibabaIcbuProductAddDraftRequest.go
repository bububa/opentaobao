package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ICBU商品发布草稿接口 API请求
alibaba.icbu.product.add.draft

发布商品草稿,支持sourcing/一口价商品，支持英文和多种语言原发商品
*/
type AlibabaIcbuProductAddDraftRequest struct {
    model.Params
    // 商品属性和属性值
    attributes   []ProductAttribute
    // 根据数量设置的折扣价
    bulkDiscountPrices   []BulkDiscountPrice
    // 类目ID
    categoryId   int64
    // 商品详情描述，可包含图片中心的图片URL
    description   string
    // 补充信息
    extraContext   string
    // 分组ID
    groupId   int64
    // 关键词，不要包含特殊符号（如,;），最多三个
    keywords   []string
    // 语种，参见FAQ 语种枚举值
    language   string
    // 商品主图
    mainImage   *MainImage
    // 商品SKU定义
    productSku   *ProductSku
    // 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)，值为wholesale时，必须填写wholesale_trade
    productType   string
    // 询盘商品交易信息
    sourcingTrade   *SourcingTrade
    // 商品名称，最多128个字符
    subject   string
    // 在线批发商品交易信息
    wholesaleTrade   *WholesaleTrade
    // 发布的市场，支持main/onesite，默认main发到主市场，填onesite发布为商机通产品
    market   string
    // 是否智能编辑，如果不传，默认为false
    isSmartEdit   bool
    // 定制信息
    customInfo   *CustomInfo
}

// 初始化AlibabaIcbuProductAddDraftRequest对象
func NewAlibabaIcbuProductAddDraftRequest() *AlibabaIcbuProductAddDraftRequest{
    return &AlibabaIcbuProductAddDraftRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductAddDraftRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.add.draft"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductAddDraftRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Attributes Setter
// 商品属性和属性值
func (r *AlibabaIcbuProductAddDraftRequest) SetAttributes(attributes []ProductAttribute) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

// Attributes Getter
func (r AlibabaIcbuProductAddDraftRequest) GetAttributes() []ProductAttribute {
    return r.attributes
}
// BulkDiscountPrices Setter
// 根据数量设置的折扣价
func (r *AlibabaIcbuProductAddDraftRequest) SetBulkDiscountPrices(bulkDiscountPrices []BulkDiscountPrice) error {
    r.bulkDiscountPrices = bulkDiscountPrices
    r.Set("bulk_discount_prices", bulkDiscountPrices)
    return nil
}

// BulkDiscountPrices Getter
func (r AlibabaIcbuProductAddDraftRequest) GetBulkDiscountPrices() []BulkDiscountPrice {
    return r.bulkDiscountPrices
}
// CategoryId Setter
// 类目ID
func (r *AlibabaIcbuProductAddDraftRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r AlibabaIcbuProductAddDraftRequest) GetCategoryId() int64 {
    return r.categoryId
}
// Description Setter
// 商品详情描述，可包含图片中心的图片URL
func (r *AlibabaIcbuProductAddDraftRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r AlibabaIcbuProductAddDraftRequest) GetDescription() string {
    return r.description
}
// ExtraContext Setter
// 补充信息
func (r *AlibabaIcbuProductAddDraftRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

// ExtraContext Getter
func (r AlibabaIcbuProductAddDraftRequest) GetExtraContext() string {
    return r.extraContext
}
// GroupId Setter
// 分组ID
func (r *AlibabaIcbuProductAddDraftRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r AlibabaIcbuProductAddDraftRequest) GetGroupId() int64 {
    return r.groupId
}
// Keywords Setter
// 关键词，不要包含特殊符号（如,;），最多三个
func (r *AlibabaIcbuProductAddDraftRequest) SetKeywords(keywords []string) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

// Keywords Getter
func (r AlibabaIcbuProductAddDraftRequest) GetKeywords() []string {
    return r.keywords
}
// Language Setter
// 语种，参见FAQ 语种枚举值
func (r *AlibabaIcbuProductAddDraftRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

// Language Getter
func (r AlibabaIcbuProductAddDraftRequest) GetLanguage() string {
    return r.language
}
// MainImage Setter
// 商品主图
func (r *AlibabaIcbuProductAddDraftRequest) SetMainImage(mainImage *MainImage) error {
    r.mainImage = mainImage
    r.Set("main_image", mainImage)
    return nil
}

// MainImage Getter
func (r AlibabaIcbuProductAddDraftRequest) GetMainImage() *MainImage {
    return r.mainImage
}
// ProductSku Setter
// 商品SKU定义
func (r *AlibabaIcbuProductAddDraftRequest) SetProductSku(productSku *ProductSku) error {
    r.productSku = productSku
    r.Set("product_sku", productSku)
    return nil
}

// ProductSku Getter
func (r AlibabaIcbuProductAddDraftRequest) GetProductSku() *ProductSku {
    return r.productSku
}
// ProductType Setter
// 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)，值为wholesale时，必须填写wholesale_trade
func (r *AlibabaIcbuProductAddDraftRequest) SetProductType(productType string) error {
    r.productType = productType
    r.Set("product_type", productType)
    return nil
}

// ProductType Getter
func (r AlibabaIcbuProductAddDraftRequest) GetProductType() string {
    return r.productType
}
// SourcingTrade Setter
// 询盘商品交易信息
func (r *AlibabaIcbuProductAddDraftRequest) SetSourcingTrade(sourcingTrade *SourcingTrade) error {
    r.sourcingTrade = sourcingTrade
    r.Set("sourcing_trade", sourcingTrade)
    return nil
}

// SourcingTrade Getter
func (r AlibabaIcbuProductAddDraftRequest) GetSourcingTrade() *SourcingTrade {
    return r.sourcingTrade
}
// Subject Setter
// 商品名称，最多128个字符
func (r *AlibabaIcbuProductAddDraftRequest) SetSubject(subject string) error {
    r.subject = subject
    r.Set("subject", subject)
    return nil
}

// Subject Getter
func (r AlibabaIcbuProductAddDraftRequest) GetSubject() string {
    return r.subject
}
// WholesaleTrade Setter
// 在线批发商品交易信息
func (r *AlibabaIcbuProductAddDraftRequest) SetWholesaleTrade(wholesaleTrade *WholesaleTrade) error {
    r.wholesaleTrade = wholesaleTrade
    r.Set("wholesale_trade", wholesaleTrade)
    return nil
}

// WholesaleTrade Getter
func (r AlibabaIcbuProductAddDraftRequest) GetWholesaleTrade() *WholesaleTrade {
    return r.wholesaleTrade
}
// Market Setter
// 发布的市场，支持main/onesite，默认main发到主市场，填onesite发布为商机通产品
func (r *AlibabaIcbuProductAddDraftRequest) SetMarket(market string) error {
    r.market = market
    r.Set("market", market)
    return nil
}

// Market Getter
func (r AlibabaIcbuProductAddDraftRequest) GetMarket() string {
    return r.market
}
// IsSmartEdit Setter
// 是否智能编辑，如果不传，默认为false
func (r *AlibabaIcbuProductAddDraftRequest) SetIsSmartEdit(isSmartEdit bool) error {
    r.isSmartEdit = isSmartEdit
    r.Set("is_smart_edit", isSmartEdit)
    return nil
}

// IsSmartEdit Getter
func (r AlibabaIcbuProductAddDraftRequest) GetIsSmartEdit() bool {
    return r.isSmartEdit
}
// CustomInfo Setter
// 定制信息
func (r *AlibabaIcbuProductAddDraftRequest) SetCustomInfo(customInfo *CustomInfo) error {
    r.customInfo = customInfo
    r.Set("custom_info", customInfo)
    return nil
}

// CustomInfo Getter
func (r AlibabaIcbuProductAddDraftRequest) GetCustomInfo() *CustomInfo {
    return r.customInfo
}
