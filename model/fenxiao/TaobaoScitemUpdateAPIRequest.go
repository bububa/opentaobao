package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemUpdateAPIRequest 根据商品ID或商家编码修改后端商品 API请求
// taobao.scitem.update
//
// 根据商品ID或商家编码修改后端商品
type TaobaoScitemUpdateAPIRequest struct {
	model.Params
	// 商家编码，跟item_id必须指定一个
	_outerCode string
	// 商品名称
	_itemName string
	// 需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3
	_updateProperties string
	// 条形码
	_barCode string
	// 仓储商编码
	_wmsCode string
	// remark
	_remark string
	// brand_Name
	_brandName string
	// 移除商品属性P列表,P由系统分配：p1；p2
	_removeProperties string
	// 后端商品ID，跟outer_code必须指定一个
	_itemId int64
	// 0.普通供应链商品 1.供应链组合主商品
	_itemType int64
	// 是否易碎 0：不是  1：是
	_isFriable int64
	// 是否危险 0：不是  0：是
	_isDangerous int64
	// 是否是贵重品 0:不是 1：是
	_isCostly int64
	// 是否保质期：0:不是 1：是
	_isWarranty int64
	// weight
	_weight int64
	// 长度 单位：mm
	_length int64
	// 宽 单位：mm
	_width int64
	// 高 单位：mm
	_height int64
	// 体积：立方厘米
	_volume int64
	// price
	_price int64
	// 0:液体，1：粉体，2：固体
	_matterStatus int64
	// 品牌id
	_brandId int64
	// 淘宝SKU产品级编码CSPU ID
	_spuId int64
	// 1表示区域销售，0或是空是非区域销售
	_isAreaSale int64
}

// NewTaobaoScitemUpdateRequest 初始化TaobaoScitemUpdateAPIRequest对象
func NewTaobaoScitemUpdateRequest() *TaobaoScitemUpdateAPIRequest {
	return &TaobaoScitemUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoScitemUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoScitemUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoScitemUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterCode is OuterCode Setter
// 商家编码，跟item_id必须指定一个
func (r *TaobaoScitemUpdateAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// GetOuterCode OuterCode Getter
func (r TaobaoScitemUpdateAPIRequest) GetOuterCode() string {
	return r._outerCode
}

// SetItemName is ItemName Setter
// 商品名称
func (r *TaobaoScitemUpdateAPIRequest) SetItemName(_itemName string) error {
	r._itemName = _itemName
	r.Set("item_name", _itemName)
	return nil
}

// GetItemName ItemName Getter
func (r TaobaoScitemUpdateAPIRequest) GetItemName() string {
	return r._itemName
}

// SetUpdateProperties is UpdateProperties Setter
// 需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3
func (r *TaobaoScitemUpdateAPIRequest) SetUpdateProperties(_updateProperties string) error {
	r._updateProperties = _updateProperties
	r.Set("update_properties", _updateProperties)
	return nil
}

// GetUpdateProperties UpdateProperties Getter
func (r TaobaoScitemUpdateAPIRequest) GetUpdateProperties() string {
	return r._updateProperties
}

// SetBarCode is BarCode Setter
// 条形码
func (r *TaobaoScitemUpdateAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// GetBarCode BarCode Getter
func (r TaobaoScitemUpdateAPIRequest) GetBarCode() string {
	return r._barCode
}

// SetWmsCode is WmsCode Setter
// 仓储商编码
func (r *TaobaoScitemUpdateAPIRequest) SetWmsCode(_wmsCode string) error {
	r._wmsCode = _wmsCode
	r.Set("wms_code", _wmsCode)
	return nil
}

// GetWmsCode WmsCode Getter
func (r TaobaoScitemUpdateAPIRequest) GetWmsCode() string {
	return r._wmsCode
}

// SetRemark is Remark Setter
// remark
func (r *TaobaoScitemUpdateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoScitemUpdateAPIRequest) GetRemark() string {
	return r._remark
}

// SetBrandName is BrandName Setter
// brand_Name
func (r *TaobaoScitemUpdateAPIRequest) SetBrandName(_brandName string) error {
	r._brandName = _brandName
	r.Set("brand_name", _brandName)
	return nil
}

// GetBrandName BrandName Getter
func (r TaobaoScitemUpdateAPIRequest) GetBrandName() string {
	return r._brandName
}

// SetRemoveProperties is RemoveProperties Setter
// 移除商品属性P列表,P由系统分配：p1；p2
func (r *TaobaoScitemUpdateAPIRequest) SetRemoveProperties(_removeProperties string) error {
	r._removeProperties = _removeProperties
	r.Set("remove_properties", _removeProperties)
	return nil
}

// GetRemoveProperties RemoveProperties Getter
func (r TaobaoScitemUpdateAPIRequest) GetRemoveProperties() string {
	return r._removeProperties
}

// SetItemId is ItemId Setter
// 后端商品ID，跟outer_code必须指定一个
func (r *TaobaoScitemUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoScitemUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemType is ItemType Setter
// 0.普通供应链商品 1.供应链组合主商品
func (r *TaobaoScitemUpdateAPIRequest) SetItemType(_itemType int64) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// GetItemType ItemType Getter
func (r TaobaoScitemUpdateAPIRequest) GetItemType() int64 {
	return r._itemType
}

// SetIsFriable is IsFriable Setter
// 是否易碎 0：不是  1：是
func (r *TaobaoScitemUpdateAPIRequest) SetIsFriable(_isFriable int64) error {
	r._isFriable = _isFriable
	r.Set("is_friable", _isFriable)
	return nil
}

// GetIsFriable IsFriable Getter
func (r TaobaoScitemUpdateAPIRequest) GetIsFriable() int64 {
	return r._isFriable
}

// SetIsDangerous is IsDangerous Setter
// 是否危险 0：不是  0：是
func (r *TaobaoScitemUpdateAPIRequest) SetIsDangerous(_isDangerous int64) error {
	r._isDangerous = _isDangerous
	r.Set("is_dangerous", _isDangerous)
	return nil
}

// GetIsDangerous IsDangerous Getter
func (r TaobaoScitemUpdateAPIRequest) GetIsDangerous() int64 {
	return r._isDangerous
}

// SetIsCostly is IsCostly Setter
// 是否是贵重品 0:不是 1：是
func (r *TaobaoScitemUpdateAPIRequest) SetIsCostly(_isCostly int64) error {
	r._isCostly = _isCostly
	r.Set("is_costly", _isCostly)
	return nil
}

// GetIsCostly IsCostly Getter
func (r TaobaoScitemUpdateAPIRequest) GetIsCostly() int64 {
	return r._isCostly
}

// SetIsWarranty is IsWarranty Setter
// 是否保质期：0:不是 1：是
func (r *TaobaoScitemUpdateAPIRequest) SetIsWarranty(_isWarranty int64) error {
	r._isWarranty = _isWarranty
	r.Set("is_warranty", _isWarranty)
	return nil
}

// GetIsWarranty IsWarranty Getter
func (r TaobaoScitemUpdateAPIRequest) GetIsWarranty() int64 {
	return r._isWarranty
}

// SetWeight is Weight Setter
// weight
func (r *TaobaoScitemUpdateAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r TaobaoScitemUpdateAPIRequest) GetWeight() int64 {
	return r._weight
}

// SetLength is Length Setter
// 长度 单位：mm
func (r *TaobaoScitemUpdateAPIRequest) SetLength(_length int64) error {
	r._length = _length
	r.Set("length", _length)
	return nil
}

// GetLength Length Getter
func (r TaobaoScitemUpdateAPIRequest) GetLength() int64 {
	return r._length
}

// SetWidth is Width Setter
// 宽 单位：mm
func (r *TaobaoScitemUpdateAPIRequest) SetWidth(_width int64) error {
	r._width = _width
	r.Set("width", _width)
	return nil
}

// GetWidth Width Getter
func (r TaobaoScitemUpdateAPIRequest) GetWidth() int64 {
	return r._width
}

// SetHeight is Height Setter
// 高 单位：mm
func (r *TaobaoScitemUpdateAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// GetHeight Height Getter
func (r TaobaoScitemUpdateAPIRequest) GetHeight() int64 {
	return r._height
}

// SetVolume is Volume Setter
// 体积：立方厘米
func (r *TaobaoScitemUpdateAPIRequest) SetVolume(_volume int64) error {
	r._volume = _volume
	r.Set("volume", _volume)
	return nil
}

// GetVolume Volume Getter
func (r TaobaoScitemUpdateAPIRequest) GetVolume() int64 {
	return r._volume
}

// SetPrice is Price Setter
// price
func (r *TaobaoScitemUpdateAPIRequest) SetPrice(_price int64) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TaobaoScitemUpdateAPIRequest) GetPrice() int64 {
	return r._price
}

// SetMatterStatus is MatterStatus Setter
// 0:液体，1：粉体，2：固体
func (r *TaobaoScitemUpdateAPIRequest) SetMatterStatus(_matterStatus int64) error {
	r._matterStatus = _matterStatus
	r.Set("matter_status", _matterStatus)
	return nil
}

// GetMatterStatus MatterStatus Getter
func (r TaobaoScitemUpdateAPIRequest) GetMatterStatus() int64 {
	return r._matterStatus
}

// SetBrandId is BrandId Setter
// 品牌id
func (r *TaobaoScitemUpdateAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r TaobaoScitemUpdateAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// SetSpuId is SpuId Setter
// 淘宝SKU产品级编码CSPU ID
func (r *TaobaoScitemUpdateAPIRequest) SetSpuId(_spuId int64) error {
	r._spuId = _spuId
	r.Set("spu_id", _spuId)
	return nil
}

// GetSpuId SpuId Getter
func (r TaobaoScitemUpdateAPIRequest) GetSpuId() int64 {
	return r._spuId
}

// SetIsAreaSale is IsAreaSale Setter
// 1表示区域销售，0或是空是非区域销售
func (r *TaobaoScitemUpdateAPIRequest) SetIsAreaSale(_isAreaSale int64) error {
	r._isAreaSale = _isAreaSale
	r.Set("is_area_sale", _isAreaSale)
	return nil
}

// GetIsAreaSale IsAreaSale Getter
func (r TaobaoScitemUpdateAPIRequest) GetIsAreaSale() int64 {
	return r._isAreaSale
}
