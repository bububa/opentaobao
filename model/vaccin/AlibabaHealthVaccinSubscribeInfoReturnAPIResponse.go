package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinsubscribeinforeturnAPIResponse 自有pov预约信息回传 API返回值
// alibaba.health.vaccin.subscribe.info.return
//
// 自有pov预约信息回传
type AlibabahealthvaccinsubscribeinforeturnAPIResponse struct {
	model.CommonResponse
	AlibabahealthvaccinsubscribeinforeturnAPIResponseModel
}

// AlibabahealthvaccinsubscribeinforeturnAPIResponseModel is 自有pov预约信息回传 成功返回结果
type AlibabahealthvaccinsubscribeinforeturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_subscribe_info_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// info
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *IsvPovSubscribeInfoResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
