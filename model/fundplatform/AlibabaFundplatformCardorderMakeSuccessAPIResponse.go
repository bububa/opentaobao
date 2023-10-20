package fundplatform

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformcardordermakesuccessAPIResponse 通知制卡成功 API返回值
// alibaba.fundplatform.cardorder.make.success
//
// 当外部业务方调用资金平台异步制卡接口后，资金平台制卡成功后通知异步通知业务方
type AlibabafundplatformcardordermakesuccessAPIResponse struct {
	model.CommonResponse
	AlibabafundplatformcardordermakesuccessAPIResponseModel
}

// AlibabafundplatformcardordermakesuccessAPIResponseModel is 通知制卡成功 成功返回结果
type AlibabafundplatformcardordermakesuccessAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_cardorder_make_success_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误详情
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 错误CODE
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
