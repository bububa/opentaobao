package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelDeleteAPIRequest 删除酒店接口 API请求
// taobao.xhotel.delete
//
// 删除飞猪酒店数据接口
type TaobaoXhotelDeleteAPIRequest struct {
	model.Params
	// 酒店vendor
	_vendor string
	// 酒店编码
	_outerId string
	// 酒店id，传参方式  hid   或者   outer_id+vendor
	_hid int64
}

// NewTaobaoXhotelDeleteRequest 初始化TaobaoXhotelDeleteAPIRequest对象
func NewTaobaoXhotelDeleteRequest() *TaobaoXhotelDeleteAPIRequest {
	return &TaobaoXhotelDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelDeleteAPIRequest) Reset() {
	r._vendor = ""
	r._outerId = ""
	r._hid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 酒店vendor
func (r *TaobaoXhotelDeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelDeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOuterId is OuterId Setter
// 酒店编码
func (r *TaobaoXhotelDeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoXhotelDeleteAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetHid is Hid Setter
// 酒店id，传参方式  hid   或者   outer_id+vendor
func (r *TaobaoXhotelDeleteAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelDeleteAPIRequest) GetHid() int64 {
	return r._hid
}

var poolTaobaoXhotelDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelDeleteRequest()
	},
}

// GetTaobaoXhotelDeleteRequest 从 sync.Pool 获取 TaobaoXhotelDeleteAPIRequest
func GetTaobaoXhotelDeleteAPIRequest() *TaobaoXhotelDeleteAPIRequest {
	return poolTaobaoXhotelDeleteAPIRequest.Get().(*TaobaoXhotelDeleteAPIRequest)
}

// ReleaseTaobaoXhotelDeleteAPIRequest 将 TaobaoXhotelDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelDeleteAPIRequest(v *TaobaoXhotelDeleteAPIRequest) {
	v.Reset()
	poolTaobaoXhotelDeleteAPIRequest.Put(v)
}
