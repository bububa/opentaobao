package icburfq

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuquotationpostAPIResponse 供应商提交报价 API返回值
// alibaba.icbu.quotation.post
//
// 供应商对RFQ进行报价
type AlibabaicbuquotationpostAPIResponse struct {
	model.CommonResponse
	AlibabaicbuquotationpostAPIResponseModel
}

// AlibabaicbuquotationpostAPIResponseModel is 供应商提交报价 成功返回结果
type AlibabaicbuquotationpostAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_quotation_post_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求返回结果信息
	Result *RfqRemoteServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
