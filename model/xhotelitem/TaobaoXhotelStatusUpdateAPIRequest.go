package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelStatusUpdateAPIRequest 酒店状态更新 API请求
// taobao.xhotel.status.update
//
// 酒店状态更新
type TaobaoXhotelStatusUpdateAPIRequest struct {
	model.Params
	// 对接系统商，不填默认taobao
	_vendor string
	// 外部酒店id
	_outerId string
	// 飞猪酒店id
	_hid int64
	// （0在线，-1删除， -2停售）
	_status int64
}

// NewTaobaoXhotelStatusUpdateRequest 初始化TaobaoXhotelStatusUpdateAPIRequest对象
func NewTaobaoXhotelStatusUpdateRequest() *TaobaoXhotelStatusUpdateAPIRequest {
	return &TaobaoXhotelStatusUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelStatusUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelStatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelStatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 对接系统商，不填默认taobao
func (r *TaobaoXhotelStatusUpdateAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelStatusUpdateAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOuterId is OuterId Setter
// 外部酒店id
func (r *TaobaoXhotelStatusUpdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoXhotelStatusUpdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetHid is Hid Setter
// 飞猪酒店id
func (r *TaobaoXhotelStatusUpdateAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelStatusUpdateAPIRequest) GetHid() int64 {
	return r._hid
}

// SetStatus is Status Setter
// （0在线，-1删除， -2停售）
func (r *TaobaoXhotelStatusUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoXhotelStatusUpdateAPIRequest) GetStatus() int64 {
	return r._status
}
