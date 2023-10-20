package ascp

import (
	"encoding/xml"

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

// AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponseModel is 送货入户并安装服务商取消增加费用 成功返回结果
type AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_inquiry_extracharge_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
