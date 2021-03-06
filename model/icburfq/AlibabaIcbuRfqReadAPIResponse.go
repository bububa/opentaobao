package icburfq

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqReadAPIResponse 是否已读RFQ API返回值
// alibaba.icbu.rfq.read
//
// 是否已读RFQ
type AlibabaIcbuRfqReadAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuRfqReadAPIResponseModel
}

// AlibabaIcbuRfqReadAPIResponseModel is 是否已读RFQ 成功返回结果
type AlibabaIcbuRfqReadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_rfq_read_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
