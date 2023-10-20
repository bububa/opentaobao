package alitrippoi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformPoiRawSaverawpoiAPIResponse POI开放存储能力 API返回值
// alitrip.platform.poi.raw.saverawpoi
//
// POI开放存储提供离线/在线/纬错更新的能力
type AlitripPlatformPoiRawSaverawpoiAPIResponse struct {
	model.CommonResponse
	AlitripPlatformPoiRawSaverawpoiAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPlatformPoiRawSaverawpoiAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPlatformPoiRawSaverawpoiAPIResponseModel).Reset()
}

// AlitripPlatformPoiRawSaverawpoiAPIResponseModel is POI开放存储能力 成功返回结果
type AlitripPlatformPoiRawSaverawpoiAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_platform_poi_raw_saverawpoi_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripPlatformPoiRawSaverawpoiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPlatformPoiRawSaverawpoiAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPlatformPoiRawSaverawpoiAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPlatformPoiRawSaverawpoiAPIResponse)
	},
}

// GetAlitripPlatformPoiRawSaverawpoiAPIResponse 从 sync.Pool 获取 AlitripPlatformPoiRawSaverawpoiAPIResponse
func GetAlitripPlatformPoiRawSaverawpoiAPIResponse() *AlitripPlatformPoiRawSaverawpoiAPIResponse {
	return poolAlitripPlatformPoiRawSaverawpoiAPIResponse.Get().(*AlitripPlatformPoiRawSaverawpoiAPIResponse)
}

// ReleaseAlitripPlatformPoiRawSaverawpoiAPIResponse 将 AlitripPlatformPoiRawSaverawpoiAPIResponse 保存到 sync.Pool
func ReleaseAlitripPlatformPoiRawSaverawpoiAPIResponse(v *AlitripPlatformPoiRawSaverawpoiAPIResponse) {
	v.Reset()
	poolAlitripPlatformPoiRawSaverawpoiAPIResponse.Put(v)
}
