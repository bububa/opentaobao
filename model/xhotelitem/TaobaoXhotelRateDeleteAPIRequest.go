package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelratedeleteAPIRequest rate删除接口 API请求
// taobao.xhotel.rate.delete
//
// 酒店产品库rate删除
type TaobaoxhotelratedeleteAPIRequest struct {
	model.Params
	// 系统商，一般不用填写，使用须申请
	_vendor string
	// 商家价格政策编码
	_rateplanCode string
	// 商家房型ID
	_outRid string
}

// NewTaobaoxhotelratedeleteRequest 初始化TaobaoxhotelratedeleteAPIRequest对象
func NewTaobaoxhotelratedeleteRequest() *TaobaoxhotelratedeleteAPIRequest {
	return &TaobaoxhotelratedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelratedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rate.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelratedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelratedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 系统商，一般不用填写，使用须申请
func (r *TaobaoxhotelratedeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelratedeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRateplanCode is RateplanCode Setter
// 商家价格政策编码
func (r *TaobaoxhotelratedeleteAPIRequest) SetRateplanCode(_rateplanCode string) error {
	r._rateplanCode = _rateplanCode
	r.Set("rateplan_code", _rateplanCode)
	return nil
}

// GetRateplanCode RateplanCode Getter
func (r TaobaoxhotelratedeleteAPIRequest) GetRateplanCode() string {
	return r._rateplanCode
}

// SetOutRid is OutRid Setter
// 商家房型ID
func (r *TaobaoxhotelratedeleteAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoxhotelratedeleteAPIRequest) GetOutRid() string {
	return r._outRid
}
