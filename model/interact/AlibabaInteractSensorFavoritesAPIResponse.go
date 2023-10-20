package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorFavoritesAPIResponse 手淘开放收藏夹鉴权接口 API返回值
// alibaba.interact.sensor.favorites
//
// 手淘开放鉴权专用接口，无数据输出输入，仅用于鉴权。
type AlibabaInteractSensorFavoritesAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorFavoritesAPIResponseModel
}

// AlibabaInteractSensorFavoritesAPIResponseModel is 手淘开放收藏夹鉴权接口 成功返回结果
type AlibabaInteractSensorFavoritesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_favorites_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// r=0
	R string `json:"r,omitempty" xml:"r,omitempty"`
}
