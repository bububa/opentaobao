package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流宝商品修改 API请求
taobao.wlb.item.update

修改物流宝商品信息
*/
type TaobaoWlbItemUpdateRequest struct {
    model.Params
    // 需要修改的商品属性值的列表，如果属性不存在，则新增属性
    updatePropertyKeyList   string
    // 需要删除的商品属性key列表
    deletePropertyKeyList   string
    // 需要修改的属性值的列表
    updatePropertyValueList   string
    // 要修改的商品id
    id   int64
    // 要修改的商品名称
    name   string
    // 要修改的商品标题
    title   string
    // 要修改的商品备注
    remark   string
    // 是否易碎品
    isFriable   bool
    // 是否危险品
    isDangerous   bool
    // 商品颜色
    color   string
    // 商品重量，单位G
    weight   int64
    // 商品长度，单位厘米
    length   int64
    // 商品宽度，单位厘米
    width   int64
    // 商品高度，单位厘米
    height   int64
    // 商品体积，单位立方厘米
    volume   int64
    // 商品货类
    goodsCat   string
    // 商品计价货类
    pricingCat   string
    // 商品包装材料类型
    packageMaterial   string
}

// 初始化TaobaoWlbItemUpdateRequest对象
func NewTaobaoWlbItemUpdateRequest() *TaobaoWlbItemUpdateRequest{
    return &TaobaoWlbItemUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemUpdateRequest) GetApiMethodName() string {
    return "taobao.wlb.item.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UpdatePropertyKeyList Setter
// 需要修改的商品属性值的列表，如果属性不存在，则新增属性
func (r *TaobaoWlbItemUpdateRequest) SetUpdatePropertyKeyList(updatePropertyKeyList string) error {
    r.updatePropertyKeyList = updatePropertyKeyList
    r.Set("update_property_key_list", updatePropertyKeyList)
    return nil
}

// UpdatePropertyKeyList Getter
func (r TaobaoWlbItemUpdateRequest) GetUpdatePropertyKeyList() string {
    return r.updatePropertyKeyList
}
// DeletePropertyKeyList Setter
// 需要删除的商品属性key列表
func (r *TaobaoWlbItemUpdateRequest) SetDeletePropertyKeyList(deletePropertyKeyList string) error {
    r.deletePropertyKeyList = deletePropertyKeyList
    r.Set("delete_property_key_list", deletePropertyKeyList)
    return nil
}

// DeletePropertyKeyList Getter
func (r TaobaoWlbItemUpdateRequest) GetDeletePropertyKeyList() string {
    return r.deletePropertyKeyList
}
// UpdatePropertyValueList Setter
// 需要修改的属性值的列表
func (r *TaobaoWlbItemUpdateRequest) SetUpdatePropertyValueList(updatePropertyValueList string) error {
    r.updatePropertyValueList = updatePropertyValueList
    r.Set("update_property_value_list", updatePropertyValueList)
    return nil
}

// UpdatePropertyValueList Getter
func (r TaobaoWlbItemUpdateRequest) GetUpdatePropertyValueList() string {
    return r.updatePropertyValueList
}
// Id Setter
// 要修改的商品id
func (r *TaobaoWlbItemUpdateRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoWlbItemUpdateRequest) GetId() int64 {
    return r.id
}
// Name Setter
// 要修改的商品名称
func (r *TaobaoWlbItemUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoWlbItemUpdateRequest) GetName() string {
    return r.name
}
// Title Setter
// 要修改的商品标题
func (r *TaobaoWlbItemUpdateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoWlbItemUpdateRequest) GetTitle() string {
    return r.title
}
// Remark Setter
// 要修改的商品备注
func (r *TaobaoWlbItemUpdateRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r TaobaoWlbItemUpdateRequest) GetRemark() string {
    return r.remark
}
// IsFriable Setter
// 是否易碎品
func (r *TaobaoWlbItemUpdateRequest) SetIsFriable(isFriable bool) error {
    r.isFriable = isFriable
    r.Set("is_friable", isFriable)
    return nil
}

// IsFriable Getter
func (r TaobaoWlbItemUpdateRequest) GetIsFriable() bool {
    return r.isFriable
}
// IsDangerous Setter
// 是否危险品
func (r *TaobaoWlbItemUpdateRequest) SetIsDangerous(isDangerous bool) error {
    r.isDangerous = isDangerous
    r.Set("is_dangerous", isDangerous)
    return nil
}

// IsDangerous Getter
func (r TaobaoWlbItemUpdateRequest) GetIsDangerous() bool {
    return r.isDangerous
}
// Color Setter
// 商品颜色
func (r *TaobaoWlbItemUpdateRequest) SetColor(color string) error {
    r.color = color
    r.Set("color", color)
    return nil
}

// Color Getter
func (r TaobaoWlbItemUpdateRequest) GetColor() string {
    return r.color
}
// Weight Setter
// 商品重量，单位G
func (r *TaobaoWlbItemUpdateRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

// Weight Getter
func (r TaobaoWlbItemUpdateRequest) GetWeight() int64 {
    return r.weight
}
// Length Setter
// 商品长度，单位厘米
func (r *TaobaoWlbItemUpdateRequest) SetLength(length int64) error {
    r.length = length
    r.Set("length", length)
    return nil
}

// Length Getter
func (r TaobaoWlbItemUpdateRequest) GetLength() int64 {
    return r.length
}
// Width Setter
// 商品宽度，单位厘米
func (r *TaobaoWlbItemUpdateRequest) SetWidth(width int64) error {
    r.width = width
    r.Set("width", width)
    return nil
}

// Width Getter
func (r TaobaoWlbItemUpdateRequest) GetWidth() int64 {
    return r.width
}
// Height Setter
// 商品高度，单位厘米
func (r *TaobaoWlbItemUpdateRequest) SetHeight(height int64) error {
    r.height = height
    r.Set("height", height)
    return nil
}

// Height Getter
func (r TaobaoWlbItemUpdateRequest) GetHeight() int64 {
    return r.height
}
// Volume Setter
// 商品体积，单位立方厘米
func (r *TaobaoWlbItemUpdateRequest) SetVolume(volume int64) error {
    r.volume = volume
    r.Set("volume", volume)
    return nil
}

// Volume Getter
func (r TaobaoWlbItemUpdateRequest) GetVolume() int64 {
    return r.volume
}
// GoodsCat Setter
// 商品货类
func (r *TaobaoWlbItemUpdateRequest) SetGoodsCat(goodsCat string) error {
    r.goodsCat = goodsCat
    r.Set("goods_cat", goodsCat)
    return nil
}

// GoodsCat Getter
func (r TaobaoWlbItemUpdateRequest) GetGoodsCat() string {
    return r.goodsCat
}
// PricingCat Setter
// 商品计价货类
func (r *TaobaoWlbItemUpdateRequest) SetPricingCat(pricingCat string) error {
    r.pricingCat = pricingCat
    r.Set("pricing_cat", pricingCat)
    return nil
}

// PricingCat Getter
func (r TaobaoWlbItemUpdateRequest) GetPricingCat() string {
    return r.pricingCat
}
// PackageMaterial Setter
// 商品包装材料类型
func (r *TaobaoWlbItemUpdateRequest) SetPackageMaterial(packageMaterial string) error {
    r.packageMaterial = packageMaterial
    r.Set("package_material", packageMaterial)
    return nil
}

// PackageMaterial Getter
func (r TaobaoWlbItemUpdateRequest) GetPackageMaterial() string {
    return r.packageMaterial
}
