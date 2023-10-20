package interact

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaInteractSensorFavoritesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorFavoritesAPIResponseModel).Reset()
}

// AlibabaInteractSensorFavoritesAPIResponseModel is 手淘开放收藏夹鉴权接口 成功返回结果
type AlibabaInteractSensorFavoritesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_favorites_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// r=0
	R string `json:"r,omitempty" xml:"r,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorFavoritesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.R = ""
}

var poolAlibabaInteractSensorFavoritesAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorFavoritesAPIResponse)
	},
}

// GetAlibabaInteractSensorFavoritesAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorFavoritesAPIResponse
func GetAlibabaInteractSensorFavoritesAPIResponse() *AlibabaInteractSensorFavoritesAPIResponse {
	return poolAlibabaInteractSensorFavoritesAPIResponse.Get().(*AlibabaInteractSensorFavoritesAPIResponse)
}

// ReleaseAlibabaInteractSensorFavoritesAPIResponse 将 AlibabaInteractSensorFavoritesAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorFavoritesAPIResponse(v *AlibabaInteractSensorFavoritesAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorFavoritesAPIResponse.Put(v)
}
