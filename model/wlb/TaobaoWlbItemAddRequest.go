package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加单个物流宝商品 APIRequest
taobao.wlb.item.add

添加物流宝商品，支持物流宝子商品和属性添加
*/
type TaobaoWlbItemAddRequest struct {
    model.Params

    // 商品名称
    name   string 

    // 商品标题
    title   string 

    // 商品编码
    itemCode   string 

    // 商品备注
    remark   string 

    // NORMAL--普通商品<br/>COMBINE--组合商品<br/>DISTRIBUTION--分销
    type   string 

    // 是否sku
    isSku   string 

    // 属性名列表,目前支持的属性：<br/>毛重:GWeight	<br/>净重:Nweight<br/>皮重: Tweight<br/>自定义属性：<br/>paramkey1<br/>paramkey2<br/>paramkey3<br/>paramkey4
    proNameList   string 

    // 属性值列表：<br/>10,8
    proValueList   string 

    // 是否易碎品
    isFriable   bool 

    // 是否危险品
    isDangerous   bool 

    // 商品颜色
    color   string 

    // 商品重量，单位G
    weight   int64 

    // 商品长度，单位毫米
    length   int64 

    // 商品宽度，单位毫米
    width   int64 

    // 商品高度，单位毫米
    height   int64 

    // 商品体积，单位立方厘米
    volume   int64 

    // 货类
    goodsCat   string 

    // 计价货类
    pricingCat   string 

    // 商品包装材料类型
    packageMaterial   string 

    // 商品价格，单位：分
    price   int64 

    // 是否支持批次
    supportBatch   bool 

}

func NewTaobaoWlbItemAddRequest() *TaobaoWlbItemAddRequest{
    return &TaobaoWlbItemAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbItemAddRequest) GetApiMethodName() string {
    return "taobao.wlb.item.add"
}

func (r TaobaoWlbItemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbItemAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetName() string {
    return r.name
}

func (r *TaobaoWlbItemAddRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoWlbItemAddRequest) SetItemCode(itemCode string) error {
    r.itemCode = itemCode
    r.Set("item_code", itemCode)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetItemCode() string {
    return r.itemCode
}

func (r *TaobaoWlbItemAddRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetRemark() string {
    return r.remark
}

func (r *TaobaoWlbItemAddRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetType() string {
    return r.type
}

func (r *TaobaoWlbItemAddRequest) SetIsSku(isSku string) error {
    r.isSku = isSku
    r.Set("is_sku", isSku)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetIsSku() string {
    return r.isSku
}

func (r *TaobaoWlbItemAddRequest) SetProNameList(proNameList string) error {
    r.proNameList = proNameList
    r.Set("pro_name_list", proNameList)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetProNameList() string {
    return r.proNameList
}

func (r *TaobaoWlbItemAddRequest) SetProValueList(proValueList string) error {
    r.proValueList = proValueList
    r.Set("pro_value_list", proValueList)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetProValueList() string {
    return r.proValueList
}

func (r *TaobaoWlbItemAddRequest) SetIsFriable(isFriable bool) error {
    r.isFriable = isFriable
    r.Set("is_friable", isFriable)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetIsFriable() bool {
    return r.isFriable
}

func (r *TaobaoWlbItemAddRequest) SetIsDangerous(isDangerous bool) error {
    r.isDangerous = isDangerous
    r.Set("is_dangerous", isDangerous)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetIsDangerous() bool {
    return r.isDangerous
}

func (r *TaobaoWlbItemAddRequest) SetColor(color string) error {
    r.color = color
    r.Set("color", color)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetColor() string {
    return r.color
}

func (r *TaobaoWlbItemAddRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetWeight() int64 {
    return r.weight
}

func (r *TaobaoWlbItemAddRequest) SetLength(length int64) error {
    r.length = length
    r.Set("length", length)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetLength() int64 {
    return r.length
}

func (r *TaobaoWlbItemAddRequest) SetWidth(width int64) error {
    r.width = width
    r.Set("width", width)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetWidth() int64 {
    return r.width
}

func (r *TaobaoWlbItemAddRequest) SetHeight(height int64) error {
    r.height = height
    r.Set("height", height)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetHeight() int64 {
    return r.height
}

func (r *TaobaoWlbItemAddRequest) SetVolume(volume int64) error {
    r.volume = volume
    r.Set("volume", volume)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetVolume() int64 {
    return r.volume
}

func (r *TaobaoWlbItemAddRequest) SetGoodsCat(goodsCat string) error {
    r.goodsCat = goodsCat
    r.Set("goods_cat", goodsCat)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetGoodsCat() string {
    return r.goodsCat
}

func (r *TaobaoWlbItemAddRequest) SetPricingCat(pricingCat string) error {
    r.pricingCat = pricingCat
    r.Set("pricing_cat", pricingCat)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetPricingCat() string {
    return r.pricingCat
}

func (r *TaobaoWlbItemAddRequest) SetPackageMaterial(packageMaterial string) error {
    r.packageMaterial = packageMaterial
    r.Set("package_material", packageMaterial)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetPackageMaterial() string {
    return r.packageMaterial
}

func (r *TaobaoWlbItemAddRequest) SetPrice(price int64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetPrice() int64 {
    return r.price
}

func (r *TaobaoWlbItemAddRequest) SetSupportBatch(supportBatch bool) error {
    r.supportBatch = supportBatch
    r.Set("support_batch", supportBatch)
    return nil
}

func (r TaobaoWlbItemAddRequest) GetSupportBatch() bool {
    return r.supportBatch
}

