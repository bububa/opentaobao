package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorMakeupAPIResponse 美妆虚拟试装 API返回值
// alibaba.interact.sensor.makeup
//
// 手机淘宝美妆类目虚拟试妆权限，客户端能力（JS－API）
type AlibabaInteractSensorMakeupAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorMakeupAPIResponseModel
}

// AlibabaInteractSensorMakeupAPIResponseModel is 美妆虚拟试装 成功返回结果
type AlibabaInteractSensorMakeupAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_makeup_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 非restAPI，为jsapi  result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
