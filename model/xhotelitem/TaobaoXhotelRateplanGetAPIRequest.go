package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateplanGetAPIRequest 价格计划rateplan查询 API请求
// taobao.xhotel.rateplan.get
//
// 酒店产品库rateplan查询
type TaobaoXhotelRateplanGetAPIRequest struct {
	model.Params
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 卖家自己系统的Code，简称RateCode
	_rateplanCode string
	// 废弃，使用rateplan_code
	_rpid int64
}

// NewTaobaoXhotelRateplanGetRequest 初始化TaobaoXhotelRateplanGetAPIRequest对象
func NewTaobaoXhotelRateplanGetRequest() *TaobaoXhotelRateplanGetAPIRequest {
	return &TaobaoXhotelRateplanGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateplanGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rateplan.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateplanGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRateplanGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelRateplanGetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelRateplanGetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRateplanCode is RateplanCode Setter
// 卖家自己系统的Code，简称RateCode
func (r *TaobaoXhotelRateplanGetAPIRequest) SetRateplanCode(_rateplanCode string) error {
	r._rateplanCode = _rateplanCode
	r.Set("rateplan_code", _rateplanCode)
	return nil
}

// GetRateplanCode RateplanCode Getter
func (r TaobaoXhotelRateplanGetAPIRequest) GetRateplanCode() string {
	return r._rateplanCode
}

// SetRpid is Rpid Setter
// 废弃，使用rateplan_code
func (r *TaobaoXhotelRateplanGetAPIRequest) SetRpid(_rpid int64) error {
	r._rpid = _rpid
	r.Set("rpid", _rpid)
	return nil
}

// GetRpid Rpid Getter
func (r TaobaoXhotelRateplanGetAPIRequest) GetRpid() int64 {
	return r._rpid
}
