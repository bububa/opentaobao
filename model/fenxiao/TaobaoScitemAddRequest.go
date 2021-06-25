package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发布后端商品 APIRequest
taobao.scitem.add

发布后端商品
*/
type TaobaoScitemAddRequest struct {
    model.Params

    // 商品名称
    itemName   string 

    // 商家编码
    outerCode   string 

    // 0.普通供应链商品 1.供应链组合主商品
    itemType   int64 

    // 商品属性格式是  p1:v1,p2:v2,p3:v3
    properties   string 

    // 条形码
    barCode   string 

    // 仓储商编码
    wmsCode   string 

    // 是否易碎 0：不是  1：是
    isFriable   int64 

    // 是否危险 0：不是  1：是
    isDangerous   int64 

    // 是否是贵重品 0:不是 1：是
    isCostly   int64 

    // 是否保质期：0:不是 1：是
    isWarranty   int64 

    // 重量 单位：g
    weight   int64 

    // 长度 单位：mm
    length   int64 

    // 宽 单位：mm
    width   int64 

    // 高 单位：mm
    height   int64 

    // 体积：立方厘米
    volume   int64 

    // 价格 单位：分
    price   int64 

    // remark
    remark   string 

    // 0:液体，1：粉体，2：固体
    matterStatus   int64 

    // 品牌id
    brandId   int64 

    // brand_Name
    brandName   string 

    // spuId或是cspuid
    spuId   int64 

    // 1表示区域销售，0或是空是非区域销售
    isAreaSale   int64 

}

func NewTaobaoScitemAddRequest() *TaobaoScitemAddRequest{
    return &TaobaoScitemAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoScitemAddRequest) GetApiMethodName() string {
    return "taobao.scitem.add"
}

func (r TaobaoScitemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoScitemAddRequest) SetItemName(itemName string) error {
    r.itemName = itemName
    r.Set("item_name", itemName)
    return nil
}

func (r TaobaoScitemAddRequest) GetItemName() string {
    return r.itemName
}

func (r *TaobaoScitemAddRequest) SetOuterCode(outerCode string) error {
    r.outerCode = outerCode
    r.Set("outer_code", outerCode)
    return nil
}

func (r TaobaoScitemAddRequest) GetOuterCode() string {
    return r.outerCode
}

func (r *TaobaoScitemAddRequest) SetItemType(itemType int64) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

func (r TaobaoScitemAddRequest) GetItemType() int64 {
    return r.itemType
}

func (r *TaobaoScitemAddRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

func (r TaobaoScitemAddRequest) GetProperties() string {
    return r.properties
}

func (r *TaobaoScitemAddRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

func (r TaobaoScitemAddRequest) GetBarCode() string {
    return r.barCode
}

func (r *TaobaoScitemAddRequest) SetWmsCode(wmsCode string) error {
    r.wmsCode = wmsCode
    r.Set("wms_code", wmsCode)
    return nil
}

func (r TaobaoScitemAddRequest) GetWmsCode() string {
    return r.wmsCode
}

func (r *TaobaoScitemAddRequest) SetIsFriable(isFriable int64) error {
    r.isFriable = isFriable
    r.Set("is_friable", isFriable)
    return nil
}

func (r TaobaoScitemAddRequest) GetIsFriable() int64 {
    return r.isFriable
}

func (r *TaobaoScitemAddRequest) SetIsDangerous(isDangerous int64) error {
    r.isDangerous = isDangerous
    r.Set("is_dangerous", isDangerous)
    return nil
}

func (r TaobaoScitemAddRequest) GetIsDangerous() int64 {
    return r.isDangerous
}

func (r *TaobaoScitemAddRequest) SetIsCostly(isCostly int64) error {
    r.isCostly = isCostly
    r.Set("is_costly", isCostly)
    return nil
}

func (r TaobaoScitemAddRequest) GetIsCostly() int64 {
    return r.isCostly
}

func (r *TaobaoScitemAddRequest) SetIsWarranty(isWarranty int64) error {
    r.isWarranty = isWarranty
    r.Set("is_warranty", isWarranty)
    return nil
}

func (r TaobaoScitemAddRequest) GetIsWarranty() int64 {
    return r.isWarranty
}

func (r *TaobaoScitemAddRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

func (r TaobaoScitemAddRequest) GetWeight() int64 {
    return r.weight
}

func (r *TaobaoScitemAddRequest) SetLength(length int64) error {
    r.length = length
    r.Set("length", length)
    return nil
}

func (r TaobaoScitemAddRequest) GetLength() int64 {
    return r.length
}

func (r *TaobaoScitemAddRequest) SetWidth(width int64) error {
    r.width = width
    r.Set("width", width)
    return nil
}

func (r TaobaoScitemAddRequest) GetWidth() int64 {
    return r.width
}

func (r *TaobaoScitemAddRequest) SetHeight(height int64) error {
    r.height = height
    r.Set("height", height)
    return nil
}

func (r TaobaoScitemAddRequest) GetHeight() int64 {
    return r.height
}

func (r *TaobaoScitemAddRequest) SetVolume(volume int64) error {
    r.volume = volume
    r.Set("volume", volume)
    return nil
}

func (r TaobaoScitemAddRequest) GetVolume() int64 {
    return r.volume
}

func (r *TaobaoScitemAddRequest) SetPrice(price int64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

func (r TaobaoScitemAddRequest) GetPrice() int64 {
    return r.price
}

func (r *TaobaoScitemAddRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoScitemAddRequest) GetRemark() string {
    return r.remark
}

func (r *TaobaoScitemAddRequest) SetMatterStatus(matterStatus int64) error {
    r.matterStatus = matterStatus
    r.Set("matter_status", matterStatus)
    return nil
}

func (r TaobaoScitemAddRequest) GetMatterStatus() int64 {
    return r.matterStatus
}

func (r *TaobaoScitemAddRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

func (r TaobaoScitemAddRequest) GetBrandId() int64 {
    return r.brandId
}

func (r *TaobaoScitemAddRequest) SetBrandName(brandName string) error {
    r.brandName = brandName
    r.Set("brand_name", brandName)
    return nil
}

func (r TaobaoScitemAddRequest) GetBrandName() string {
    return r.brandName
}

func (r *TaobaoScitemAddRequest) SetSpuId(spuId int64) error {
    r.spuId = spuId
    r.Set("spu_id", spuId)
    return nil
}

func (r TaobaoScitemAddRequest) GetSpuId() int64 {
    return r.spuId
}

func (r *TaobaoScitemAddRequest) SetIsAreaSale(isAreaSale int64) error {
    r.isAreaSale = isAreaSale
    r.Set("is_area_sale", isAreaSale)
    return nil
}

func (r TaobaoScitemAddRequest) GetIsAreaSale() int64 {
    return r.isAreaSale
}

