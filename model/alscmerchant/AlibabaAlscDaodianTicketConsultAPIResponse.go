package alscmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscdaodianticketconsultAPIResponse 券码预览 API返回值
// alibaba.alsc.daodian.ticket.consult
//
// 券码预览
type AlibabaalscdaodianticketconsultAPIResponse struct {
	model.CommonResponse
	AlibabaalscdaodianticketconsultAPIResponseModel
}

// AlibabaalscdaodianticketconsultAPIResponseModel is 券码预览 成功返回结果
type AlibabaalscdaodianticketconsultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_daodian_ticket_consult_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Result *AlibabaalscdaodianticketconsultResult `json:"result,omitempty" xml:"result,omitempty"`
}
