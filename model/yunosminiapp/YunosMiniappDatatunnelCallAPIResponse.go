package yunosminiapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosminiappdatatunnelcallAPIResponse 车载小程序外部服务调用 API返回值
// yunos.miniapp.datatunnel.call
//
// 对客户提供的api进行统一封装调用。
type YunosminiappdatatunnelcallAPIResponse struct {
	model.CommonResponse
	YunosminiappdatatunnelcallAPIResponseModel
}

// YunosminiappdatatunnelcallAPIResponseModel is 车载小程序外部服务调用 成功返回结果
type YunosminiappdatatunnelcallAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_miniapp_datatunnel_call_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *YunosminiappdatatunnelcallMapResult `json:"result,omitempty" xml:"result,omitempty"`
}
