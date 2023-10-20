package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse 送货入户并安装服务商取消增加费用 API返回值
// alibaba.ascp.industry.inquiry.extracharge.cancel
//
// 送货入户并安装服务商取消增加费用
type AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse struct {
	model.CommonResponse
	AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponseModel).Reset()
}

// AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponseModel is 送货入户并安装服务商取消增加费用 成功返回结果
type AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_inquiry_extracharge_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse)
	},
}

// GetAlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse 从 sync.Pool 获取 AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse
func GetAlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse() *AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse {
	return poolAlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse.Get().(*AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse)
}

// ReleaseAlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse 将 AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse(v *AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse) {
	v.Reset()
	poolAlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse.Put(v)
}
