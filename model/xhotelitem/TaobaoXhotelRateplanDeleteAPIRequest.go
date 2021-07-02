package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateplanDeleteAPIRequest 价格计划rateplan删除 API请求
// taobao.xhotel.rateplan.delete
//
// 酒店产品库rateplan删除，同时删除级联的rate，请谨慎使用
type TaobaoXhotelRateplanDeleteAPIRequest struct {
	model.Params
	// ratepland标识
	_rpId int64
	// 系统商，一般不用填写，使用须申请
	_vendor string
	// 商家价格政策编码
	_rateplanCode string
}

// NewTaobaoXhotelRateplanDeleteRequest 初始化TaobaoXhotelRateplanDeleteAPIRequest对象
func NewTaobaoXhotelRateplanDeleteRequest() *TaobaoXhotelRateplanDeleteAPIRequest {
	return &TaobaoXhotelRateplanDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateplanDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rateplan.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateplanDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRpId is RpId Setter
// ratepland标识
func (r *TaobaoXhotelRateplanDeleteAPIRequest) SetRpId(_rpId int64) error {
	r._rpId = _rpId
	r.Set("rp_id", _rpId)
	return nil
}

// GetRpId RpId Getter
func (r TaobaoXhotelRateplanDeleteAPIRequest) GetRpId() int64 {
	return r._rpId
}

// SetVendor is Vendor Setter
// 系统商，一般不用填写，使用须申请
func (r *TaobaoXhotelRateplanDeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelRateplanDeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRateplanCode is RateplanCode Setter
// 商家价格政策编码
func (r *TaobaoXhotelRateplanDeleteAPIRequest) SetRateplanCode(_rateplanCode string) error {
	r._rateplanCode = _rateplanCode
	r.Set("rateplan_code", _rateplanCode)
	return nil
}

// GetRateplanCode RateplanCode Getter
func (r TaobaoXhotelRateplanDeleteAPIRequest) GetRateplanCode() string {
	return r._rateplanCode
}
