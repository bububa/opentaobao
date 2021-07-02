package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractIsvGatewayAPIResponse isv调用gateway API返回值
// alibaba.interact.isv.gateway
//
// isv能够调用jae本身的server
type AlibabaInteractIsvGatewayAPIResponse struct {
	model.CommonResponse
	AlibabaInteractIsvGatewayAPIResponseModel
}

// AlibabaInteractIsvGatewayAPIResponseModel is isv调用gateway 成功返回结果
type AlibabaInteractIsvGatewayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_isv_gateway_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ret=0
	Ret string `json:"ret,omitempty" xml:"ret,omitempty"`
}
