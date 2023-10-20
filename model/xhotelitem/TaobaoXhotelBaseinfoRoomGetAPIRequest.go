package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbaseinforoomgetAPIRequest 酒店房型与房价查询 API请求
// taobao.xhotel.baseinfo.room.get
//
// 根据outHid/hid获取酒店的房型和价格信息
type TaobaoxhotelbaseinforoomgetAPIRequest struct {
	model.Params
	// 卖家系统中的酒店ID。
	_outHid string
	// 用于标示该酒店发布的渠道信息
	_vendor string
	// 是否需要房价基本信息（false为不需要），默认为需要
	_isNeedRatePlan bool
}

// NewTaobaoxhotelbaseinforoomgetRequest 初始化TaobaoxhotelbaseinforoomgetAPIRequest对象
func NewTaobaoxhotelbaseinforoomgetRequest() *TaobaoxhotelbaseinforoomgetAPIRequest {
	return &TaobaoxhotelbaseinforoomgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbaseinforoomgetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.baseinfo.room.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbaseinforoomgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbaseinforoomgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutHid is OutHid Setter
// 卖家系统中的酒店ID。
func (r *TaobaoxhotelbaseinforoomgetAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoxhotelbaseinforoomgetAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetVendor is Vendor Setter
// 用于标示该酒店发布的渠道信息
func (r *TaobaoxhotelbaseinforoomgetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelbaseinforoomgetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetIsNeedRatePlan is IsNeedRatePlan Setter
// 是否需要房价基本信息（false为不需要），默认为需要
func (r *TaobaoxhotelbaseinforoomgetAPIRequest) SetIsNeedRatePlan(_isNeedRatePlan bool) error {
	r._isNeedRatePlan = _isNeedRatePlan
	r.Set("is_need_rate_plan", _isNeedRatePlan)
	return nil
}

// GetIsNeedRatePlan IsNeedRatePlan Getter
func (r TaobaoxhotelbaseinforoomgetAPIRequest) GetIsNeedRatePlan() bool {
	return r._isNeedRatePlan
}
