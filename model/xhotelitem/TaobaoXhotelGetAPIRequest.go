package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelgetAPIRequest 酒店查询接口 API请求
// taobao.xhotel.get
//
// 酒店查询接口
type TaobaoxhotelgetAPIRequest struct {
	model.Params
	// 卖家系统中的酒店ID。
	_outerId string
	// 系统商，一般不用填写，使用须申请
	_vendor string
	// 废弃，请使用outer_id
	_hid int64
	// 是否需要在售状态(默认false不需要)
	_needSaleInfo bool
}

// NewTaobaoxhotelgetRequest 初始化TaobaoxhotelgetAPIRequest对象
func NewTaobaoxhotelgetRequest() *TaobaoxhotelgetAPIRequest {
	return &TaobaoxhotelgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelgetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 卖家系统中的酒店ID。
func (r *TaobaoxhotelgetAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoxhotelgetAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetVendor is Vendor Setter
// 系统商，一般不用填写，使用须申请
func (r *TaobaoxhotelgetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelgetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetHid is Hid Setter
// 废弃，请使用outer_id
func (r *TaobaoxhotelgetAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoxhotelgetAPIRequest) GetHid() int64 {
	return r._hid
}

// SetNeedSaleInfo is NeedSaleInfo Setter
// 是否需要在售状态(默认false不需要)
func (r *TaobaoxhotelgetAPIRequest) SetNeedSaleInfo(_needSaleInfo bool) error {
	r._needSaleInfo = _needSaleInfo
	r.Set("need_sale_info", _needSaleInfo)
	return nil
}

// GetNeedSaleInfo NeedSaleInfo Getter
func (r TaobaoxhotelgetAPIRequest) GetNeedSaleInfo() bool {
	return r._needSaleInfo
}
