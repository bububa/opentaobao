package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoroomtypestatusupdateAPIRequest top房型状态修改 API请求
// taobao.roomtype.status.update
//
// top房型状态修改
type TaobaoroomtypestatusupdateAPIRequest struct {
	model.Params
	// 对接系统商，不填默认taobao
	_vendor string
	// 卖家房型id
	_outerId string
	// 飞猪房型id
	_rid int64
	// 0在线，-1删除， -2停售）
	_status int64
}

// NewTaobaoroomtypestatusupdateRequest 初始化TaobaoroomtypestatusupdateAPIRequest对象
func NewTaobaoroomtypestatusupdateRequest() *TaobaoroomtypestatusupdateAPIRequest {
	return &TaobaoroomtypestatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoroomtypestatusupdateAPIRequest) GetApiMethodName() string {
	return "taobao.roomtype.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoroomtypestatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoroomtypestatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 对接系统商，不填默认taobao
func (r *TaobaoroomtypestatusupdateAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoroomtypestatusupdateAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOuterId is OuterId Setter
// 卖家房型id
func (r *TaobaoroomtypestatusupdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoroomtypestatusupdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetRid is Rid Setter
// 飞猪房型id
func (r *TaobaoroomtypestatusupdateAPIRequest) SetRid(_rid int64) error {
	r._rid = _rid
	r.Set("rid", _rid)
	return nil
}

// GetRid Rid Getter
func (r TaobaoroomtypestatusupdateAPIRequest) GetRid() int64 {
	return r._rid
}

// SetStatus is Status Setter
// 0在线，-1删除， -2停售）
func (r *TaobaoroomtypestatusupdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoroomtypestatusupdateAPIRequest) GetStatus() int64 {
	return r._status
}
