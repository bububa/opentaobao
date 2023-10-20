package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScOptimusMaterialAPIRequest 淘宝客-服务商-物料精选 API请求
// taobao.tbk.sc.optimus.material
//
// 服务商使用。支持入参推广者对应的“推广位”和官方提供的“物料id”，获取指定物料信息和推广者对应的推广链接，还可入参用户信息提供智能推荐（需智能推荐请先前协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaoTbkScOptimusMaterialAPIRequest struct {
	model.Params
	// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	_deviceType string
	// 智能匹配-设备号加密类型：MD5，类型为OAID时不传
	_deviceEncrypt string
	// 智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值
	_deviceValue string
	// 内容专用-内容渠道信息
	_contentSource string
	// 商品ID，用于相似商品推荐
	_itemId string
	// 页大小，默认20，1~100
	_pageSize int64
	// 第几页，默认：1
	_pageNo int64
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneId int64
	// 官方提供的物料Id（详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a）
	_materialId int64
	// mm_xxx_xxx_xxx的第2段数字
	_siteId int64
	// 内容专用-内容详情ID
	_contentId int64
}

// NewTaobaoTbkScOptimusMaterialRequest 初始化TaobaoTbkScOptimusMaterialAPIRequest对象
func NewTaobaoTbkScOptimusMaterialRequest() *TaobaoTbkScOptimusMaterialAPIRequest {
	return &TaobaoTbkScOptimusMaterialAPIRequest{
		Params: model.NewParams(11),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkScOptimusMaterialAPIRequest) Reset() {
	r._deviceType = ""
	r._deviceEncrypt = ""
	r._deviceValue = ""
	r._contentSource = ""
	r._itemId = ""
	r._pageSize = 0
	r._pageNo = 0
	r._adzoneId = 0
	r._materialId = 0
	r._siteId = 0
	r._contentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.optimus.material"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceType is DeviceType Setter
// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
func (r *TaobaoTbkScOptimusMaterialAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetDeviceEncrypt is DeviceEncrypt Setter
// 智能匹配-设备号加密类型：MD5，类型为OAID时不传
func (r *TaobaoTbkScOptimusMaterialAPIRequest) SetDeviceEncrypt(_deviceEncrypt string) error {
	r._deviceEncrypt = _deviceEncrypt
	r.Set("device_encrypt", _deviceEncrypt)
	return nil
}

// GetDeviceEncrypt DeviceEncrypt Getter
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetDeviceEncrypt() string {
	return r._deviceEncrypt
}

// SetDeviceValue is DeviceValue Setter
// 智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值
func (r *TaobaoTbkScOptimusMaterialAPIRequest) SetDeviceValue(_deviceValue string) error {
	r._deviceValue = _deviceValue
	r.Set("device_value", _deviceValue)
	return nil
}

// GetDeviceValue DeviceValue Getter
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetDeviceValue() string {
	return r._deviceValue
}

// SetContentSource is ContentSource Setter
// 内容专用-内容渠道信息
func (r *TaobaoTbkScOptimusMaterialAPIRequest) SetContentSource(_contentSource string) error {
	r._contentSource = _contentSource
	r.Set("content_source", _contentSource)
	return nil
}

// GetContentSource ContentSource Getter
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetContentSource() string {
	return r._contentSource
}

// SetItemId is ItemId Setter
// 商品ID，用于相似商品推荐
func (r *TaobaoTbkScOptimusMaterialAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetItemId() string {
	return r._itemId
}

// SetPageSize is PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoTbkScOptimusMaterialAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 第几页，默认：1
func (r *TaobaoTbkScOptimusMaterialAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaoTbkScOptimusMaterialAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetMaterialId is MaterialId Setter
// 官方提供的物料Id（详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a）
func (r *TaobaoTbkScOptimusMaterialAPIRequest) SetMaterialId(_materialId int64) error {
	r._materialId = _materialId
	r.Set("material_id", _materialId)
	return nil
}

// GetMaterialId MaterialId Getter
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetMaterialId() int64 {
	return r._materialId
}

// SetSiteId is SiteId Setter
// mm_xxx_xxx_xxx的第2段数字
func (r *TaobaoTbkScOptimusMaterialAPIRequest) SetSiteId(_siteId int64) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetSiteId() int64 {
	return r._siteId
}

// SetContentId is ContentId Setter
// 内容专用-内容详情ID
func (r *TaobaoTbkScOptimusMaterialAPIRequest) SetContentId(_contentId int64) error {
	r._contentId = _contentId
	r.Set("content_id", _contentId)
	return nil
}

// GetContentId ContentId Getter
func (r TaobaoTbkScOptimusMaterialAPIRequest) GetContentId() int64 {
	return r._contentId
}

var poolTaobaoTbkScOptimusMaterialAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkScOptimusMaterialRequest()
	},
}

// GetTaobaoTbkScOptimusMaterialRequest 从 sync.Pool 获取 TaobaoTbkScOptimusMaterialAPIRequest
func GetTaobaoTbkScOptimusMaterialAPIRequest() *TaobaoTbkScOptimusMaterialAPIRequest {
	return poolTaobaoTbkScOptimusMaterialAPIRequest.Get().(*TaobaoTbkScOptimusMaterialAPIRequest)
}

// ReleaseTaobaoTbkScOptimusMaterialAPIRequest 将 TaobaoTbkScOptimusMaterialAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkScOptimusMaterialAPIRequest(v *TaobaoTbkScOptimusMaterialAPIRequest) {
	v.Reset()
	poolTaobaoTbkScOptimusMaterialAPIRequest.Put(v)
}
