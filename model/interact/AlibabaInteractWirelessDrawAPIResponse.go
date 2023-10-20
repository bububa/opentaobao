package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractwirelessdrawAPIResponse 双11到店互动无线端抽奖接口鉴权 API返回值
// alibaba.interact.wireless.draw
//
// 双11到店互动无线端mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 坯子
type AlibabainteractwirelessdrawAPIResponse struct {
	model.CommonResponse
	AlibabainteractwirelessdrawAPIResponseModel
}

// AlibabainteractwirelessdrawAPIResponseModel is 双11到店互动无线端抽奖接口鉴权 成功返回结果
type AlibabainteractwirelessdrawAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_wireless_draw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
