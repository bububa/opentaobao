package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgoptimusmaterialAPIRequest 淘宝客-推广者-物料精选 API请求
// taobao.tbk.dg.optimus.material
//
// 支持入参对应的“推广位”和官方提供的“物料id”，获取指定物料信息和推广链接，还可入参用户信息提供智能推荐（需智能推荐请先前协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaotbkdgoptimusmaterialAPIRequest struct {
	model.Params
	// 智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值
	_deviceValue string
	// 智能匹配-设备号加密类型：MD5，类型为OAID时不传
	_deviceEncrypt string
	// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	_deviceType string
	// 内容专用-内容渠道信息
	_contentSource string
	// 商品ID，用于相似商品推荐
	_itemId string
	// 选品库投放id
	_favoritesId string
	// 页大小，默认20，1~100
	_pageSize int64
	// 第几页，默认：1
	_pageNo int64
	// mm_xxx_xxx_xxx的第三位
	_adzoneId int64
	// 官方的物料Id(详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a)
	_materialId int64
	// 内容专用-内容详情ID
	_contentId int64
}

// NewTaobaotbkdgoptimusmaterialRequest 初始化TaobaotbkdgoptimusmaterialAPIRequest对象
func NewTaobaotbkdgoptimusmaterialRequest() *TaobaotbkdgoptimusmaterialAPIRequest {
	return &TaobaotbkdgoptimusmaterialAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.optimus.material"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceValue is DeviceValue Setter
// 智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值
func (r *TaobaotbkdgoptimusmaterialAPIRequest) SetDeviceValue(_deviceValue string) error {
	r._deviceValue = _deviceValue
	r.Set("device_value", _deviceValue)
	return nil
}

// GetDeviceValue DeviceValue Getter
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetDeviceValue() string {
	return r._deviceValue
}

// SetDeviceEncrypt is DeviceEncrypt Setter
// 智能匹配-设备号加密类型：MD5，类型为OAID时不传
func (r *TaobaotbkdgoptimusmaterialAPIRequest) SetDeviceEncrypt(_deviceEncrypt string) error {
	r._deviceEncrypt = _deviceEncrypt
	r.Set("device_encrypt", _deviceEncrypt)
	return nil
}

// GetDeviceEncrypt DeviceEncrypt Getter
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetDeviceEncrypt() string {
	return r._deviceEncrypt
}

// SetDeviceType is DeviceType Setter
// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
func (r *TaobaotbkdgoptimusmaterialAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetContentSource is ContentSource Setter
// 内容专用-内容渠道信息
func (r *TaobaotbkdgoptimusmaterialAPIRequest) SetContentSource(_contentSource string) error {
	r._contentSource = _contentSource
	r.Set("content_source", _contentSource)
	return nil
}

// GetContentSource ContentSource Getter
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetContentSource() string {
	return r._contentSource
}

// SetItemId is ItemId Setter
// 商品ID，用于相似商品推荐
func (r *TaobaotbkdgoptimusmaterialAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetItemId() string {
	return r._itemId
}

// SetFavoritesId is FavoritesId Setter
// 选品库投放id
func (r *TaobaotbkdgoptimusmaterialAPIRequest) SetFavoritesId(_favoritesId string) error {
	r._favoritesId = _favoritesId
	r.Set("favorites_id", _favoritesId)
	return nil
}

// GetFavoritesId FavoritesId Getter
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetFavoritesId() string {
	return r._favoritesId
}

// SetPageSize is PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaotbkdgoptimusmaterialAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 第几页，默认：1
func (r *TaobaotbkdgoptimusmaterialAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaotbkdgoptimusmaterialAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetMaterialId is MaterialId Setter
// 官方的物料Id(详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a)
func (r *TaobaotbkdgoptimusmaterialAPIRequest) SetMaterialId(_materialId int64) error {
	r._materialId = _materialId
	r.Set("material_id", _materialId)
	return nil
}

// GetMaterialId MaterialId Getter
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetMaterialId() int64 {
	return r._materialId
}

// SetContentId is ContentId Setter
// 内容专用-内容详情ID
func (r *TaobaotbkdgoptimusmaterialAPIRequest) SetContentId(_contentId int64) error {
	r._contentId = _contentId
	r.Set("content_id", _contentId)
	return nil
}

// GetContentId ContentId Getter
func (r TaobaotbkdgoptimusmaterialAPIRequest) GetContentId() int64 {
	return r._contentId
}
