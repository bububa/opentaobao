package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryInquiryResultCallbackAPIResponse 送货入户并安装服务商询价结果返回 API返回值
// alibaba.ascp.industry.inquiry.result.callback
//
// 送货入户并安装服务商询价结果返回
type AlibabaAscpIndustryInquiryResultCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaAscpIndustryInquiryResultCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryInquiryResultCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpIndustryInquiryResultCallbackAPIResponseModel).Reset()
}

// AlibabaAscpIndustryInquiryResultCallbackAPIResponseModel is 送货入户并安装服务商询价结果返回 成功返回结果
type AlibabaAscpIndustryInquiryResultCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_inquiry_result_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryInquiryResultCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpIndustryInquiryResultCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpIndustryInquiryResultCallbackAPIResponse)
	},
}

// GetAlibabaAscpIndustryInquiryResultCallbackAPIResponse 从 sync.Pool 获取 AlibabaAscpIndustryInquiryResultCallbackAPIResponse
func GetAlibabaAscpIndustryInquiryResultCallbackAPIResponse() *AlibabaAscpIndustryInquiryResultCallbackAPIResponse {
	return poolAlibabaAscpIndustryInquiryResultCallbackAPIResponse.Get().(*AlibabaAscpIndustryInquiryResultCallbackAPIResponse)
}

// ReleaseAlibabaAscpIndustryInquiryResultCallbackAPIResponse 将 AlibabaAscpIndustryInquiryResultCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpIndustryInquiryResultCallbackAPIResponse(v *AlibabaAscpIndustryInquiryResultCallbackAPIResponse) {
	v.Reset()
	poolAlibabaAscpIndustryInquiryResultCallbackAPIResponse.Put(v)
}
