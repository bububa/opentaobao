package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBaseinfoRoomGetAPIRequest 酒店房型与房价查询 API请求
// taobao.xhotel.baseinfo.room.get
//
// 根据outHid/hid获取酒店的房型和价格信息
type TaobaoXhotelBaseinfoRoomGetAPIRequest struct {
	model.Params
	// 卖家系统中的酒店ID。
	_outHid string
	// 用于标示该酒店发布的渠道信息
	_vendor string
	// 是否需要房价基本信息（false为不需要），默认为需要
	_isNeedRatePlan bool
}

// NewTaobaoXhotelBaseinfoRoomGetRequest 初始化TaobaoXhotelBaseinfoRoomGetAPIRequest对象
func NewTaobaoXhotelBaseinfoRoomGetRequest() *TaobaoXhotelBaseinfoRoomGetAPIRequest {
	return &TaobaoXhotelBaseinfoRoomGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBaseinfoRoomGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.baseinfo.room.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBaseinfoRoomGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOutHid is OutHid Setter
// 卖家系统中的酒店ID。
func (r *TaobaoXhotelBaseinfoRoomGetAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoXhotelBaseinfoRoomGetAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetVendor is Vendor Setter
// 用于标示该酒店发布的渠道信息
func (r *TaobaoXhotelBaseinfoRoomGetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelBaseinfoRoomGetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetIsNeedRatePlan is IsNeedRatePlan Setter
// 是否需要房价基本信息（false为不需要），默认为需要
func (r *TaobaoXhotelBaseinfoRoomGetAPIRequest) SetIsNeedRatePlan(_isNeedRatePlan bool) error {
	r._isNeedRatePlan = _isNeedRatePlan
	r.Set("is_need_rate_plan", _isNeedRatePlan)
	return nil
}

// GetIsNeedRatePlan IsNeedRatePlan Getter
func (r TaobaoXhotelBaseinfoRoomGetAPIRequest) GetIsNeedRatePlan() bool {
	return r._isNeedRatePlan
}
