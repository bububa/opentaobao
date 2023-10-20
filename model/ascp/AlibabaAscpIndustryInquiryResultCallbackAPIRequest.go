package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryInquiryResultCallbackAPIRequest 送货入户并安装服务商询价结果返回 API请求
// alibaba.ascp.industry.inquiry.result.callback
//
// 送货入户并安装服务商询价结果返回
type AlibabaAscpIndustryInquiryResultCallbackAPIRequest struct {
	model.Params
	// 询价结果
	_inquiryResult *InquiryResult
}

// NewAlibabaAscpIndustryInquiryResultCallbackRequest 初始化AlibabaAscpIndustryInquiryResultCallbackAPIRequest对象
func NewAlibabaAscpIndustryInquiryResultCallbackRequest() *AlibabaAscpIndustryInquiryResultCallbackAPIRequest {
	return &AlibabaAscpIndustryInquiryResultCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpIndustryInquiryResultCallbackAPIRequest) Reset() {
	r._inquiryResult = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryInquiryResultCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.inquiry.result.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryInquiryResultCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpIndustryInquiryResultCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInquiryResult is InquiryResult Setter
// 询价结果
func (r *AlibabaAscpIndustryInquiryResultCallbackAPIRequest) SetInquiryResult(_inquiryResult *InquiryResult) error {
	r._inquiryResult = _inquiryResult
	r.Set("inquiry_result", _inquiryResult)
	return nil
}

// GetInquiryResult InquiryResult Getter
func (r AlibabaAscpIndustryInquiryResultCallbackAPIRequest) GetInquiryResult() *InquiryResult {
	return r._inquiryResult
}

var poolAlibabaAscpIndustryInquiryResultCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpIndustryInquiryResultCallbackRequest()
	},
}

// GetAlibabaAscpIndustryInquiryResultCallbackRequest 从 sync.Pool 获取 AlibabaAscpIndustryInquiryResultCallbackAPIRequest
func GetAlibabaAscpIndustryInquiryResultCallbackAPIRequest() *AlibabaAscpIndustryInquiryResultCallbackAPIRequest {
	return poolAlibabaAscpIndustryInquiryResultCallbackAPIRequest.Get().(*AlibabaAscpIndustryInquiryResultCallbackAPIRequest)
}

// ReleaseAlibabaAscpIndustryInquiryResultCallbackAPIRequest 将 AlibabaAscpIndustryInquiryResultCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpIndustryInquiryResultCallbackAPIRequest(v *AlibabaAscpIndustryInquiryResultCallbackAPIRequest) {
	v.Reset()
	poolAlibabaAscpIndustryInquiryResultCallbackAPIRequest.Put(v)
}
