package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新产品 APIRequest
taobao.fenxiao.product.update

更新分销平台产品数据，不传更新数据返回失败<br><br/>1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br>
*/
type TaobaoFenxiaoProductUpdateRequest struct {
    model.Params

    // 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）。
    postageType   string 

    // 是否有发票，可选值：false（否）、true（是），默认false。
    haveInvoice   string 

    // 是否有保修，可选值：false（否）、true（是），默认false。
    haveQuarantee   string 

    // 发布状态，可选值：up（上架）、down（下架）、delete（删除），输入非法字符则忽略。
    status   string 

    // 产品ID
    pid   int64 

    // 产品名称，长度不超过60个字节。
    name   string 

    // 采购基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    standardPrice   string 

    // 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    costPrice   string 

    // 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    retailPriceLow   string 

    // 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。
    retailPriceHigh   string 

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

    // 运费模板ID，参考taobao.postages.get。更新时必须指定运费类型为 buyer，否则不更新。
    postageId   int64 

    // 平邮费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
    postageOrdinary   string 

    // 快递费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
    postageFast   string 

    // ems费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。
    postageEms   string 

    // sku id列表，例：1001,1002,1003。如果传入sku_properties将忽略此参数。
    skuIds   string 

    // sku采购价格，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。
    skuCostPrices   string 

    // sku库存，单位元，例："10,20,30"，字段必须和sku_ids或sku_properties保持一致。
    skuQuantitys   string 

    // sku商家编码 ，单位元，例："S1000,S1002,S1003"，字段必须和上面的id或sku_properties保持一致，如果没有可以写成",,"
    skuOuterIds   string 

    // 折扣ID
    discountId   int64 

    // sku采购基准价，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。
    skuStandardPrices   string 

    // sku属性。格式:pid:vid;pid:vid,表示一组属性如:1627207:3232483;1630696:3284570,表示一组:机身颜色:军绿色;手机套餐:一电一充。多组之间用逗号“,”区分。(属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid)<br/>通过此字段可新增和更新sku。若传入此值将忽略sku_ids字段。sku其他字段与此值保持一致。
    skuProperties   string 

    // 根据sku属性删除sku信息。需要按组删除属性。
    skuPropertiesDel   string 

    // 产品是否需要授权isAuthz:yes|no <br/>yes:需要授权 <br/>no:不需要授权
    isAuthz   string 

    // 产品主图图片空间相对路径或绝对路径
    picPath   string 

    // 主图图片，如果pic_path参数不空，则优先使用pic_path，忽略该参数
    image   []byte 

    // 产品属性
    properties   string 

    // 属性别名
    propertyAlias   string 

    // 自定义属性。格式为pid:value;pid:value
    inputProperties   string 

    // 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    dealerCostPrice   string 

    // sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。
    skuDealerCostPrices   string 

    // 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。
    categoryId   int64 

    // 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。
    standardRetailPrice   string 

}

func NewTaobaoFenxiaoProductUpdateRequest() *TaobaoFenxiaoProductUpdateRequest{
    return &TaobaoFenxiaoProductUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoProductUpdateRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.update"
}

func (r TaobaoFenxiaoProductUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoProductUpdateRequest) SetPostageType(postageType string) error {
    r.postageType = postageType
    r.Set("postage_type", postageType)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetPostageType() string {
    return r.postageType
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetHaveInvoice(haveInvoice string) error {
    r.haveInvoice = haveInvoice
    r.Set("have_invoice", haveInvoice)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetHaveInvoice() string {
    return r.haveInvoice
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetHaveQuarantee(haveQuarantee string) error {
    r.haveQuarantee = haveQuarantee
    r.Set("have_quarantee", haveQuarantee)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetHaveQuarantee() string {
    return r.haveQuarantee
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetPid(pid int64) error {
    r.pid = pid
    r.Set("pid", pid)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetPid() int64 {
    return r.pid
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetName() string {
    return r.name
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetStandardPrice(standardPrice string) error {
    r.standardPrice = standardPrice
    r.Set("standard_price", standardPrice)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetStandardPrice() string {
    return r.standardPrice
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetCostPrice(costPrice string) error {
    r.costPrice = costPrice
    r.Set("cost_price", costPrice)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetCostPrice() string {
    return r.costPrice
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetRetailPriceLow(retailPriceLow string) error {
    r.retailPriceLow = retailPriceLow
    r.Set("retail_price_low", retailPriceLow)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetRetailPriceLow() string {
    return r.retailPriceLow
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetRetailPriceHigh(retailPriceHigh string) error {
    r.retailPriceHigh = retailPriceHigh
    r.Set("retail_price_high", retailPriceHigh)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetRetailPriceHigh() string {
    return r.retailPriceHigh
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetQuantity() int64 {
    return r.quantity
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetDesc() string {
    return r.desc
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetProv(prov string) error {
    r.prov = prov
    r.Set("prov", prov)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetProv() string {
    return r.prov
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetCity() string {
    return r.city
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetPostageId(postageId int64) error {
    r.postageId = postageId
    r.Set("postage_id", postageId)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetPostageId() int64 {
    return r.postageId
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetPostageOrdinary(postageOrdinary string) error {
    r.postageOrdinary = postageOrdinary
    r.Set("postage_ordinary", postageOrdinary)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetPostageOrdinary() string {
    return r.postageOrdinary
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetPostageFast(postageFast string) error {
    r.postageFast = postageFast
    r.Set("postage_fast", postageFast)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetPostageFast() string {
    return r.postageFast
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetPostageEms(postageEms string) error {
    r.postageEms = postageEms
    r.Set("postage_ems", postageEms)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetPostageEms() string {
    return r.postageEms
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetSkuIds(skuIds string) error {
    r.skuIds = skuIds
    r.Set("sku_ids", skuIds)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetSkuIds() string {
    return r.skuIds
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetSkuCostPrices(skuCostPrices string) error {
    r.skuCostPrices = skuCostPrices
    r.Set("sku_cost_prices", skuCostPrices)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetSkuCostPrices() string {
    return r.skuCostPrices
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetSkuQuantitys(skuQuantitys string) error {
    r.skuQuantitys = skuQuantitys
    r.Set("sku_quantitys", skuQuantitys)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetSkuQuantitys() string {
    return r.skuQuantitys
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetSkuOuterIds(skuOuterIds string) error {
    r.skuOuterIds = skuOuterIds
    r.Set("sku_outer_ids", skuOuterIds)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetSkuOuterIds() string {
    return r.skuOuterIds
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetDiscountId(discountId int64) error {
    r.discountId = discountId
    r.Set("discount_id", discountId)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetDiscountId() int64 {
    return r.discountId
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetSkuStandardPrices(skuStandardPrices string) error {
    r.skuStandardPrices = skuStandardPrices
    r.Set("sku_standard_prices", skuStandardPrices)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetSkuStandardPrices() string {
    return r.skuStandardPrices
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetSkuProperties(skuProperties string) error {
    r.skuProperties = skuProperties
    r.Set("sku_properties", skuProperties)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetSkuProperties() string {
    return r.skuProperties
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetSkuPropertiesDel(skuPropertiesDel string) error {
    r.skuPropertiesDel = skuPropertiesDel
    r.Set("sku_properties_del", skuPropertiesDel)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetSkuPropertiesDel() string {
    return r.skuPropertiesDel
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetIsAuthz(isAuthz string) error {
    r.isAuthz = isAuthz
    r.Set("is_authz", isAuthz)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetIsAuthz() string {
    return r.isAuthz
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetPicPath(picPath string) error {
    r.picPath = picPath
    r.Set("pic_path", picPath)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetPicPath() string {
    return r.picPath
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetImage(image []byte) error {
    r.image = image
    r.Set("image", image)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetImage() []byte {
    return r.image
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetProperties() string {
    return r.properties
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetPropertyAlias(propertyAlias string) error {
    r.propertyAlias = propertyAlias
    r.Set("property_alias", propertyAlias)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetPropertyAlias() string {
    return r.propertyAlias
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetInputProperties(inputProperties string) error {
    r.inputProperties = inputProperties
    r.Set("input_properties", inputProperties)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetInputProperties() string {
    return r.inputProperties
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetDealerCostPrice(dealerCostPrice string) error {
    r.dealerCostPrice = dealerCostPrice
    r.Set("dealer_cost_price", dealerCostPrice)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetDealerCostPrice() string {
    return r.dealerCostPrice
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetSkuDealerCostPrices(skuDealerCostPrices string) error {
    r.skuDealerCostPrices = skuDealerCostPrices
    r.Set("sku_dealer_cost_prices", skuDealerCostPrices)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetSkuDealerCostPrices() string {
    return r.skuDealerCostPrices
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetCategoryId() int64 {
    return r.categoryId
}

func (r *TaobaoFenxiaoProductUpdateRequest) SetStandardRetailPrice(standardRetailPrice string) error {
    r.standardRetailPrice = standardRetailPrice
    r.Set("standard_retail_price", standardRetailPrice)
    return nil
}

func (r TaobaoFenxiaoProductUpdateRequest) GetStandardRetailPrice() string {
    return r.standardRetailPrice
}

