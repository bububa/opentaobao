package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加产品 API请求
taobao.fenxiao.product.add

添加分销平台产品数据。业务逻辑与分销系统前台页面一致。<br/><br/>    * 产品图片默认为空<br/>    * 产品发布后默认为下架状态
*/
type TaobaoFenxiaoProductAddRequest struct {
    model.Params
    // 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。
    postageType   string
    // 是否有发票，可选值：false（否）、true（是），默认false。
    haveInvoice   string
    // 是否有保修，可选值：false（否）、true（是），默认false。
    haveQuarantee   string
    // 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
    tradeType   string
    // 产品名称，长度不超过60个字节。
    name   string
    // 产品线ID
    productcatId   int64
    // 采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    standardPrice   string
    // 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    costPrice   string
    // 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    retailPriceLow   string
    // 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
    retailPriceHigh   string
    // 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
    categoryId   int64
    // 商家编码，长度不能超过60个字节。
    outerId   string
    // 产品库存必须是1到999999。
    quantity   int64
    // 产品描述，长度为5到25000字符。
    desc   string
    // 所在地：省，例：“浙江”
    prov   string
    // 所在地：市，例：“杭州”
    city   string
    // 运费模板ID，参考taobao.postages.get。
    postageId   int64
    // 平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
    postageOrdinary   string
    // 快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
    postageFast   string
    // ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。
    postageEms   string
    // 折扣ID
    discountId   int64
    // 添加产品时，添加入参isAuthz:yes|no <br/>yes:需要授权 <br/>no:不需要授权 <br/>默认是需要授权
    isAuthz   string
    // 产品主图图片空间相对路径或绝对路径
    picPath   string
    // 产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片
    image   []*model.File
    // 产品属性，格式为pid:vid;pid:vid
    properties   string
    // 属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名）
    propertyAlias   string
    // 自定义属性。格式为pid:value;pid:value
    inputProperties   string
    // sku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    skuStandardPrices   string
    // sku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    skuCostPrices   string
    // sku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    skuOuterIds   string
    // sku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    skuQuantitys   string
    // sku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
    skuProperties   string
    // 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    dealerCostPrice   string
    // sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
    skuDealerCostPrices   string
    // 导入的商品ID
    itemId   int64
    // 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    standardRetailPrice   string
    // 产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传
    spuId   int64
}

// 初始化TaobaoFenxiaoProductAddRequest对象
func NewTaobaoFenxiaoProductAddRequest() *TaobaoFenxiaoProductAddRequest{
    return &TaobaoFenxiaoProductAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductAddRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PostageType Setter
// 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。
func (r *TaobaoFenxiaoProductAddRequest) SetPostageType(postageType string) error {
    r.postageType = postageType
    r.Set("postage_type", postageType)
    return nil
}

// PostageType Getter
func (r TaobaoFenxiaoProductAddRequest) GetPostageType() string {
    return r.postageType
}
// HaveInvoice Setter
// 是否有发票，可选值：false（否）、true（是），默认false。
func (r *TaobaoFenxiaoProductAddRequest) SetHaveInvoice(haveInvoice string) error {
    r.haveInvoice = haveInvoice
    r.Set("have_invoice", haveInvoice)
    return nil
}

// HaveInvoice Getter
func (r TaobaoFenxiaoProductAddRequest) GetHaveInvoice() string {
    return r.haveInvoice
}
// HaveQuarantee Setter
// 是否有保修，可选值：false（否）、true（是），默认false。
func (r *TaobaoFenxiaoProductAddRequest) SetHaveQuarantee(haveQuarantee string) error {
    r.haveQuarantee = haveQuarantee
    r.Set("have_quarantee", haveQuarantee)
    return nil
}

// HaveQuarantee Getter
func (r TaobaoFenxiaoProductAddRequest) GetHaveQuarantee() string {
    return r.haveQuarantee
}
// TradeType Setter
// 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
func (r *TaobaoFenxiaoProductAddRequest) SetTradeType(tradeType string) error {
    r.tradeType = tradeType
    r.Set("trade_type", tradeType)
    return nil
}

// TradeType Getter
func (r TaobaoFenxiaoProductAddRequest) GetTradeType() string {
    return r.tradeType
}
// Name Setter
// 产品名称，长度不超过60个字节。
func (r *TaobaoFenxiaoProductAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoFenxiaoProductAddRequest) GetName() string {
    return r.name
}
// ProductcatId Setter
// 产品线ID
func (r *TaobaoFenxiaoProductAddRequest) SetProductcatId(productcatId int64) error {
    r.productcatId = productcatId
    r.Set("productcat_id", productcatId)
    return nil
}

// ProductcatId Getter
func (r TaobaoFenxiaoProductAddRequest) GetProductcatId() int64 {
    return r.productcatId
}
// StandardPrice Setter
// 采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetStandardPrice(standardPrice string) error {
    r.standardPrice = standardPrice
    r.Set("standard_price", standardPrice)
    return nil
}

// StandardPrice Getter
func (r TaobaoFenxiaoProductAddRequest) GetStandardPrice() string {
    return r.standardPrice
}
// CostPrice Setter
// 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetCostPrice(costPrice string) error {
    r.costPrice = costPrice
    r.Set("cost_price", costPrice)
    return nil
}

// CostPrice Getter
func (r TaobaoFenxiaoProductAddRequest) GetCostPrice() string {
    return r.costPrice
}
// RetailPriceLow Setter
// 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetRetailPriceLow(retailPriceLow string) error {
    r.retailPriceLow = retailPriceLow
    r.Set("retail_price_low", retailPriceLow)
    return nil
}

// RetailPriceLow Getter
func (r TaobaoFenxiaoProductAddRequest) GetRetailPriceLow() string {
    return r.retailPriceLow
}
// RetailPriceHigh Setter
// 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
func (r *TaobaoFenxiaoProductAddRequest) SetRetailPriceHigh(retailPriceHigh string) error {
    r.retailPriceHigh = retailPriceHigh
    r.Set("retail_price_high", retailPriceHigh)
    return nil
}

// RetailPriceHigh Getter
func (r TaobaoFenxiaoProductAddRequest) GetRetailPriceHigh() string {
    return r.retailPriceHigh
}
// CategoryId Setter
// 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
func (r *TaobaoFenxiaoProductAddRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r TaobaoFenxiaoProductAddRequest) GetCategoryId() int64 {
    return r.categoryId
}
// OuterId Setter
// 商家编码，长度不能超过60个字节。
func (r *TaobaoFenxiaoProductAddRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoFenxiaoProductAddRequest) GetOuterId() string {
    return r.outerId
}
// Quantity Setter
// 产品库存必须是1到999999。
func (r *TaobaoFenxiaoProductAddRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

// Quantity Getter
func (r TaobaoFenxiaoProductAddRequest) GetQuantity() int64 {
    return r.quantity
}
// Desc Setter
// 产品描述，长度为5到25000字符。
func (r *TaobaoFenxiaoProductAddRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

// Desc Getter
func (r TaobaoFenxiaoProductAddRequest) GetDesc() string {
    return r.desc
}
// Prov Setter
// 所在地：省，例：“浙江”
func (r *TaobaoFenxiaoProductAddRequest) SetProv(prov string) error {
    r.prov = prov
    r.Set("prov", prov)
    return nil
}

// Prov Getter
func (r TaobaoFenxiaoProductAddRequest) GetProv() string {
    return r.prov
}
// City Setter
// 所在地：市，例：“杭州”
func (r *TaobaoFenxiaoProductAddRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r TaobaoFenxiaoProductAddRequest) GetCity() string {
    return r.city
}
// PostageId Setter
// 运费模板ID，参考taobao.postages.get。
func (r *TaobaoFenxiaoProductAddRequest) SetPostageId(postageId int64) error {
    r.postageId = postageId
    r.Set("postage_id", postageId)
    return nil
}

// PostageId Getter
func (r TaobaoFenxiaoProductAddRequest) GetPostageId() int64 {
    return r.postageId
}
// PostageOrdinary Setter
// 平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetPostageOrdinary(postageOrdinary string) error {
    r.postageOrdinary = postageOrdinary
    r.Set("postage_ordinary", postageOrdinary)
    return nil
}

// PostageOrdinary Getter
func (r TaobaoFenxiaoProductAddRequest) GetPostageOrdinary() string {
    return r.postageOrdinary
}
// PostageFast Setter
// 快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetPostageFast(postageFast string) error {
    r.postageFast = postageFast
    r.Set("postage_fast", postageFast)
    return nil
}

// PostageFast Getter
func (r TaobaoFenxiaoProductAddRequest) GetPostageFast() string {
    return r.postageFast
}
// PostageEms Setter
// ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetPostageEms(postageEms string) error {
    r.postageEms = postageEms
    r.Set("postage_ems", postageEms)
    return nil
}

// PostageEms Getter
func (r TaobaoFenxiaoProductAddRequest) GetPostageEms() string {
    return r.postageEms
}
// DiscountId Setter
// 折扣ID
func (r *TaobaoFenxiaoProductAddRequest) SetDiscountId(discountId int64) error {
    r.discountId = discountId
    r.Set("discount_id", discountId)
    return nil
}

// DiscountId Getter
func (r TaobaoFenxiaoProductAddRequest) GetDiscountId() int64 {
    return r.discountId
}
// IsAuthz Setter
// 添加产品时，添加入参isAuthz:yes|no <br/>yes:需要授权 <br/>no:不需要授权 <br/>默认是需要授权
func (r *TaobaoFenxiaoProductAddRequest) SetIsAuthz(isAuthz string) error {
    r.isAuthz = isAuthz
    r.Set("is_authz", isAuthz)
    return nil
}

// IsAuthz Getter
func (r TaobaoFenxiaoProductAddRequest) GetIsAuthz() string {
    return r.isAuthz
}
// PicPath Setter
// 产品主图图片空间相对路径或绝对路径
func (r *TaobaoFenxiaoProductAddRequest) SetPicPath(picPath string) error {
    r.picPath = picPath
    r.Set("pic_path", picPath)
    return nil
}

// PicPath Getter
func (r TaobaoFenxiaoProductAddRequest) GetPicPath() string {
    return r.picPath
}
// Image Setter
// 产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片
func (r *TaobaoFenxiaoProductAddRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

// Image Getter
func (r TaobaoFenxiaoProductAddRequest) GetImage() []*model.File {
    return r.image
}
// Properties Setter
// 产品属性，格式为pid:vid;pid:vid
func (r *TaobaoFenxiaoProductAddRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

// Properties Getter
func (r TaobaoFenxiaoProductAddRequest) GetProperties() string {
    return r.properties
}
// PropertyAlias Setter
// 属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名）
func (r *TaobaoFenxiaoProductAddRequest) SetPropertyAlias(propertyAlias string) error {
    r.propertyAlias = propertyAlias
    r.Set("property_alias", propertyAlias)
    return nil
}

// PropertyAlias Getter
func (r TaobaoFenxiaoProductAddRequest) GetPropertyAlias() string {
    return r.propertyAlias
}
// InputProperties Setter
// 自定义属性。格式为pid:value;pid:value
func (r *TaobaoFenxiaoProductAddRequest) SetInputProperties(inputProperties string) error {
    r.inputProperties = inputProperties
    r.Set("input_properties", inputProperties)
    return nil
}

// InputProperties Getter
func (r TaobaoFenxiaoProductAddRequest) GetInputProperties() string {
    return r.inputProperties
}
// SkuStandardPrices Setter
// sku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddRequest) SetSkuStandardPrices(skuStandardPrices string) error {
    r.skuStandardPrices = skuStandardPrices
    r.Set("sku_standard_prices", skuStandardPrices)
    return nil
}

// SkuStandardPrices Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuStandardPrices() string {
    return r.skuStandardPrices
}
// SkuCostPrices Setter
// sku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddRequest) SetSkuCostPrices(skuCostPrices string) error {
    r.skuCostPrices = skuCostPrices
    r.Set("sku_cost_prices", skuCostPrices)
    return nil
}

// SkuCostPrices Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuCostPrices() string {
    return r.skuCostPrices
}
// SkuOuterIds Setter
// sku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddRequest) SetSkuOuterIds(skuOuterIds string) error {
    r.skuOuterIds = skuOuterIds
    r.Set("sku_outer_ids", skuOuterIds)
    return nil
}

// SkuOuterIds Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuOuterIds() string {
    return r.skuOuterIds
}
// SkuQuantitys Setter
// sku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddRequest) SetSkuQuantitys(skuQuantitys string) error {
    r.skuQuantitys = skuQuantitys
    r.Set("sku_quantitys", skuQuantitys)
    return nil
}

// SkuQuantitys Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuQuantitys() string {
    return r.skuQuantitys
}
// SkuProperties Setter
// sku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序
func (r *TaobaoFenxiaoProductAddRequest) SetSkuProperties(skuProperties string) error {
    r.skuProperties = skuProperties
    r.Set("sku_properties", skuProperties)
    return nil
}

// SkuProperties Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuProperties() string {
    return r.skuProperties
}
// DealerCostPrice Setter
// 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetDealerCostPrice(dealerCostPrice string) error {
    r.dealerCostPrice = dealerCostPrice
    r.Set("dealer_cost_price", dealerCostPrice)
    return nil
}

// DealerCostPrice Getter
func (r TaobaoFenxiaoProductAddRequest) GetDealerCostPrice() string {
    return r.dealerCostPrice
}
// SkuDealerCostPrices Setter
// sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetSkuDealerCostPrices(skuDealerCostPrices string) error {
    r.skuDealerCostPrices = skuDealerCostPrices
    r.Set("sku_dealer_cost_prices", skuDealerCostPrices)
    return nil
}

// SkuDealerCostPrices Getter
func (r TaobaoFenxiaoProductAddRequest) GetSkuDealerCostPrices() string {
    return r.skuDealerCostPrices
}
// ItemId Setter
// 导入的商品ID
func (r *TaobaoFenxiaoProductAddRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoFenxiaoProductAddRequest) GetItemId() int64 {
    return r.itemId
}
// StandardRetailPrice Setter
// 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
func (r *TaobaoFenxiaoProductAddRequest) SetStandardRetailPrice(standardRetailPrice string) error {
    r.standardRetailPrice = standardRetailPrice
    r.Set("standard_retail_price", standardRetailPrice)
    return nil
}

// StandardRetailPrice Getter
func (r TaobaoFenxiaoProductAddRequest) GetStandardRetailPrice() string {
    return r.standardRetailPrice
}
// SpuId Setter
// 产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传
func (r *TaobaoFenxiaoProductAddRequest) SetSpuId(spuId int64) error {
    r.spuId = spuId
    r.Set("spu_id", spuId)
    return nil
}

// SpuId Getter
func (r TaobaoFenxiaoProductAddRequest) GetSpuId() int64 {
    return r.spuId
}
