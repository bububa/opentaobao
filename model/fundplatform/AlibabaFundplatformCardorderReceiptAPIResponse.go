package fundplatform

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardorderReceiptAPIResponse
通知确认收货 API返回值
alibaba.fundplatform.cardorder.receipt

告知卡商这一批储值卡已经被用户确认收货 */
type AlibabaFundplatformCardorderReceiptAPIResponse struct {
	model.CommonResponse
	AlibabaFundplatformCardorderReceiptAPIResponseModel
}

// AlibabaFundplatformCardorderReceiptAPIResponseModel is 通知确认收货 成功返回结果
type AlibabaFundplatformCardorderReceiptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_cardorder_receipt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误详情
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 错误CODE
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
