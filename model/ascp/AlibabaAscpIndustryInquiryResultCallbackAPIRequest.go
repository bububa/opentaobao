package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustryinquiryresultcallbackAPIRequest 送货入户并安装服务商询价结果返回 API请求
// alibaba.ascp.industry.inquiry.result.callback
//
// 送货入户并安装服务商询价结果返回
type AlibabaascpindustryinquiryresultcallbackAPIRequest struct {
	model.Params
	// 询价结果
	_inquiryResult *InquiryResult
}

// NewAlibabaascpindustryinquiryresultcallbackRequest 初始化AlibabaascpindustryinquiryresultcallbackAPIRequest对象
func NewAlibabaascpindustryinquiryresultcallbackRequest() *AlibabaascpindustryinquiryresultcallbackAPIRequest {
	return &AlibabaascpindustryinquiryresultcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpindustryinquiryresultcallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.inquiry.result.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpindustryinquiryresultcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpindustryinquiryresultcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInquiryResult is InquiryResult Setter
// 询价结果
func (r *AlibabaascpindustryinquiryresultcallbackAPIRequest) SetInquiryResult(_inquiryResult *InquiryResult) error {
	r._inquiryResult = _inquiryResult
	r.Set("inquiry_result", _inquiryResult)
	return nil
}

// GetInquiryResult InquiryResult Getter
func (r AlibabaascpindustryinquiryresultcallbackAPIRequest) GetInquiryResult() *InquiryResult {
	return r._inquiryResult
}
