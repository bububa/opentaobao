package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbhouseDeleteAPIRequest 民宿门店删除接口 API请求
// taobao.xhotel.bnbhouse.delete
//
// 支持门店相关的门店删除，删除门店会级联删除门店下面的房源
type TaobaoXhotelBnbhouseDeleteAPIRequest struct {
	model.Params
	// 对接系统商名称：可为空不要乱填，需要申请后使用
	_vendor string
	// 门店Id，系统商outer_id
	_outerId string
	// 门店Id，传参方式为hid或outer_id+vendor
	_hid int64
}

// NewTaobaoXhotelBnbhouseDeleteRequest 初始化TaobaoXhotelBnbhouseDeleteAPIRequest对象
func NewTaobaoXhotelBnbhouseDeleteRequest() *TaobaoXhotelBnbhouseDeleteAPIRequest {
	return &TaobaoXhotelBnbhouseDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelBnbhouseDeleteAPIRequest) Reset() {
	r._vendor = ""
	r._outerId = ""
	r._hid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbhouseDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbhouse.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbhouseDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelBnbhouseDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 对接系统商名称：可为空不要乱填，需要申请后使用
func (r *TaobaoXhotelBnbhouseDeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelBnbhouseDeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOuterId is OuterId Setter
// 门店Id，系统商outer_id
func (r *TaobaoXhotelBnbhouseDeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoXhotelBnbhouseDeleteAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetHid is Hid Setter
// 门店Id，传参方式为hid或outer_id+vendor
func (r *TaobaoXhotelBnbhouseDeleteAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelBnbhouseDeleteAPIRequest) GetHid() int64 {
	return r._hid
}

var poolTaobaoXhotelBnbhouseDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelBnbhouseDeleteRequest()
	},
}

// GetTaobaoXhotelBnbhouseDeleteRequest 从 sync.Pool 获取 TaobaoXhotelBnbhouseDeleteAPIRequest
func GetTaobaoXhotelBnbhouseDeleteAPIRequest() *TaobaoXhotelBnbhouseDeleteAPIRequest {
	return poolTaobaoXhotelBnbhouseDeleteAPIRequest.Get().(*TaobaoXhotelBnbhouseDeleteAPIRequest)
}

// ReleaseTaobaoXhotelBnbhouseDeleteAPIRequest 将 TaobaoXhotelBnbhouseDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelBnbhouseDeleteAPIRequest(v *TaobaoXhotelBnbhouseDeleteAPIRequest) {
	v.Reset()
	poolTaobaoXhotelBnbhouseDeleteAPIRequest.Put(v)
}
