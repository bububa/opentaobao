package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID或商家编码修改后端商品 API请求
taobao.scitem.update

根据商品ID或商家编码修改后端商品
*/
type TaobaoScitemUpdateRequest struct {
    model.Params
    // 后端商品ID，跟outer_code必须指定一个
    itemId   int64
    // 商家编码，跟item_id必须指定一个
    outerCode   string
    // 商品名称
    itemName   string
    // 0.普通供应链商品 1.供应链组合主商品
    itemType   int64
    // 需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3
    updateProperties   string
    // 条形码
    barCode   string
    // 仓储商编码
    wmsCode   string
    // 是否易碎 0：不是  1：是
    isFriable   int64
    // 是否危险 0：不是  0：是
    isDangerous   int64
    // 是否是贵重品 0:不是 1：是
    isCostly   int64
    // 是否保质期：0:不是 1：是
    isWarranty   int64
    // weight
    weight   int64
    // 长度 单位：mm
    length   int64
    // 宽 单位：mm
    width   int64
    // 高 单位：mm
    height   int64
    // 体积：立方厘米
    volume   int64
    // price
    price   int64
    // remark
    remark   string
    // 0:液体，1：粉体，2：固体
    matterStatus   int64
    // 品牌id
    brandId   int64
    // brand_Name
    brandName   string
    // 淘宝SKU产品级编码CSPU ID
    spuId   int64
    // 移除商品属性P列表,P由系统分配：p1；p2
    removeProperties   string
    // 1表示区域销售，0或是空是非区域销售
    isAreaSale   int64
}

// 初始化TaobaoScitemUpdateRequest对象
func NewTaobaoScitemUpdateRequest() *TaobaoScitemUpdateRequest{
    return &TaobaoScitemUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoScitemUpdateRequest) GetApiMethodName() string {
    return "taobao.scitem.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoScitemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 后端商品ID，跟outer_code必须指定一个
func (r *TaobaoScitemUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoScitemUpdateRequest) GetItemId() int64 {
    return r.itemId
}
// OuterCode Setter
// 商家编码，跟item_id必须指定一个
func (r *TaobaoScitemUpdateRequest) SetOuterCode(outerCode string) error {
    r.outerCode = outerCode
    r.Set("outer_code", outerCode)
    return nil
}

// OuterCode Getter
func (r TaobaoScitemUpdateRequest) GetOuterCode() string {
    return r.outerCode
}
// ItemName Setter
// 商品名称
func (r *TaobaoScitemUpdateRequest) SetItemName(itemName string) error {
    r.itemName = itemName
    r.Set("item_name", itemName)
    return nil
}

// ItemName Getter
func (r TaobaoScitemUpdateRequest) GetItemName() string {
    return r.itemName
}
// ItemType Setter
// 0.普通供应链商品 1.供应链组合主商品
func (r *TaobaoScitemUpdateRequest) SetItemType(itemType int64) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

// ItemType Getter
func (r TaobaoScitemUpdateRequest) GetItemType() int64 {
    return r.itemType
}
// UpdateProperties Setter
// 需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3
func (r *TaobaoScitemUpdateRequest) SetUpdateProperties(updateProperties string) error {
    r.updateProperties = updateProperties
    r.Set("update_properties", updateProperties)
    return nil
}

// UpdateProperties Getter
func (r TaobaoScitemUpdateRequest) GetUpdateProperties() string {
    return r.updateProperties
}
// BarCode Setter
// 条形码
func (r *TaobaoScitemUpdateRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

// BarCode Getter
func (r TaobaoScitemUpdateRequest) GetBarCode() string {
    return r.barCode
}
// WmsCode Setter
// 仓储商编码
func (r *TaobaoScitemUpdateRequest) SetWmsCode(wmsCode string) error {
    r.wmsCode = wmsCode
    r.Set("wms_code", wmsCode)
    return nil
}

// WmsCode Getter
func (r TaobaoScitemUpdateRequest) GetWmsCode() string {
    return r.wmsCode
}
// IsFriable Setter
// 是否易碎 0：不是  1：是
func (r *TaobaoScitemUpdateRequest) SetIsFriable(isFriable int64) error {
    r.isFriable = isFriable
    r.Set("is_friable", isFriable)
    return nil
}

// IsFriable Getter
func (r TaobaoScitemUpdateRequest) GetIsFriable() int64 {
    return r.isFriable
}
// IsDangerous Setter
// 是否危险 0：不是  0：是
func (r *TaobaoScitemUpdateRequest) SetIsDangerous(isDangerous int64) error {
    r.isDangerous = isDangerous
    r.Set("is_dangerous", isDangerous)
    return nil
}

// IsDangerous Getter
func (r TaobaoScitemUpdateRequest) GetIsDangerous() int64 {
    return r.isDangerous
}
// IsCostly Setter
// 是否是贵重品 0:不是 1：是
func (r *TaobaoScitemUpdateRequest) SetIsCostly(isCostly int64) error {
    r.isCostly = isCostly
    r.Set("is_costly", isCostly)
    return nil
}

// IsCostly Getter
func (r TaobaoScitemUpdateRequest) GetIsCostly() int64 {
    return r.isCostly
}
// IsWarranty Setter
// 是否保质期：0:不是 1：是
func (r *TaobaoScitemUpdateRequest) SetIsWarranty(isWarranty int64) error {
    r.isWarranty = isWarranty
    r.Set("is_warranty", isWarranty)
    return nil
}

// IsWarranty Getter
func (r TaobaoScitemUpdateRequest) GetIsWarranty() int64 {
    return r.isWarranty
}
// Weight Setter
// weight
func (r *TaobaoScitemUpdateRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

// Weight Getter
func (r TaobaoScitemUpdateRequest) GetWeight() int64 {
    return r.weight
}
// Length Setter
// 长度 单位：mm
func (r *TaobaoScitemUpdateRequest) SetLength(length int64) error {
    r.length = length
    r.Set("length", length)
    return nil
}

// Length Getter
func (r TaobaoScitemUpdateRequest) GetLength() int64 {
    return r.length
}
// Width Setter
// 宽 单位：mm
func (r *TaobaoScitemUpdateRequest) SetWidth(width int64) error {
    r.width = width
    r.Set("width", width)
    return nil
}

// Width Getter
func (r TaobaoScitemUpdateRequest) GetWidth() int64 {
    return r.width
}
// Height Setter
// 高 单位：mm
func (r *TaobaoScitemUpdateRequest) SetHeight(height int64) error {
    r.height = height
    r.Set("height", height)
    return nil
}

// Height Getter
func (r TaobaoScitemUpdateRequest) GetHeight() int64 {
    return r.height
}
// Volume Setter
// 体积：立方厘米
func (r *TaobaoScitemUpdateRequest) SetVolume(volume int64) error {
    r.volume = volume
    r.Set("volume", volume)
    return nil
}

// Volume Getter
func (r TaobaoScitemUpdateRequest) GetVolume() int64 {
    return r.volume
}
// Price Setter
// price
func (r *TaobaoScitemUpdateRequest) SetPrice(price int64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

// Price Getter
func (r TaobaoScitemUpdateRequest) GetPrice() int64 {
    return r.price
}
// Remark Setter
// remark
func (r *TaobaoScitemUpdateRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r TaobaoScitemUpdateRequest) GetRemark() string {
    return r.remark
}
// MatterStatus Setter
// 0:液体，1：粉体，2：固体
func (r *TaobaoScitemUpdateRequest) SetMatterStatus(matterStatus int64) error {
    r.matterStatus = matterStatus
    r.Set("matter_status", matterStatus)
    return nil
}

// MatterStatus Getter
func (r TaobaoScitemUpdateRequest) GetMatterStatus() int64 {
    return r.matterStatus
}
// BrandId Setter
// 品牌id
func (r *TaobaoScitemUpdateRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

// BrandId Getter
func (r TaobaoScitemUpdateRequest) GetBrandId() int64 {
    return r.brandId
}
// BrandName Setter
// brand_Name
func (r *TaobaoScitemUpdateRequest) SetBrandName(brandName string) error {
    r.brandName = brandName
    r.Set("brand_name", brandName)
    return nil
}

// BrandName Getter
func (r TaobaoScitemUpdateRequest) GetBrandName() string {
    return r.brandName
}
// SpuId Setter
// 淘宝SKU产品级编码CSPU ID
func (r *TaobaoScitemUpdateRequest) SetSpuId(spuId int64) error {
    r.spuId = spuId
    r.Set("spu_id", spuId)
    return nil
}

// SpuId Getter
func (r TaobaoScitemUpdateRequest) GetSpuId() int64 {
    return r.spuId
}
// RemoveProperties Setter
// 移除商品属性P列表,P由系统分配：p1；p2
func (r *TaobaoScitemUpdateRequest) SetRemoveProperties(removeProperties string) error {
    r.removeProperties = removeProperties
    r.Set("remove_properties", removeProperties)
    return nil
}

// RemoveProperties Getter
func (r TaobaoScitemUpdateRequest) GetRemoveProperties() string {
    return r.removeProperties
}
// IsAreaSale Setter
// 1表示区域销售，0或是空是非区域销售
func (r *TaobaoScitemUpdateRequest) SetIsAreaSale(isAreaSale int64) error {
    r.isAreaSale = isAreaSale
    r.Set("is_area_sale", isAreaSale)
    return nil
}

// IsAreaSale Getter
func (r TaobaoScitemUpdateRequest) GetIsAreaSale() int64 {
    return r.isAreaSale
}
