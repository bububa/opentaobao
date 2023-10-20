package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleOnespuRegisterUpdateAPIResponse 闲鱼 ONESPU 挂载接口 API返回值
// alibaba.idle.onespu.register.update
//
// 闲鱼 ONESPU 挂载接口
type AlibabaIdleOnespuRegisterUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaIdleOnespuRegisterUpdateAPIResponseModel
}

// AlibabaIdleOnespuRegisterUpdateAPIResponseModel is 闲鱼 ONESPU 挂载接口 成功返回结果
type AlibabaIdleOnespuRegisterUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_onespu_register_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaIdleOnespuRegisterUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
