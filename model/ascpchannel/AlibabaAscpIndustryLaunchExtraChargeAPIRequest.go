package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustrylaunchextrachargeAPIRequest 阿里巴巴.行业.增加费用.服务商发起 API请求
// alibaba.ascp.industry.launch.extra.charge
//
// 阿里巴巴.行业.增加费用.服务商发起
type AlibabaascpindustrylaunchextrachargeAPIRequest struct {
	model.Params
	// 请求对象
	_omsLaunchExtraChargeParameter *OmsLaunchExtraChargeParameter
}

// NewAlibabaascpindustrylaunchextrachargeRequest 初始化AlibabaascpindustrylaunchextrachargeAPIRequest对象
func NewAlibabaascpindustrylaunchextrachargeRequest() *AlibabaascpindustrylaunchextrachargeAPIRequest {
	return &AlibabaascpindustrylaunchextrachargeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpindustrylaunchextrachargeAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.launch.extra.charge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpindustrylaunchextrachargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpindustrylaunchextrachargeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOmsLaunchExtraChargeParameter is OmsLaunchExtraChargeParameter Setter
// 请求对象
func (r *AlibabaascpindustrylaunchextrachargeAPIRequest) SetOmsLaunchExtraChargeParameter(_omsLaunchExtraChargeParameter *OmsLaunchExtraChargeParameter) error {
	r._omsLaunchExtraChargeParameter = _omsLaunchExtraChargeParameter
	r.Set("oms_launch_extra_charge_parameter", _omsLaunchExtraChargeParameter)
	return nil
}

// GetOmsLaunchExtraChargeParameter OmsLaunchExtraChargeParameter Getter
func (r AlibabaascpindustrylaunchextrachargeAPIRequest) GetOmsLaunchExtraChargeParameter() *OmsLaunchExtraChargeParameter {
	return r._omsLaunchExtraChargeParameter
}
