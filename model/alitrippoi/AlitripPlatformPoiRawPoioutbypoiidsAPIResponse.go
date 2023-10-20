package alitrippoi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformPoiRawPoioutbypoiidsAPIResponse 根据poiId输出飞猪poi数据 API返回值
// alitrip.platform.poi.raw.poioutbypoiids
//
// 根据poiId输出飞猪poi数据
type AlitripPlatformPoiRawPoioutbypoiidsAPIResponse struct {
	model.CommonResponse
	AlitripPlatformPoiRawPoioutbypoiidsAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPlatformPoiRawPoioutbypoiidsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPlatformPoiRawPoioutbypoiidsAPIResponseModel).Reset()
}

// AlitripPlatformPoiRawPoioutbypoiidsAPIResponseModel is 根据poiId输出飞猪poi数据 成功返回结果
type AlitripPlatformPoiRawPoioutbypoiidsAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_platform_poi_raw_poioutbypoiids_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripPlatformPoiRawPoioutbypoiidsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPlatformPoiRawPoioutbypoiidsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPlatformPoiRawPoioutbypoiidsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPlatformPoiRawPoioutbypoiidsAPIResponse)
	},
}

// GetAlitripPlatformPoiRawPoioutbypoiidsAPIResponse 从 sync.Pool 获取 AlitripPlatformPoiRawPoioutbypoiidsAPIResponse
func GetAlitripPlatformPoiRawPoioutbypoiidsAPIResponse() *AlitripPlatformPoiRawPoioutbypoiidsAPIResponse {
	return poolAlitripPlatformPoiRawPoioutbypoiidsAPIResponse.Get().(*AlitripPlatformPoiRawPoioutbypoiidsAPIResponse)
}

// ReleaseAlitripPlatformPoiRawPoioutbypoiidsAPIResponse 将 AlitripPlatformPoiRawPoioutbypoiidsAPIResponse 保存到 sync.Pool
func ReleaseAlitripPlatformPoiRawPoioutbypoiidsAPIResponse(v *AlitripPlatformPoiRawPoioutbypoiidsAPIResponse) {
	v.Reset()
	poolAlitripPlatformPoiRawPoioutbypoiidsAPIResponse.Put(v)
}
