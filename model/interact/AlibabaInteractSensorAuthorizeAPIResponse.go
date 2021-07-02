package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorAuthorizeAPIResponse 客户端授权页 API返回值
// alibaba.interact.sensor.authorize
//
// 客户端授权页
type AlibabaInteractSensorAuthorizeAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorAuthorizeAPIResponseModel
}

// AlibabaInteractSensorAuthorizeAPIResponseModel is 客户端授权页 成功返回结果
type AlibabaInteractSensorAuthorizeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_authorize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
