package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ICBU商品发布草稿接口 APIRequest
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

func NewAlibabaIcbuProductAddDraftRequest() *AlibabaIcbuProductAddDraftRequest{
    return &AlibabaIcbuProductAddDraftRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductAddDraftRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.add.draft"
}

func (r AlibabaIcbuProductAddDraftRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductAddDraftRequest) SetAttributes(attributes []ProductAttribute) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetAttributes() []ProductAttribute {
    return r.attributes
}

func (r *AlibabaIcbuProductAddDraftRequest) SetBulkDiscountPrices(bulkDiscountPrices []BulkDiscountPrice) error {
    r.bulkDiscountPrices = bulkDiscountPrices
    r.Set("bulk_discount_prices", bulkDiscountPrices)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetBulkDiscountPrices() []BulkDiscountPrice {
    return r.bulkDiscountPrices
}

func (r *AlibabaIcbuProductAddDraftRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetCategoryId() int64 {
    return r.categoryId
}

func (r *AlibabaIcbuProductAddDraftRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetDescription() string {
    return r.description
}

func (r *AlibabaIcbuProductAddDraftRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetExtraContext() string {
    return r.extraContext
}

func (r *AlibabaIcbuProductAddDraftRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetGroupId() int64 {
    return r.groupId
}

func (r *AlibabaIcbuProductAddDraftRequest) SetKeywords(keywords []string) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetKeywords() []string {
    return r.keywords
}

func (r *AlibabaIcbuProductAddDraftRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetLanguage() string {
    return r.language
}

func (r *AlibabaIcbuProductAddDraftRequest) SetMainImage(mainImage *MainImage) error {
    r.mainImage = mainImage
    r.Set("main_image", mainImage)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetMainImage() *MainImage {
    return r.mainImage
}

func (r *AlibabaIcbuProductAddDraftRequest) SetProductSku(productSku *ProductSku) error {
    r.productSku = productSku
    r.Set("product_sku", productSku)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetProductSku() *ProductSku {
    return r.productSku
}

func (r *AlibabaIcbuProductAddDraftRequest) SetProductType(productType string) error {
    r.productType = productType
    r.Set("product_type", productType)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetProductType() string {
    return r.productType
}

func (r *AlibabaIcbuProductAddDraftRequest) SetSourcingTrade(sourcingTrade *SourcingTrade) error {
    r.sourcingTrade = sourcingTrade
    r.Set("sourcing_trade", sourcingTrade)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetSourcingTrade() *SourcingTrade {
    return r.sourcingTrade
}

func (r *AlibabaIcbuProductAddDraftRequest) SetSubject(subject string) error {
    r.subject = subject
    r.Set("subject", subject)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetSubject() string {
    return r.subject
}

func (r *AlibabaIcbuProductAddDraftRequest) SetWholesaleTrade(wholesaleTrade *WholesaleTrade) error {
    r.wholesaleTrade = wholesaleTrade
    r.Set("wholesale_trade", wholesaleTrade)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetWholesaleTrade() *WholesaleTrade {
    return r.wholesaleTrade
}

func (r *AlibabaIcbuProductAddDraftRequest) SetMarket(market string) error {
    r.market = market
    r.Set("market", market)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetMarket() string {
    return r.market
}

func (r *AlibabaIcbuProductAddDraftRequest) SetIsSmartEdit(isSmartEdit bool) error {
    r.isSmartEdit = isSmartEdit
    r.Set("is_smart_edit", isSmartEdit)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetIsSmartEdit() bool {
    return r.isSmartEdit
}

func (r *AlibabaIcbuProductAddDraftRequest) SetCustomInfo(customInfo *CustomInfo) error {
    r.customInfo = customInfo
    r.Set("custom_info", customInfo)
    return nil
}

func (r AlibabaIcbuProductAddDraftRequest) GetCustomInfo() *CustomInfo {
    return r.customInfo
}

