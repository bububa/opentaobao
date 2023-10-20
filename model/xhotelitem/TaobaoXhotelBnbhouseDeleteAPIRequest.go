package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbhousedeleteAPIRequest 民宿门店删除接口 API请求
// taobao.xhotel.bnbhouse.delete
//
// 支持门店相关的门店删除，删除门店会级联删除门店下面的房源
type TaobaoxhotelbnbhousedeleteAPIRequest struct {
	model.Params
	// 对接系统商名称：可为空不要乱填，需要申请后使用
	_vendor string
	// 门店Id，系统商outer_id
	_outerId string
	// 门店Id，传参方式为hid或outer_id+vendor
	_hid int64
}

// NewTaobaoxhotelbnbhousedeleteRequest 初始化TaobaoxhotelbnbhousedeleteAPIRequest对象
func NewTaobaoxhotelbnbhousedeleteRequest() *TaobaoxhotelbnbhousedeleteAPIRequest {
	return &TaobaoxhotelbnbhousedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbnbhousedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbhouse.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbnbhousedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbnbhousedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 对接系统商名称：可为空不要乱填，需要申请后使用
func (r *TaobaoxhotelbnbhousedeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelbnbhousedeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOuterId is OuterId Setter
// 门店Id，系统商outer_id
func (r *TaobaoxhotelbnbhousedeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoxhotelbnbhousedeleteAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetHid is Hid Setter
// 门店Id，传参方式为hid或outer_id+vendor
func (r *TaobaoxhotelbnbhousedeleteAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoxhotelbnbhousedeleteAPIRequest) GetHid() int64 {
	return r._hid
}
