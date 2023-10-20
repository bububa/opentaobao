package alitrippoi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformPoiRawFeedAPIResponse 存储poi原始数据 API返回值
// alitrip.platform.poi.raw.feed
//
// 对接外部数据源，外部数据推送poi数据到飞猪
type AlitripPlatformPoiRawFeedAPIResponse struct {
	model.CommonResponse
	AlitripPlatformPoiRawFeedAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPlatformPoiRawFeedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPlatformPoiRawFeedAPIResponseModel).Reset()
}

// AlitripPlatformPoiRawFeedAPIResponseModel is 存储poi原始数据 成功返回结果
type AlitripPlatformPoiRawFeedAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_platform_poi_raw_feed_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlitripPlatformPoiRawFeedResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPlatformPoiRawFeedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPlatformPoiRawFeedAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPlatformPoiRawFeedAPIResponse)
	},
}

// GetAlitripPlatformPoiRawFeedAPIResponse 从 sync.Pool 获取 AlitripPlatformPoiRawFeedAPIResponse
func GetAlitripPlatformPoiRawFeedAPIResponse() *AlitripPlatformPoiRawFeedAPIResponse {
	return poolAlitripPlatformPoiRawFeedAPIResponse.Get().(*AlitripPlatformPoiRawFeedAPIResponse)
}

// ReleaseAlitripPlatformPoiRawFeedAPIResponse 将 AlitripPlatformPoiRawFeedAPIResponse 保存到 sync.Pool
func ReleaseAlitripPlatformPoiRawFeedAPIResponse(v *AlitripPlatformPoiRawFeedAPIResponse) {
	v.Reset()
	poolAlitripPlatformPoiRawFeedAPIResponse.Put(v)
}
