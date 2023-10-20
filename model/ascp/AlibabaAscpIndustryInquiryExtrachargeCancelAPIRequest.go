package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest 送货入户并安装服务商取消增加费用 API请求
// alibaba.ascp.industry.inquiry.extracharge.cancel
//
// 送货入户并安装服务商取消增加费用
type AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest struct {
	model.Params
	// 请求对象
	_omsCancelExtraChargeParameter *OmsCancelExtraChargeParameter
}

// NewAlibabaAscpIndustryInquiryExtrachargeCancelRequest 初始化AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest对象
func NewAlibabaAscpIndustryInquiryExtrachargeCancelRequest() *AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest {
	return &AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest) Reset() {
	r._omsCancelExtraChargeParameter = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.inquiry.extracharge.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOmsCancelExtraChargeParameter is OmsCancelExtraChargeParameter Setter
// 请求对象
func (r *AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest) SetOmsCancelExtraChargeParameter(_omsCancelExtraChargeParameter *OmsCancelExtraChargeParameter) error {
	r._omsCancelExtraChargeParameter = _omsCancelExtraChargeParameter
	r.Set("oms_cancel_extra_charge_parameter", _omsCancelExtraChargeParameter)
	return nil
}

// GetOmsCancelExtraChargeParameter OmsCancelExtraChargeParameter Getter
func (r AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest) GetOmsCancelExtraChargeParameter() *OmsCancelExtraChargeParameter {
	return r._omsCancelExtraChargeParameter
}

var poolAlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpIndustryInquiryExtrachargeCancelRequest()
	},
}

// GetAlibabaAscpIndustryInquiryExtrachargeCancelRequest 从 sync.Pool 获取 AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest
func GetAlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest() *AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest {
	return poolAlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest.Get().(*AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest)
}

// ReleaseAlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest 将 AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest(v *AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest) {
	v.Reset()
	poolAlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest.Put(v)
}
