package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustryinquiryextrachargecancelAPIRequest 送货入户并安装服务商取消增加费用 API请求
// alibaba.ascp.industry.inquiry.extracharge.cancel
//
// 送货入户并安装服务商取消增加费用
type AlibabaascpindustryinquiryextrachargecancelAPIRequest struct {
	model.Params
	// 请求对象
	_omsCancelExtraChargeParameter *OmsCancelExtraChargeParameter
}

// NewAlibabaascpindustryinquiryextrachargecancelRequest 初始化AlibabaascpindustryinquiryextrachargecancelAPIRequest对象
func NewAlibabaascpindustryinquiryextrachargecancelRequest() *AlibabaascpindustryinquiryextrachargecancelAPIRequest {
	return &AlibabaascpindustryinquiryextrachargecancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpindustryinquiryextrachargecancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.inquiry.extracharge.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpindustryinquiryextrachargecancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpindustryinquiryextrachargecancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOmsCancelExtraChargeParameter is OmsCancelExtraChargeParameter Setter
// 请求对象
func (r *AlibabaascpindustryinquiryextrachargecancelAPIRequest) SetOmsCancelExtraChargeParameter(_omsCancelExtraChargeParameter *OmsCancelExtraChargeParameter) error {
	r._omsCancelExtraChargeParameter = _omsCancelExtraChargeParameter
	r.Set("oms_cancel_extra_charge_parameter", _omsCancelExtraChargeParameter)
	return nil
}

// GetOmsCancelExtraChargeParameter OmsCancelExtraChargeParameter Getter
func (r AlibabaascpindustryinquiryextrachargecancelAPIRequest) GetOmsCancelExtraChargeParameter() *OmsCancelExtraChargeParameter {
	return r._omsCancelExtraChargeParameter
}
