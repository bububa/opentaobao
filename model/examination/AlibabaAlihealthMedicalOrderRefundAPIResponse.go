package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalOrderRefundAPIResponse 退款接口 API返回值
// alibaba.alihealth.medical.order.refund
//
// 退款接口
type AlibabaAlihealthMedicalOrderRefundAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalOrderRefundAPIResponseModel
}

// AlibabaAlihealthMedicalOrderRefundAPIResponseModel is 退款接口 成功返回结果
type AlibabaAlihealthMedicalOrderRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_order_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihealthMedicalOrderRefundResult `json:"result,omitempty" xml:"result,omitempty"`
}
