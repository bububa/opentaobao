package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改商品 API请求
alibaba.icbu.product.update

修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
*/
type AlibabaIcbuProductUpdateRequest struct {
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
    // 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)
    productType   string
    // 询盘商品交易信息
    sourcingTrade   *SourcingTrade
    // 商品名称，最多128个字符
    subject   string
    // 在线批发商品交易信息
    wholesaleTrade   *WholesaleTrade
    // 发布的市场，支持main/onesite，默认main发到主市场，填onesite发布为商机通产品
    market   string
    // 智能编辑，不填写使用原来的。注意必须和详情的格式一致
    isSmartEdit   bool
    // 定制信息
    customInfo   *CustomInfo
    // 混淆商品ID
    productId   string
}

// 初始化AlibabaIcbuProductUpdateRequest对象
func NewAlibabaIcbuProductUpdateRequest() *AlibabaIcbuProductUpdateRequest{
    return &AlibabaIcbuProductUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductUpdateRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Attributes Setter
// 商品属性和属性值
func (r *AlibabaIcbuProductUpdateRequest) SetAttributes(attributes []ProductAttribute) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

// Attributes Getter
func (r AlibabaIcbuProductUpdateRequest) GetAttributes() []ProductAttribute {
    return r.attributes
}
// BulkDiscountPrices Setter
// 根据数量设置的折扣价
func (r *AlibabaIcbuProductUpdateRequest) SetBulkDiscountPrices(bulkDiscountPrices []BulkDiscountPrice) error {
    r.bulkDiscountPrices = bulkDiscountPrices
    r.Set("bulk_discount_prices", bulkDiscountPrices)
    return nil
}

// BulkDiscountPrices Getter
func (r AlibabaIcbuProductUpdateRequest) GetBulkDiscountPrices() []BulkDiscountPrice {
    return r.bulkDiscountPrices
}
// CategoryId Setter
// 类目ID
func (r *AlibabaIcbuProductUpdateRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r AlibabaIcbuProductUpdateRequest) GetCategoryId() int64 {
    return r.categoryId
}
// Description Setter
// 商品详情描述，可包含图片中心的图片URL
func (r *AlibabaIcbuProductUpdateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r AlibabaIcbuProductUpdateRequest) GetDescription() string {
    return r.description
}
// ExtraContext Setter
// 补充信息
func (r *AlibabaIcbuProductUpdateRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

// ExtraContext Getter
func (r AlibabaIcbuProductUpdateRequest) GetExtraContext() string {
    return r.extraContext
}
// GroupId Setter
// 分组ID
func (r *AlibabaIcbuProductUpdateRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r AlibabaIcbuProductUpdateRequest) GetGroupId() int64 {
    return r.groupId
}
// Keywords Setter
// 关键词，不要包含特殊符号（如,;），最多三个
func (r *AlibabaIcbuProductUpdateRequest) SetKeywords(keywords []string) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

// Keywords Getter
func (r AlibabaIcbuProductUpdateRequest) GetKeywords() []string {
    return r.keywords
}
// Language Setter
// 语种，参见FAQ 语种枚举值
func (r *AlibabaIcbuProductUpdateRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

// Language Getter
func (r AlibabaIcbuProductUpdateRequest) GetLanguage() string {
    return r.language
}
// MainImage Setter
// 商品主图
func (r *AlibabaIcbuProductUpdateRequest) SetMainImage(mainImage *MainImage) error {
    r.mainImage = mainImage
    r.Set("main_image", mainImage)
    return nil
}

// MainImage Getter
func (r AlibabaIcbuProductUpdateRequest) GetMainImage() *MainImage {
    return r.mainImage
}
// ProductSku Setter
// 商品SKU定义
func (r *AlibabaIcbuProductUpdateRequest) SetProductSku(productSku *ProductSku) error {
    r.productSku = productSku
    r.Set("product_sku", productSku)
    return nil
}

// ProductSku Getter
func (r AlibabaIcbuProductUpdateRequest) GetProductSku() *ProductSku {
    return r.productSku
}
// ProductType Setter
// 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)
func (r *AlibabaIcbuProductUpdateRequest) SetProductType(productType string) error {
    r.productType = productType
    r.Set("product_type", productType)
    return nil
}

// ProductType Getter
func (r AlibabaIcbuProductUpdateRequest) GetProductType() string {
    return r.productType
}
// SourcingTrade Setter
// 询盘商品交易信息
func (r *AlibabaIcbuProductUpdateRequest) SetSourcingTrade(sourcingTrade *SourcingTrade) error {
    r.sourcingTrade = sourcingTrade
    r.Set("sourcing_trade", sourcingTrade)
    return nil
}

// SourcingTrade Getter
func (r AlibabaIcbuProductUpdateRequest) GetSourcingTrade() *SourcingTrade {
    return r.sourcingTrade
}
// Subject Setter
// 商品名称，最多128个字符
func (r *AlibabaIcbuProductUpdateRequest) SetSubject(subject string) error {
    r.subject = subject
    r.Set("subject", subject)
    return nil
}

// Subject Getter
func (r AlibabaIcbuProductUpdateRequest) GetSubject() string {
    return r.subject
}
// WholesaleTrade Setter
// 在线批发商品交易信息
func (r *AlibabaIcbuProductUpdateRequest) SetWholesaleTrade(wholesaleTrade *WholesaleTrade) error {
    r.wholesaleTrade = wholesaleTrade
    r.Set("wholesale_trade", wholesaleTrade)
    return nil
}

// WholesaleTrade Getter
func (r AlibabaIcbuProductUpdateRequest) GetWholesaleTrade() *WholesaleTrade {
    return r.wholesaleTrade
}
// Market Setter
// 发布的市场，支持main/onesite，默认main发到主市场，填onesite发布为商机通产品
func (r *AlibabaIcbuProductUpdateRequest) SetMarket(market string) error {
    r.market = market
    r.Set("market", market)
    return nil
}

// Market Getter
func (r AlibabaIcbuProductUpdateRequest) GetMarket() string {
    return r.market
}
// IsSmartEdit Setter
// 智能编辑，不填写使用原来的。注意必须和详情的格式一致
func (r *AlibabaIcbuProductUpdateRequest) SetIsSmartEdit(isSmartEdit bool) error {
    r.isSmartEdit = isSmartEdit
    r.Set("is_smart_edit", isSmartEdit)
    return nil
}

// IsSmartEdit Getter
func (r AlibabaIcbuProductUpdateRequest) GetIsSmartEdit() bool {
    return r.isSmartEdit
}
// CustomInfo Setter
// 定制信息
func (r *AlibabaIcbuProductUpdateRequest) SetCustomInfo(customInfo *CustomInfo) error {
    r.customInfo = customInfo
    r.Set("custom_info", customInfo)
    return nil
}

// CustomInfo Getter
func (r AlibabaIcbuProductUpdateRequest) GetCustomInfo() *CustomInfo {
    return r.customInfo
}
// ProductId Setter
// 混淆商品ID
func (r *AlibabaIcbuProductUpdateRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AlibabaIcbuProductUpdateRequest) GetProductId() string {
    return r.productId
}
