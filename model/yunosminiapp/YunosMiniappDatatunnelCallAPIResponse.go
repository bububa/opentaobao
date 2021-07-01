package yunosminiapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosMiniappDatatunnelCallAPIResponse
车载小程序外部服务调用 API返回值
yunos.miniapp.datatunnel.call

对客户提供的api进行统一封装调用。 */
type YunosMiniappDatatunnelCallAPIResponse struct {
	model.CommonResponse
	YunosMiniappDatatunnelCallAPIResponseModel
}

// YunosMiniappDatatunnelCallAPIResponseModel is 车载小程序外部服务调用 成功返回结果
type YunosMiniappDatatunnelCallAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_miniapp_datatunnel_call_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *YunosMiniappDatatunnelCallMapResult `json:"result,omitempty" xml:"result,omitempty"`
}
