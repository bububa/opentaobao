package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改商品 APIRequest
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
    keywords   []String 

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

func NewAlibabaIcbuProductUpdateRequest() *AlibabaIcbuProductUpdateRequest{
    return &AlibabaIcbuProductUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductUpdateRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.update"
}

func (r AlibabaIcbuProductUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductUpdateRequest) SetAttributes(attributes []ProductAttribute) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetAttributes() []ProductAttribute {
    return r.attributes
}

func (r *AlibabaIcbuProductUpdateRequest) SetBulkDiscountPrices(bulkDiscountPrices []BulkDiscountPrice) error {
    r.bulkDiscountPrices = bulkDiscountPrices
    r.Set("bulk_discount_prices", bulkDiscountPrices)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetBulkDiscountPrices() []BulkDiscountPrice {
    return r.bulkDiscountPrices
}

func (r *AlibabaIcbuProductUpdateRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetCategoryId() int64 {
    return r.categoryId
}

func (r *AlibabaIcbuProductUpdateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetDescription() string {
    return r.description
}

func (r *AlibabaIcbuProductUpdateRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetExtraContext() string {
    return r.extraContext
}

func (r *AlibabaIcbuProductUpdateRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetGroupId() int64 {
    return r.groupId
}

func (r *AlibabaIcbuProductUpdateRequest) SetKeywords(keywords []String) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetKeywords() []String {
    return r.keywords
}

func (r *AlibabaIcbuProductUpdateRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetLanguage() string {
    return r.language
}

func (r *AlibabaIcbuProductUpdateRequest) SetMainImage(mainImage *MainImage) error {
    r.mainImage = mainImage
    r.Set("main_image", mainImage)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetMainImage() *MainImage {
    return r.mainImage
}

func (r *AlibabaIcbuProductUpdateRequest) SetProductSku(productSku *ProductSku) error {
    r.productSku = productSku
    r.Set("product_sku", productSku)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetProductSku() *ProductSku {
    return r.productSku
}

func (r *AlibabaIcbuProductUpdateRequest) SetProductType(productType string) error {
    r.productType = productType
    r.Set("product_type", productType)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetProductType() string {
    return r.productType
}

func (r *AlibabaIcbuProductUpdateRequest) SetSourcingTrade(sourcingTrade *SourcingTrade) error {
    r.sourcingTrade = sourcingTrade
    r.Set("sourcing_trade", sourcingTrade)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetSourcingTrade() *SourcingTrade {
    return r.sourcingTrade
}

func (r *AlibabaIcbuProductUpdateRequest) SetSubject(subject string) error {
    r.subject = subject
    r.Set("subject", subject)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetSubject() string {
    return r.subject
}

func (r *AlibabaIcbuProductUpdateRequest) SetWholesaleTrade(wholesaleTrade *WholesaleTrade) error {
    r.wholesaleTrade = wholesaleTrade
    r.Set("wholesale_trade", wholesaleTrade)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetWholesaleTrade() *WholesaleTrade {
    return r.wholesaleTrade
}

func (r *AlibabaIcbuProductUpdateRequest) SetMarket(market string) error {
    r.market = market
    r.Set("market", market)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetMarket() string {
    return r.market
}

func (r *AlibabaIcbuProductUpdateRequest) SetIsSmartEdit(isSmartEdit bool) error {
    r.isSmartEdit = isSmartEdit
    r.Set("is_smart_edit", isSmartEdit)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetIsSmartEdit() bool {
    return r.isSmartEdit
}

func (r *AlibabaIcbuProductUpdateRequest) SetCustomInfo(customInfo *CustomInfo) error {
    r.customInfo = customInfo
    r.Set("custom_info", customInfo)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetCustomInfo() *CustomInfo {
    return r.customInfo
}

func (r *AlibabaIcbuProductUpdateRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AlibabaIcbuProductUpdateRequest) GetProductId() string {
    return r.productId
}

