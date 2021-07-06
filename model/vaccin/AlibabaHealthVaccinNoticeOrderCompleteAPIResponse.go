package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeOrderCompleteAPIResponse 疫苗接种完成(带支付宝提醒) API返回值
// alibaba.health.vaccin.notice.order.complete
//
// 用户到店完成接种,ISV感知通知阿里健康完成接种,并通知用户!
type AlibabaHealthVaccinNoticeOrderCompleteAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinNoticeOrderCompleteAPIResponseModel
}

// AlibabaHealthVaccinNoticeOrderCompleteAPIResponseModel is 疫苗接种完成(带支付宝提醒) 成功返回结果
type AlibabaHealthVaccinNoticeOrderCompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_notice_order_complete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 执行是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
