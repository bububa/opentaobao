package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryLaunchExtraChargeAPIRequest 阿里巴巴.行业.增加费用.服务商发起 API请求
// alibaba.ascp.industry.launch.extra.charge
//
// 阿里巴巴.行业.增加费用.服务商发起
type AlibabaAscpIndustryLaunchExtraChargeAPIRequest struct {
	model.Params
	// 请求对象
	_omsLaunchExtraChargeParameter *OmsLaunchExtraChargeParameter
}

// NewAlibabaAscpIndustryLaunchExtraChargeRequest 初始化AlibabaAscpIndustryLaunchExtraChargeAPIRequest对象
func NewAlibabaAscpIndustryLaunchExtraChargeRequest() *AlibabaAscpIndustryLaunchExtraChargeAPIRequest {
	return &AlibabaAscpIndustryLaunchExtraChargeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryLaunchExtraChargeAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.launch.extra.charge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryLaunchExtraChargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpIndustryLaunchExtraChargeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOmsLaunchExtraChargeParameter is OmsLaunchExtraChargeParameter Setter
// 请求对象
func (r *AlibabaAscpIndustryLaunchExtraChargeAPIRequest) SetOmsLaunchExtraChargeParameter(_omsLaunchExtraChargeParameter *OmsLaunchExtraChargeParameter) error {
	r._omsLaunchExtraChargeParameter = _omsLaunchExtraChargeParameter
	r.Set("oms_launch_extra_charge_parameter", _omsLaunchExtraChargeParameter)
	return nil
}

// GetOmsLaunchExtraChargeParameter OmsLaunchExtraChargeParameter Getter
func (r AlibabaAscpIndustryLaunchExtraChargeAPIRequest) GetOmsLaunchExtraChargeParameter() *OmsLaunchExtraChargeParameter {
	return r._omsLaunchExtraChargeParameter
}
