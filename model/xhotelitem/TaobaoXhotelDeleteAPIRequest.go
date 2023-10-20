package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhoteldeleteAPIRequest 删除酒店接口 API请求
// taobao.xhotel.delete
//
// 删除飞猪酒店数据接口
type TaobaoxhoteldeleteAPIRequest struct {
	model.Params
	// 酒店vendor
	_vendor string
	// 酒店编码
	_outerId string
	// 酒店id，传参方式  hid   或者   outer_id+vendor
	_hid int64
}

// NewTaobaoxhoteldeleteRequest 初始化TaobaoxhoteldeleteAPIRequest对象
func NewTaobaoxhoteldeleteRequest() *TaobaoxhoteldeleteAPIRequest {
	return &TaobaoxhoteldeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhoteldeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhoteldeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhoteldeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 酒店vendor
func (r *TaobaoxhoteldeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhoteldeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOuterId is OuterId Setter
// 酒店编码
func (r *TaobaoxhoteldeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoxhoteldeleteAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetHid is Hid Setter
// 酒店id，传参方式  hid   或者   outer_id+vendor
func (r *TaobaoxhoteldeleteAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoxhoteldeleteAPIRequest) GetHid() int64 {
	return r._hid
}
