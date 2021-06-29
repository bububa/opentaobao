package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发布后端商品 API请求
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

// 初始化TaobaoScitemAddRequest对象
func NewTaobaoScitemAddRequest() *TaobaoScitemAddRequest{
    return &TaobaoScitemAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoScitemAddRequest) GetApiMethodName() string {
    return "taobao.scitem.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoScitemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemName Setter
// 商品名称
func (r *TaobaoScitemAddRequest) SetItemName(itemName string) error {
    r.itemName = itemName
    r.Set("item_name", itemName)
    return nil
}

// ItemName Getter
func (r TaobaoScitemAddRequest) GetItemName() string {
    return r.itemName
}
// OuterCode Setter
// 商家编码
func (r *TaobaoScitemAddRequest) SetOuterCode(outerCode string) error {
    r.outerCode = outerCode
    r.Set("outer_code", outerCode)
    return nil
}

// OuterCode Getter
func (r TaobaoScitemAddRequest) GetOuterCode() string {
    return r.outerCode
}
// ItemType Setter
// 0.普通供应链商品 1.供应链组合主商品
func (r *TaobaoScitemAddRequest) SetItemType(itemType int64) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

// ItemType Getter
func (r TaobaoScitemAddRequest) GetItemType() int64 {
    return r.itemType
}
// Properties Setter
// 商品属性格式是  p1:v1,p2:v2,p3:v3
func (r *TaobaoScitemAddRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

// Properties Getter
func (r TaobaoScitemAddRequest) GetProperties() string {
    return r.properties
}
// BarCode Setter
// 条形码
func (r *TaobaoScitemAddRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

// BarCode Getter
func (r TaobaoScitemAddRequest) GetBarCode() string {
    return r.barCode
}
// WmsCode Setter
// 仓储商编码
func (r *TaobaoScitemAddRequest) SetWmsCode(wmsCode string) error {
    r.wmsCode = wmsCode
    r.Set("wms_code", wmsCode)
    return nil
}

// WmsCode Getter
func (r TaobaoScitemAddRequest) GetWmsCode() string {
    return r.wmsCode
}
// IsFriable Setter
// 是否易碎 0：不是  1：是
func (r *TaobaoScitemAddRequest) SetIsFriable(isFriable int64) error {
    r.isFriable = isFriable
    r.Set("is_friable", isFriable)
    return nil
}

// IsFriable Getter
func (r TaobaoScitemAddRequest) GetIsFriable() int64 {
    return r.isFriable
}
// IsDangerous Setter
// 是否危险 0：不是  1：是
func (r *TaobaoScitemAddRequest) SetIsDangerous(isDangerous int64) error {
    r.isDangerous = isDangerous
    r.Set("is_dangerous", isDangerous)
    return nil
}

// IsDangerous Getter
func (r TaobaoScitemAddRequest) GetIsDangerous() int64 {
    return r.isDangerous
}
// IsCostly Setter
// 是否是贵重品 0:不是 1：是
func (r *TaobaoScitemAddRequest) SetIsCostly(isCostly int64) error {
    r.isCostly = isCostly
    r.Set("is_costly", isCostly)
    return nil
}

// IsCostly Getter
func (r TaobaoScitemAddRequest) GetIsCostly() int64 {
    return r.isCostly
}
// IsWarranty Setter
// 是否保质期：0:不是 1：是
func (r *TaobaoScitemAddRequest) SetIsWarranty(isWarranty int64) error {
    r.isWarranty = isWarranty
    r.Set("is_warranty", isWarranty)
    return nil
}

// IsWarranty Getter
func (r TaobaoScitemAddRequest) GetIsWarranty() int64 {
    return r.isWarranty
}
// Weight Setter
// 重量 单位：g
func (r *TaobaoScitemAddRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

// Weight Getter
func (r TaobaoScitemAddRequest) GetWeight() int64 {
    return r.weight
}
// Length Setter
// 长度 单位：mm
func (r *TaobaoScitemAddRequest) SetLength(length int64) error {
    r.length = length
    r.Set("length", length)
    return nil
}

// Length Getter
func (r TaobaoScitemAddRequest) GetLength() int64 {
    return r.length
}
// Width Setter
// 宽 单位：mm
func (r *TaobaoScitemAddRequest) SetWidth(width int64) error {
    r.width = width
    r.Set("width", width)
    return nil
}

// Width Getter
func (r TaobaoScitemAddRequest) GetWidth() int64 {
    return r.width
}
// Height Setter
// 高 单位：mm
func (r *TaobaoScitemAddRequest) SetHeight(height int64) error {
    r.height = height
    r.Set("height", height)
    return nil
}

// Height Getter
func (r TaobaoScitemAddRequest) GetHeight() int64 {
    return r.height
}
// Volume Setter
// 体积：立方厘米
func (r *TaobaoScitemAddRequest) SetVolume(volume int64) error {
    r.volume = volume
    r.Set("volume", volume)
    return nil
}

// Volume Getter
func (r TaobaoScitemAddRequest) GetVolume() int64 {
    return r.volume
}
// Price Setter
// 价格 单位：分
func (r *TaobaoScitemAddRequest) SetPrice(price int64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

// Price Getter
func (r TaobaoScitemAddRequest) GetPrice() int64 {
    return r.price
}
// Remark Setter
// remark
func (r *TaobaoScitemAddRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r TaobaoScitemAddRequest) GetRemark() string {
    return r.remark
}
// MatterStatus Setter
// 0:液体，1：粉体，2：固体
func (r *TaobaoScitemAddRequest) SetMatterStatus(matterStatus int64) error {
    r.matterStatus = matterStatus
    r.Set("matter_status", matterStatus)
    return nil
}

// MatterStatus Getter
func (r TaobaoScitemAddRequest) GetMatterStatus() int64 {
    return r.matterStatus
}
// BrandId Setter
// 品牌id
func (r *TaobaoScitemAddRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

// BrandId Getter
func (r TaobaoScitemAddRequest) GetBrandId() int64 {
    return r.brandId
}
// BrandName Setter
// brand_Name
func (r *TaobaoScitemAddRequest) SetBrandName(brandName string) error {
    r.brandName = brandName
    r.Set("brand_name", brandName)
    return nil
}

// BrandName Getter
func (r TaobaoScitemAddRequest) GetBrandName() string {
    return r.brandName
}
// SpuId Setter
// spuId或是cspuid
func (r *TaobaoScitemAddRequest) SetSpuId(spuId int64) error {
    r.spuId = spuId
    r.Set("spu_id", spuId)
    return nil
}

// SpuId Getter
func (r TaobaoScitemAddRequest) GetSpuId() int64 {
    return r.spuId
}
// IsAreaSale Setter
// 1表示区域销售，0或是空是非区域销售
func (r *TaobaoScitemAddRequest) SetIsAreaSale(isAreaSale int64) error {
    r.isAreaSale = isAreaSale
    r.Set("is_area_sale", isAreaSale)
    return nil
}

// IsAreaSale Getter
func (r TaobaoScitemAddRequest) GetIsAreaSale() int64 {
    return r.isAreaSale
}
