package xhotelitem

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelStatusUpdateAPIRequest) Reset() {
	r._vendor = ""
	r._outerId = ""
	r._hid = 0
	r._status = 0
	r.Params.ToZero()
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

var poolTaobaoXhotelStatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelStatusUpdateRequest()
	},
}

// GetTaobaoXhotelStatusUpdateRequest 从 sync.Pool 获取 TaobaoXhotelStatusUpdateAPIRequest
func GetTaobaoXhotelStatusUpdateAPIRequest() *TaobaoXhotelStatusUpdateAPIRequest {
	return poolTaobaoXhotelStatusUpdateAPIRequest.Get().(*TaobaoXhotelStatusUpdateAPIRequest)
}

// ReleaseTaobaoXhotelStatusUpdateAPIRequest 将 TaobaoXhotelStatusUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelStatusUpdateAPIRequest(v *TaobaoXhotelStatusUpdateAPIRequest) {
	v.Reset()
	poolTaobaoXhotelStatusUpdateAPIRequest.Put(v)
}
