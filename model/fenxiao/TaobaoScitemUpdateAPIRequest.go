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
type TaobaoScitemUpdateAPIRequest struct {
    model.Params
    // 后端商品ID，跟outer_code必须指定一个
    _itemId   int64
    // 商家编码，跟item_id必须指定一个
    _outerCode   string
    // 商品名称
    _itemName   string
    // 0.普通供应链商品 1.供应链组合主商品
    _itemType   int64
    // 需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3
    _updateProperties   string
    // 条形码
    _barCode   string
    // 仓储商编码
    _wmsCode   string
    // 是否易碎 0：不是  1：是
    _isFriable   int64
    // 是否危险 0：不是  0：是
    _isDangerous   int64
    // 是否是贵重品 0:不是 1：是
    _isCostly   int64
    // 是否保质期：0:不是 1：是
    _isWarranty   int64
    // weight
    _weight   int64
    // 长度 单位：mm
    _length   int64
    // 宽 单位：mm
    _width   int64
    // 高 单位：mm
    _height   int64
    // 体积：立方厘米
    _volume   int64
    // price
    _price   int64
    // remark
    _remark   string
    // 0:液体，1：粉体，2：固体
    _matterStatus   int64
    // 品牌id
    _brandId   int64
    // brand_Name
    _brandName   string
    // 淘宝SKU产品级编码CSPU ID
    _spuId   int64
    // 移除商品属性P列表,P由系统分配：p1；p2
    _removeProperties   string
    // 1表示区域销售，0或是空是非区域销售
    _isAreaSale   int64
}

// 初始化TaobaoScitemUpdateAPIRequest对象
func NewTaobaoScitemUpdateRequest() *TaobaoScitemUpdateAPIRequest{
    return &TaobaoScitemUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoScitemUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.scitem.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoScitemUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 后端商品ID，跟outer_code必须指定一个
func (r *TaobaoScitemUpdateAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoScitemUpdateAPIRequest) GetItemId() int64 {
    return r._itemId
}
// OuterCode Setter
// 商家编码，跟item_id必须指定一个
func (r *TaobaoScitemUpdateAPIRequest) SetOuterCode(_outerCode string) error {
    r._outerCode = _outerCode
    r.Set("outer_code", _outerCode)
    return nil
}

// OuterCode Getter
func (r TaobaoScitemUpdateAPIRequest) GetOuterCode() string {
    return r._outerCode
}
// ItemName Setter
// 商品名称
func (r *TaobaoScitemUpdateAPIRequest) SetItemName(_itemName string) error {
    r._itemName = _itemName
    r.Set("item_name", _itemName)
    return nil
}

// ItemName Getter
func (r TaobaoScitemUpdateAPIRequest) GetItemName() string {
    return r._itemName
}
// ItemType Setter
// 0.普通供应链商品 1.供应链组合主商品
func (r *TaobaoScitemUpdateAPIRequest) SetItemType(_itemType int64) error {
    r._itemType = _itemType
    r.Set("item_type", _itemType)
    return nil
}

// ItemType Getter
func (r TaobaoScitemUpdateAPIRequest) GetItemType() int64 {
    return r._itemType
}
// UpdateProperties Setter
// 需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3
func (r *TaobaoScitemUpdateAPIRequest) SetUpdateProperties(_updateProperties string) error {
    r._updateProperties = _updateProperties
    r.Set("update_properties", _updateProperties)
    return nil
}

// UpdateProperties Getter
func (r TaobaoScitemUpdateAPIRequest) GetUpdateProperties() string {
    return r._updateProperties
}
// BarCode Setter
// 条形码
func (r *TaobaoScitemUpdateAPIRequest) SetBarCode(_barCode string) error {
    r._barCode = _barCode
    r.Set("bar_code", _barCode)
    return nil
}

// BarCode Getter
func (r TaobaoScitemUpdateAPIRequest) GetBarCode() string {
    return r._barCode
}
// WmsCode Setter
// 仓储商编码
func (r *TaobaoScitemUpdateAPIRequest) SetWmsCode(_wmsCode string) error {
    r._wmsCode = _wmsCode
    r.Set("wms_code", _wmsCode)
    return nil
}

// WmsCode Getter
func (r TaobaoScitemUpdateAPIRequest) GetWmsCode() string {
    return r._wmsCode
}
// IsFriable Setter
// 是否易碎 0：不是  1：是
func (r *TaobaoScitemUpdateAPIRequest) SetIsFriable(_isFriable int64) error {
    r._isFriable = _isFriable
    r.Set("is_friable", _isFriable)
    return nil
}

// IsFriable Getter
func (r TaobaoScitemUpdateAPIRequest) GetIsFriable() int64 {
    return r._isFriable
}
// IsDangerous Setter
// 是否危险 0：不是  0：是
func (r *TaobaoScitemUpdateAPIRequest) SetIsDangerous(_isDangerous int64) error {
    r._isDangerous = _isDangerous
    r.Set("is_dangerous", _isDangerous)
    return nil
}

// IsDangerous Getter
func (r TaobaoScitemUpdateAPIRequest) GetIsDangerous() int64 {
    return r._isDangerous
}
// IsCostly Setter
// 是否是贵重品 0:不是 1：是
func (r *TaobaoScitemUpdateAPIRequest) SetIsCostly(_isCostly int64) error {
    r._isCostly = _isCostly
    r.Set("is_costly", _isCostly)
    return nil
}

// IsCostly Getter
func (r TaobaoScitemUpdateAPIRequest) GetIsCostly() int64 {
    return r._isCostly
}
// IsWarranty Setter
// 是否保质期：0:不是 1：是
func (r *TaobaoScitemUpdateAPIRequest) SetIsWarranty(_isWarranty int64) error {
    r._isWarranty = _isWarranty
    r.Set("is_warranty", _isWarranty)
    return nil
}

// IsWarranty Getter
func (r TaobaoScitemUpdateAPIRequest) GetIsWarranty() int64 {
    return r._isWarranty
}
// Weight Setter
// weight
func (r *TaobaoScitemUpdateAPIRequest) SetWeight(_weight int64) error {
    r._weight = _weight
    r.Set("weight", _weight)
    return nil
}

// Weight Getter
func (r TaobaoScitemUpdateAPIRequest) GetWeight() int64 {
    return r._weight
}
// Length Setter
// 长度 单位：mm
func (r *TaobaoScitemUpdateAPIRequest) SetLength(_length int64) error {
    r._length = _length
    r.Set("length", _length)
    return nil
}

// Length Getter
func (r TaobaoScitemUpdateAPIRequest) GetLength() int64 {
    return r._length
}
// Width Setter
// 宽 单位：mm
func (r *TaobaoScitemUpdateAPIRequest) SetWidth(_width int64) error {
    r._width = _width
    r.Set("width", _width)
    return nil
}

// Width Getter
func (r TaobaoScitemUpdateAPIRequest) GetWidth() int64 {
    return r._width
}
// Height Setter
// 高 单位：mm
func (r *TaobaoScitemUpdateAPIRequest) SetHeight(_height int64) error {
    r._height = _height
    r.Set("height", _height)
    return nil
}

// Height Getter
func (r TaobaoScitemUpdateAPIRequest) GetHeight() int64 {
    return r._height
}
// Volume Setter
// 体积：立方厘米
func (r *TaobaoScitemUpdateAPIRequest) SetVolume(_volume int64) error {
    r._volume = _volume
    r.Set("volume", _volume)
    return nil
}

// Volume Getter
func (r TaobaoScitemUpdateAPIRequest) GetVolume() int64 {
    return r._volume
}
// Price Setter
// price
func (r *TaobaoScitemUpdateAPIRequest) SetPrice(_price int64) error {
    r._price = _price
    r.Set("price", _price)
    return nil
}

// Price Getter
func (r TaobaoScitemUpdateAPIRequest) GetPrice() int64 {
    return r._price
}
// Remark Setter
// remark
func (r *TaobaoScitemUpdateAPIRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoScitemUpdateAPIRequest) GetRemark() string {
    return r._remark
}
// MatterStatus Setter
// 0:液体，1：粉体，2：固体
func (r *TaobaoScitemUpdateAPIRequest) SetMatterStatus(_matterStatus int64) error {
    r._matterStatus = _matterStatus
    r.Set("matter_status", _matterStatus)
    return nil
}

// MatterStatus Getter
func (r TaobaoScitemUpdateAPIRequest) GetMatterStatus() int64 {
    return r._matterStatus
}
// BrandId Setter
// 品牌id
func (r *TaobaoScitemUpdateAPIRequest) SetBrandId(_brandId int64) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r TaobaoScitemUpdateAPIRequest) GetBrandId() int64 {
    return r._brandId
}
// BrandName Setter
// brand_Name
func (r *TaobaoScitemUpdateAPIRequest) SetBrandName(_brandName string) error {
    r._brandName = _brandName
    r.Set("brand_name", _brandName)
    return nil
}

// BrandName Getter
func (r TaobaoScitemUpdateAPIRequest) GetBrandName() string {
    return r._brandName
}
// SpuId Setter
// 淘宝SKU产品级编码CSPU ID
func (r *TaobaoScitemUpdateAPIRequest) SetSpuId(_spuId int64) error {
    r._spuId = _spuId
    r.Set("spu_id", _spuId)
    return nil
}

// SpuId Getter
func (r TaobaoScitemUpdateAPIRequest) GetSpuId() int64 {
    return r._spuId
}
// RemoveProperties Setter
// 移除商品属性P列表,P由系统分配：p1；p2
func (r *TaobaoScitemUpdateAPIRequest) SetRemoveProperties(_removeProperties string) error {
    r._removeProperties = _removeProperties
    r.Set("remove_properties", _removeProperties)
    return nil
}

// RemoveProperties Getter
func (r TaobaoScitemUpdateAPIRequest) GetRemoveProperties() string {
    return r._removeProperties
}
// IsAreaSale Setter
// 1表示区域销售，0或是空是非区域销售
func (r *TaobaoScitemUpdateAPIRequest) SetIsAreaSale(_isAreaSale int64) error {
    r._isAreaSale = _isAreaSale
    r.Set("is_area_sale", _isAreaSale)
    return nil
}

// IsAreaSale Getter
func (r TaobaoScitemUpdateAPIRequest) GetIsAreaSale() int64 {
    return r._isAreaSale
}
