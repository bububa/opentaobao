package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelDeleteAPIRequest 删除酒店接口 API请求
// taobao.xhotel.delete
//
// 删除飞猪酒店数据接口
type TaobaoXhotelDeleteAPIRequest struct {
	model.Params
	// 酒店id，传参方式  hid   或者   outer_id+vendor
	_hid int64
	// 酒店vendor
	_vendor string
	// 酒店编码
	_outerId string
}

// NewTaobaoXhotelDeleteRequest 初始化TaobaoXhotelDeleteAPIRequest对象
func NewTaobaoXhotelDeleteRequest() *TaobaoXhotelDeleteAPIRequest {
	return &TaobaoXhotelDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Hid Setter
// 酒店id，传参方式  hid   或者   outer_id+vendor
func (r *TaobaoXhotelDeleteAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// Get Hid Getter
func (r TaobaoXhotelDeleteAPIRequest) GetHid() int64 {
	return r._hid
}

// Set is Vendor Setter
// 酒店vendor
func (r *TaobaoXhotelDeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// Get Vendor Getter
func (r TaobaoXhotelDeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// Set is OuterId Setter
// 酒店编码
func (r *TaobaoXhotelDeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r TaobaoXhotelDeleteAPIRequest) GetOuterId() string {
	return r._outerId
}
