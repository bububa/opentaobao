package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoAutofinanceLoanReceiveAPIResponse 接收放款结果通知 API返回值
// tmall.aliauto.autofinance.loan.receive
//
// 天猫汽车的金融业务场景中，需要接收外部ISV对用户支用放款的通知结果
type TmallAliautoAutofinanceLoanReceiveAPIResponse struct {
	model.CommonResponse
	TmallAliautoAutofinanceLoanReceiveAPIResponseModel
}

// TmallAliautoAutofinanceLoanReceiveAPIResponseModel is 接收放款结果通知 成功返回结果
type TmallAliautoAutofinanceLoanReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_autofinance_loan_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErorCode string `json:"eror_code,omitempty" xml:"eror_code,omitempty"`
	// 错误信息描述
	ErorMessage string `json:"eror_message,omitempty" xml:"eror_message,omitempty"`
	// 业务结果说明
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Succes bool `json:"succes,omitempty" xml:"succes,omitempty"`
}
