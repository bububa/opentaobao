package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoAutofinanceCreditReceiveAPIResponse 接收授信结果通知 API返回值
// tmall.aliauto.autofinance.credit.receive
//
// 天猫汽车的金融业务场景中，需要接收外部ISV对用户授信申请的通知结果.
type TmallAliautoAutofinanceCreditReceiveAPIResponse struct {
	model.CommonResponse
	TmallAliautoAutofinanceCreditReceiveAPIResponseModel
}

// TmallAliautoAutofinanceCreditReceiveAPIResponseModel is 接收授信结果通知 成功返回结果
type TmallAliautoAutofinanceCreditReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_autofinance_credit_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务结果说明
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息描述
	ErorMessage string `json:"eror_message,omitempty" xml:"eror_message,omitempty"`
	// 错误码
	ErorCode string `json:"eror_code,omitempty" xml:"eror_code,omitempty"`
	// 是否成功
	Succes bool `json:"succes,omitempty" xml:"succes,omitempty"`
}
