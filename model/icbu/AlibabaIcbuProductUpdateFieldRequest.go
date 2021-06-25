package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品按字段更新 APIRequest
alibaba.icbu.product.update.field

按字段修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
*/
type AlibabaIcbuProductUpdateFieldRequest struct {
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

    // 语种，当前只有english
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

    // 定制信息
    customInfo   *CustomInfo 

    // 商品详情种类，true表示智能编辑，不填默认取商品原来的详情种类
    isSmartEdit   bool 

    // 使用SKU价的时候需要传入这个参数
    useSkuPrice   bool 

    // 混淆商品ID
    productId   string 

}

func NewAlibabaIcbuProductUpdateFieldRequest() *AlibabaIcbuProductUpdateFieldRequest{
    return &AlibabaIcbuProductUpdateFieldRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.update.field"
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductUpdateFieldRequest) SetAttributes(attributes []ProductAttribute) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetAttributes() []ProductAttribute {
    return r.attributes
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetBulkDiscountPrices(bulkDiscountPrices []BulkDiscountPrice) error {
    r.bulkDiscountPrices = bulkDiscountPrices
    r.Set("bulk_discount_prices", bulkDiscountPrices)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetBulkDiscountPrices() []BulkDiscountPrice {
    return r.bulkDiscountPrices
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetCategoryId() int64 {
    return r.categoryId
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetDescription() string {
    return r.description
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetExtraContext() string {
    return r.extraContext
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetGroupId() int64 {
    return r.groupId
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetKeywords(keywords []String) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetKeywords() []String {
    return r.keywords
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetLanguage() string {
    return r.language
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetMainImage(mainImage *MainImage) error {
    r.mainImage = mainImage
    r.Set("main_image", mainImage)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetMainImage() *MainImage {
    return r.mainImage
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetProductSku(productSku *ProductSku) error {
    r.productSku = productSku
    r.Set("product_sku", productSku)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetProductSku() *ProductSku {
    return r.productSku
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetProductType(productType string) error {
    r.productType = productType
    r.Set("product_type", productType)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetProductType() string {
    return r.productType
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetSourcingTrade(sourcingTrade *SourcingTrade) error {
    r.sourcingTrade = sourcingTrade
    r.Set("sourcing_trade", sourcingTrade)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetSourcingTrade() *SourcingTrade {
    return r.sourcingTrade
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetSubject(subject string) error {
    r.subject = subject
    r.Set("subject", subject)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetSubject() string {
    return r.subject
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetWholesaleTrade(wholesaleTrade *WholesaleTrade) error {
    r.wholesaleTrade = wholesaleTrade
    r.Set("wholesale_trade", wholesaleTrade)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetWholesaleTrade() *WholesaleTrade {
    return r.wholesaleTrade
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetMarket(market string) error {
    r.market = market
    r.Set("market", market)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetMarket() string {
    return r.market
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetCustomInfo(customInfo *CustomInfo) error {
    r.customInfo = customInfo
    r.Set("custom_info", customInfo)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetCustomInfo() *CustomInfo {
    return r.customInfo
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetIsSmartEdit(isSmartEdit bool) error {
    r.isSmartEdit = isSmartEdit
    r.Set("is_smart_edit", isSmartEdit)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetIsSmartEdit() bool {
    return r.isSmartEdit
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetUseSkuPrice(useSkuPrice bool) error {
    r.useSkuPrice = useSkuPrice
    r.Set("use_sku_price", useSkuPrice)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetUseSkuPrice() bool {
    return r.useSkuPrice
}

func (r *AlibabaIcbuProductUpdateFieldRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AlibabaIcbuProductUpdateFieldRequest) GetProductId() string {
    return r.productId
}

