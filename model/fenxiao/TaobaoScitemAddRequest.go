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
    _itemName   string
    // 商家编码
    _outerCode   string
    // 0.普通供应链商品 1.供应链组合主商品
    _itemType   int64
    // 商品属性格式是  p1:v1,p2:v2,p3:v3
    _properties   string
    // 条形码
    _barCode   string
    // 仓储商编码
    _wmsCode   string
    // 是否易碎 0：不是  1：是
    _isFriable   int64
    // 是否危险 0：不是  1：是
    _isDangerous   int64
    // 是否是贵重品 0:不是 1：是
    _isCostly   int64
    // 是否保质期：0:不是 1：是
    _isWarranty   int64
    // 重量 单位：g
    _weight   int64
    // 长度 单位：mm
    _length   int64
    // 宽 单位：mm
    _width   int64
    // 高 单位：mm
    _height   int64
    // 体积：立方厘米
    _volume   int64
    // 价格 单位：分
    _price   int64
    // remark
    _remark   string
    // 0:液体，1：粉体，2：固体
    _matterStatus   int64
    // 品牌id
    _brandId   int64
    // brand_Name
    _brandName   string
    // spuId或是cspuid
    _spuId   int64
    // 1表示区域销售，0或是空是非区域销售
    _isAreaSale   int64
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
func (r *TaobaoScitemAddRequest) SetItemName(_itemName string) error {
    r._itemName = _itemName
    r.Set("item_name", _itemName)
    return nil
}

// ItemName Getter
func (r TaobaoScitemAddRequest) GetItemName() string {
    return r._itemName
}
// OuterCode Setter
// 商家编码
func (r *TaobaoScitemAddRequest) SetOuterCode(_outerCode string) error {
    r._outerCode = _outerCode
    r.Set("outer_code", _outerCode)
    return nil
}

// OuterCode Getter
func (r TaobaoScitemAddRequest) GetOuterCode() string {
    return r._outerCode
}
// ItemType Setter
// 0.普通供应链商品 1.供应链组合主商品
func (r *TaobaoScitemAddRequest) SetItemType(_itemType int64) error {
    r._itemType = _itemType
    r.Set("item_type", _itemType)
    return nil
}

// ItemType Getter
func (r TaobaoScitemAddRequest) GetItemType() int64 {
    return r._itemType
}
// Properties Setter
// 商品属性格式是  p1:v1,p2:v2,p3:v3
func (r *TaobaoScitemAddRequest) SetProperties(_properties string) error {
    r._properties = _properties
    r.Set("properties", _properties)
    return nil
}

// Properties Getter
func (r TaobaoScitemAddRequest) GetProperties() string {
    return r._properties
}
// BarCode Setter
// 条形码
func (r *TaobaoScitemAddRequest) SetBarCode(_barCode string) error {
    r._barCode = _barCode
    r.Set("bar_code", _barCode)
    return nil
}

// BarCode Getter
func (r TaobaoScitemAddRequest) GetBarCode() string {
    return r._barCode
}
// WmsCode Setter
// 仓储商编码
func (r *TaobaoScitemAddRequest) SetWmsCode(_wmsCode string) error {
    r._wmsCode = _wmsCode
    r.Set("wms_code", _wmsCode)
    return nil
}

// WmsCode Getter
func (r TaobaoScitemAddRequest) GetWmsCode() string {
    return r._wmsCode
}
// IsFriable Setter
// 是否易碎 0：不是  1：是
func (r *TaobaoScitemAddRequest) SetIsFriable(_isFriable int64) error {
    r._isFriable = _isFriable
    r.Set("is_friable", _isFriable)
    return nil
}

// IsFriable Getter
func (r TaobaoScitemAddRequest) GetIsFriable() int64 {
    return r._isFriable
}
// IsDangerous Setter
// 是否危险 0：不是  1：是
func (r *TaobaoScitemAddRequest) SetIsDangerous(_isDangerous int64) error {
    r._isDangerous = _isDangerous
    r.Set("is_dangerous", _isDangerous)
    return nil
}

// IsDangerous Getter
func (r TaobaoScitemAddRequest) GetIsDangerous() int64 {
    return r._isDangerous
}
// IsCostly Setter
// 是否是贵重品 0:不是 1：是
func (r *TaobaoScitemAddRequest) SetIsCostly(_isCostly int64) error {
    r._isCostly = _isCostly
    r.Set("is_costly", _isCostly)
    return nil
}

// IsCostly Getter
func (r TaobaoScitemAddRequest) GetIsCostly() int64 {
    return r._isCostly
}
// IsWarranty Setter
// 是否保质期：0:不是 1：是
func (r *TaobaoScitemAddRequest) SetIsWarranty(_isWarranty int64) error {
    r._isWarranty = _isWarranty
    r.Set("is_warranty", _isWarranty)
    return nil
}

// IsWarranty Getter
func (r TaobaoScitemAddRequest) GetIsWarranty() int64 {
    return r._isWarranty
}
// Weight Setter
// 重量 单位：g
func (r *TaobaoScitemAddRequest) SetWeight(_weight int64) error {
    r._weight = _weight
    r.Set("weight", _weight)
    return nil
}

// Weight Getter
func (r TaobaoScitemAddRequest) GetWeight() int64 {
    return r._weight
}
// Length Setter
// 长度 单位：mm
func (r *TaobaoScitemAddRequest) SetLength(_length int64) error {
    r._length = _length
    r.Set("length", _length)
    return nil
}

// Length Getter
func (r TaobaoScitemAddRequest) GetLength() int64 {
    return r._length
}
// Width Setter
// 宽 单位：mm
func (r *TaobaoScitemAddRequest) SetWidth(_width int64) error {
    r._width = _width
    r.Set("width", _width)
    return nil
}

// Width Getter
func (r TaobaoScitemAddRequest) GetWidth() int64 {
    return r._width
}
// Height Setter
// 高 单位：mm
func (r *TaobaoScitemAddRequest) SetHeight(_height int64) error {
    r._height = _height
    r.Set("height", _height)
    return nil
}

// Height Getter
func (r TaobaoScitemAddRequest) GetHeight() int64 {
    return r._height
}
// Volume Setter
// 体积：立方厘米
func (r *TaobaoScitemAddRequest) SetVolume(_volume int64) error {
    r._volume = _volume
    r.Set("volume", _volume)
    return nil
}

// Volume Getter
func (r TaobaoScitemAddRequest) GetVolume() int64 {
    return r._volume
}
// Price Setter
// 价格 单位：分
func (r *TaobaoScitemAddRequest) SetPrice(_price int64) error {
    r._price = _price
    r.Set("price", _price)
    return nil
}

// Price Getter
func (r TaobaoScitemAddRequest) GetPrice() int64 {
    return r._price
}
// Remark Setter
// remark
func (r *TaobaoScitemAddRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoScitemAddRequest) GetRemark() string {
    return r._remark
}
// MatterStatus Setter
// 0:液体，1：粉体，2：固体
func (r *TaobaoScitemAddRequest) SetMatterStatus(_matterStatus int64) error {
    r._matterStatus = _matterStatus
    r.Set("matter_status", _matterStatus)
    return nil
}

// MatterStatus Getter
func (r TaobaoScitemAddRequest) GetMatterStatus() int64 {
    return r._matterStatus
}
// BrandId Setter
// 品牌id
func (r *TaobaoScitemAddRequest) SetBrandId(_brandId int64) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r TaobaoScitemAddRequest) GetBrandId() int64 {
    return r._brandId
}
// BrandName Setter
// brand_Name
func (r *TaobaoScitemAddRequest) SetBrandName(_brandName string) error {
    r._brandName = _brandName
    r.Set("brand_name", _brandName)
    return nil
}

// BrandName Getter
func (r TaobaoScitemAddRequest) GetBrandName() string {
    return r._brandName
}
// SpuId Setter
// spuId或是cspuid
func (r *TaobaoScitemAddRequest) SetSpuId(_spuId int64) error {
    r._spuId = _spuId
    r.Set("spu_id", _spuId)
    return nil
}

// SpuId Getter
func (r TaobaoScitemAddRequest) GetSpuId() int64 {
    return r._spuId
}
// IsAreaSale Setter
// 1表示区域销售，0或是空是非区域销售
func (r *TaobaoScitemAddRequest) SetIsAreaSale(_isAreaSale int64) error {
    r._isAreaSale = _isAreaSale
    r.Set("is_area_sale", _isAreaSale)
    return nil
}

// IsAreaSale Getter
func (r TaobaoScitemAddRequest) GetIsAreaSale() int64 {
    return r._isAreaSale
}
