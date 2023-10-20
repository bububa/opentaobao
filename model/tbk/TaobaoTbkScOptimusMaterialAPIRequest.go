package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscoptimusmaterialAPIRequest 淘宝客-服务商-物料精选 API请求
// taobao.tbk.sc.optimus.material
//
// 服务商使用。支持入参推广者对应的“推广位”和官方提供的“物料id”，获取指定物料信息和推广者对应的推广链接，还可入参用户信息提供智能推荐（需智能推荐请先前协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaotbkscoptimusmaterialAPIRequest struct {
	model.Params
	// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	_devicetype string
	// 智能匹配-设备号加密类型：MD5，类型为OAID时不传
	_deviceencrypt string
	// 智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值
	_devicevalue string
	// 内容专用-内容渠道信息
	_contentsource string
	// 商品ID，用于相似商品推荐
	_itemid string
	// 页大小，默认20，1~100
	_pagesize int64
	// 第几页，默认：1
	_pageno int64
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneid int64
	// 官方提供的物料Id（详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a）
	_materialid int64
	// mm_xxx_xxx_xxx的第2段数字
	_siteid int64
	// 内容专用-内容详情ID
	_contentid int64
}

// NewTaobaotbkscoptimusmaterialRequest 初始化TaobaotbkscoptimusmaterialAPIRequest对象
func NewTaobaotbkscoptimusmaterialRequest() *TaobaotbkscoptimusmaterialAPIRequest {
	return &TaobaotbkscoptimusmaterialAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscoptimusmaterialAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.optimus.material"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscoptimusmaterialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscoptimusmaterialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDevicetype is Devicetype Setter
// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
func (r *TaobaotbkscoptimusmaterialAPIRequest) SetDevicetype(_devicetype string) error {
	r._devicetype = _devicetype
	r.Set("device_type", _devicetype)
	return nil
}

// GetDevicetype Devicetype Getter
func (r TaobaotbkscoptimusmaterialAPIRequest) GetDevicetype() string {
	return r._devicetype
}

// SetDeviceencrypt is Deviceencrypt Setter
// 智能匹配-设备号加密类型：MD5，类型为OAID时不传
func (r *TaobaotbkscoptimusmaterialAPIRequest) SetDeviceencrypt(_deviceencrypt string) error {
	r._deviceencrypt = _deviceencrypt
	r.Set("device_encrypt", _deviceencrypt)
	return nil
}

// GetDeviceencrypt Deviceencrypt Getter
func (r TaobaotbkscoptimusmaterialAPIRequest) GetDeviceencrypt() string {
	return r._deviceencrypt
}

// SetDevicevalue is Devicevalue Setter
// 智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值
func (r *TaobaotbkscoptimusmaterialAPIRequest) SetDevicevalue(_devicevalue string) error {
	r._devicevalue = _devicevalue
	r.Set("device_value", _devicevalue)
	return nil
}

// GetDevicevalue Devicevalue Getter
func (r TaobaotbkscoptimusmaterialAPIRequest) GetDevicevalue() string {
	return r._devicevalue
}

// SetContentsource is Contentsource Setter
// 内容专用-内容渠道信息
func (r *TaobaotbkscoptimusmaterialAPIRequest) SetContentsource(_contentsource string) error {
	r._contentsource = _contentsource
	r.Set("content_source", _contentsource)
	return nil
}

// GetContentsource Contentsource Getter
func (r TaobaotbkscoptimusmaterialAPIRequest) GetContentsource() string {
	return r._contentsource
}

// SetItemid is Itemid Setter
// 商品ID，用于相似商品推荐
func (r *TaobaotbkscoptimusmaterialAPIRequest) SetItemid(_itemid string) error {
	r._itemid = _itemid
	r.Set("item_id", _itemid)
	return nil
}

// GetItemid Itemid Getter
func (r TaobaotbkscoptimusmaterialAPIRequest) GetItemid() string {
	return r._itemid
}

// SetPagesize is Pagesize Setter
// 页大小，默认20，1~100
func (r *TaobaotbkscoptimusmaterialAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkscoptimusmaterialAPIRequest) GetPagesize() int64 {
	return r._pagesize
}

// SetPageno is Pageno Setter
// 第几页，默认：1
func (r *TaobaotbkscoptimusmaterialAPIRequest) SetPageno(_pageno int64) error {
	r._pageno = _pageno
	r.Set("page_no", _pageno)
	return nil
}

// GetPageno Pageno Getter
func (r TaobaotbkscoptimusmaterialAPIRequest) GetPageno() int64 {
	return r._pageno
}

// SetAdzoneid is Adzoneid Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaotbkscoptimusmaterialAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkscoptimusmaterialAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}

// SetMaterialid is Materialid Setter
// 官方提供的物料Id（详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a）
func (r *TaobaotbkscoptimusmaterialAPIRequest) SetMaterialid(_materialid int64) error {
	r._materialid = _materialid
	r.Set("material_id", _materialid)
	return nil
}

// GetMaterialid Materialid Getter
func (r TaobaotbkscoptimusmaterialAPIRequest) GetMaterialid() int64 {
	return r._materialid
}

// SetSiteid is Siteid Setter
// mm_xxx_xxx_xxx的第2段数字
func (r *TaobaotbkscoptimusmaterialAPIRequest) SetSiteid(_siteid int64) error {
	r._siteid = _siteid
	r.Set("site_id", _siteid)
	return nil
}

// GetSiteid Siteid Getter
func (r TaobaotbkscoptimusmaterialAPIRequest) GetSiteid() int64 {
	return r._siteid
}

// SetContentid is Contentid Setter
// 内容专用-内容详情ID
func (r *TaobaotbkscoptimusmaterialAPIRequest) SetContentid(_contentid int64) error {
	r._contentid = _contentid
	r.Set("content_id", _contentid)
	return nil
}

// GetContentid Contentid Getter
func (r TaobaotbkscoptimusmaterialAPIRequest) GetContentid() int64 {
	return r._contentid
}
