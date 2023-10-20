package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoscitemaddAPIRequest 发布后端商品 API请求
// taobao.scitem.add
//
// 发布后端商品
type TaobaoscitemaddAPIRequest struct {
	model.Params
	// 商品名称
	_itemName string
	// 商家编码
	_outerCode string
	// 商品属性格式是  p1:v1,p2:v2,p3:v3
	_properties string
	// 条形码
	_barCode string
	// 仓储商编码
	_wmsCode string
	// remark
	_remark string
	// brand_Name
	_brandName string
	// 0.普通供应链商品 1.供应链组合主商品
	_itemType int64
	// 是否易碎 0：不是  1：是
	_isFriable int64
	// 是否危险 0：不是  1：是
	_isDangerous int64
	// 是否是贵重品 0:不是 1：是
	_isCostly int64
	// 是否保质期：0:不是 1：是
	_isWarranty int64
	// 重量 单位：g
	_weight int64
	// 长度 单位：mm
	_length int64
	// 宽 单位：mm
	_width int64
	// 高 单位：mm
	_height int64
	// 体积：立方厘米
	_volume int64
	// 价格 单位：分
	_price int64
	// 0:液体，1：粉体，2：固体
	_matterStatus int64
	// 品牌id
	_brandId int64
	// spuId或是cspuid
	_spuId int64
	// 1表示区域销售，0或是空是非区域销售
	_isAreaSale int64
}

// NewTaobaoscitemaddRequest 初始化TaobaoscitemaddAPIRequest对象
func NewTaobaoscitemaddRequest() *TaobaoscitemaddAPIRequest {
	return &TaobaoscitemaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoscitemaddAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoscitemaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoscitemaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemName is ItemName Setter
// 商品名称
func (r *TaobaoscitemaddAPIRequest) SetItemName(_itemName string) error {
	r._itemName = _itemName
	r.Set("item_name", _itemName)
	return nil
}

// GetItemName ItemName Getter
func (r TaobaoscitemaddAPIRequest) GetItemName() string {
	return r._itemName
}

// SetOuterCode is OuterCode Setter
// 商家编码
func (r *TaobaoscitemaddAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// GetOuterCode OuterCode Getter
func (r TaobaoscitemaddAPIRequest) GetOuterCode() string {
	return r._outerCode
}

// SetProperties is Properties Setter
// 商品属性格式是  p1:v1,p2:v2,p3:v3
func (r *TaobaoscitemaddAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoscitemaddAPIRequest) GetProperties() string {
	return r._properties
}

// SetBarCode is BarCode Setter
// 条形码
func (r *TaobaoscitemaddAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// GetBarCode BarCode Getter
func (r TaobaoscitemaddAPIRequest) GetBarCode() string {
	return r._barCode
}

// SetWmsCode is WmsCode Setter
// 仓储商编码
func (r *TaobaoscitemaddAPIRequest) SetWmsCode(_wmsCode string) error {
	r._wmsCode = _wmsCode
	r.Set("wms_code", _wmsCode)
	return nil
}

// GetWmsCode WmsCode Getter
func (r TaobaoscitemaddAPIRequest) GetWmsCode() string {
	return r._wmsCode
}

// SetRemark is Remark Setter
// remark
func (r *TaobaoscitemaddAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoscitemaddAPIRequest) GetRemark() string {
	return r._remark
}

// SetBrandName is BrandName Setter
// brand_Name
func (r *TaobaoscitemaddAPIRequest) SetBrandName(_brandName string) error {
	r._brandName = _brandName
	r.Set("brand_name", _brandName)
	return nil
}

// GetBrandName BrandName Getter
func (r TaobaoscitemaddAPIRequest) GetBrandName() string {
	return r._brandName
}

// SetItemType is ItemType Setter
// 0.普通供应链商品 1.供应链组合主商品
func (r *TaobaoscitemaddAPIRequest) SetItemType(_itemType int64) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// GetItemType ItemType Getter
func (r TaobaoscitemaddAPIRequest) GetItemType() int64 {
	return r._itemType
}

// SetIsFriable is IsFriable Setter
// 是否易碎 0：不是  1：是
func (r *TaobaoscitemaddAPIRequest) SetIsFriable(_isFriable int64) error {
	r._isFriable = _isFriable
	r.Set("is_friable", _isFriable)
	return nil
}

// GetIsFriable IsFriable Getter
func (r TaobaoscitemaddAPIRequest) GetIsFriable() int64 {
	return r._isFriable
}

// SetIsDangerous is IsDangerous Setter
// 是否危险 0：不是  1：是
func (r *TaobaoscitemaddAPIRequest) SetIsDangerous(_isDangerous int64) error {
	r._isDangerous = _isDangerous
	r.Set("is_dangerous", _isDangerous)
	return nil
}

// GetIsDangerous IsDangerous Getter
func (r TaobaoscitemaddAPIRequest) GetIsDangerous() int64 {
	return r._isDangerous
}

// SetIsCostly is IsCostly Setter
// 是否是贵重品 0:不是 1：是
func (r *TaobaoscitemaddAPIRequest) SetIsCostly(_isCostly int64) error {
	r._isCostly = _isCostly
	r.Set("is_costly", _isCostly)
	return nil
}

// GetIsCostly IsCostly Getter
func (r TaobaoscitemaddAPIRequest) GetIsCostly() int64 {
	return r._isCostly
}

// SetIsWarranty is IsWarranty Setter
// 是否保质期：0:不是 1：是
func (r *TaobaoscitemaddAPIRequest) SetIsWarranty(_isWarranty int64) error {
	r._isWarranty = _isWarranty
	r.Set("is_warranty", _isWarranty)
	return nil
}

// GetIsWarranty IsWarranty Getter
func (r TaobaoscitemaddAPIRequest) GetIsWarranty() int64 {
	return r._isWarranty
}

// SetWeight is Weight Setter
// 重量 单位：g
func (r *TaobaoscitemaddAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r TaobaoscitemaddAPIRequest) GetWeight() int64 {
	return r._weight
}

// SetLength is Length Setter
// 长度 单位：mm
func (r *TaobaoscitemaddAPIRequest) SetLength(_length int64) error {
	r._length = _length
	r.Set("length", _length)
	return nil
}

// GetLength Length Getter
func (r TaobaoscitemaddAPIRequest) GetLength() int64 {
	return r._length
}

// SetWidth is Width Setter
// 宽 单位：mm
func (r *TaobaoscitemaddAPIRequest) SetWidth(_width int64) error {
	r._width = _width
	r.Set("width", _width)
	return nil
}

// GetWidth Width Getter
func (r TaobaoscitemaddAPIRequest) GetWidth() int64 {
	return r._width
}

// SetHeight is Height Setter
// 高 单位：mm
func (r *TaobaoscitemaddAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// GetHeight Height Getter
func (r TaobaoscitemaddAPIRequest) GetHeight() int64 {
	return r._height
}

// SetVolume is Volume Setter
// 体积：立方厘米
func (r *TaobaoscitemaddAPIRequest) SetVolume(_volume int64) error {
	r._volume = _volume
	r.Set("volume", _volume)
	return nil
}

// GetVolume Volume Getter
func (r TaobaoscitemaddAPIRequest) GetVolume() int64 {
	return r._volume
}

// SetPrice is Price Setter
// 价格 单位：分
func (r *TaobaoscitemaddAPIRequest) SetPrice(_price int64) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TaobaoscitemaddAPIRequest) GetPrice() int64 {
	return r._price
}

// SetMatterStatus is MatterStatus Setter
// 0:液体，1：粉体，2：固体
func (r *TaobaoscitemaddAPIRequest) SetMatterStatus(_matterStatus int64) error {
	r._matterStatus = _matterStatus
	r.Set("matter_status", _matterStatus)
	return nil
}

// GetMatterStatus MatterStatus Getter
func (r TaobaoscitemaddAPIRequest) GetMatterStatus() int64 {
	return r._matterStatus
}

// SetBrandId is BrandId Setter
// 品牌id
func (r *TaobaoscitemaddAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r TaobaoscitemaddAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// SetSpuId is SpuId Setter
// spuId或是cspuid
func (r *TaobaoscitemaddAPIRequest) SetSpuId(_spuId int64) error {
	r._spuId = _spuId
	r.Set("spu_id", _spuId)
	return nil
}

// GetSpuId SpuId Getter
func (r TaobaoscitemaddAPIRequest) GetSpuId() int64 {
	return r._spuId
}

// SetIsAreaSale is IsAreaSale Setter
// 1表示区域销售，0或是空是非区域销售
func (r *TaobaoscitemaddAPIRequest) SetIsAreaSale(_isAreaSale int64) error {
	r._isAreaSale = _isAreaSale
	r.Set("is_area_sale", _isAreaSale)
	return nil
}

// GetIsAreaSale IsAreaSale Getter
func (r TaobaoscitemaddAPIRequest) GetIsAreaSale() int64 {
	return r._isAreaSale
}
