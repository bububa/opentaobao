package alitrippoi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformPoiRawPoioutAPIResponse 飞猪poi输出 API返回值
// alitrip.platform.poi.raw.poiout
//
// 输出指定城市poi指定信息
type AlitripPlatformPoiRawPoioutAPIResponse struct {
	model.CommonResponse
	AlitripPlatformPoiRawPoioutAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPlatformPoiRawPoioutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPlatformPoiRawPoioutAPIResponseModel).Reset()
}

// AlitripPlatformPoiRawPoioutAPIResponseModel is 飞猪poi输出 成功返回结果
type AlitripPlatformPoiRawPoioutAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_platform_poi_raw_poiout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripPlatformPoiRawPoioutResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPlatformPoiRawPoioutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPlatformPoiRawPoioutAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPlatformPoiRawPoioutAPIResponse)
	},
}

// GetAlitripPlatformPoiRawPoioutAPIResponse 从 sync.Pool 获取 AlitripPlatformPoiRawPoioutAPIResponse
func GetAlitripPlatformPoiRawPoioutAPIResponse() *AlitripPlatformPoiRawPoioutAPIResponse {
	return poolAlitripPlatformPoiRawPoioutAPIResponse.Get().(*AlitripPlatformPoiRawPoioutAPIResponse)
}

// ReleaseAlitripPlatformPoiRawPoioutAPIResponse 将 AlitripPlatformPoiRawPoioutAPIResponse 保存到 sync.Pool
func ReleaseAlitripPlatformPoiRawPoioutAPIResponse(v *AlitripPlatformPoiRawPoioutAPIResponse) {
	v.Reset()
	poolAlitripPlatformPoiRawPoioutAPIResponse.Put(v)
}
