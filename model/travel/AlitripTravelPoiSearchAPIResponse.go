package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelPoiSearchAPIResponse POI信息查询 API返回值
// alitrip.travel.poi.search
//
// POI信息查询，用于商品更新使用
type AlitripTravelPoiSearchAPIResponse struct {
	model.CommonResponse
	AlitripTravelPoiSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelPoiSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelPoiSearchAPIResponseModel).Reset()
}

// AlitripTravelPoiSearchAPIResponseModel is POI信息查询 成功返回结果
type AlitripTravelPoiSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_poi_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// POI详情
	Results []Poi `json:"results,omitempty" xml:"results>poi,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelPoiSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolAlitripTravelPoiSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelPoiSearchAPIResponse)
	},
}

// GetAlitripTravelPoiSearchAPIResponse 从 sync.Pool 获取 AlitripTravelPoiSearchAPIResponse
func GetAlitripTravelPoiSearchAPIResponse() *AlitripTravelPoiSearchAPIResponse {
	return poolAlitripTravelPoiSearchAPIResponse.Get().(*AlitripTravelPoiSearchAPIResponse)
}

// ReleaseAlitripTravelPoiSearchAPIResponse 将 AlitripTravelPoiSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelPoiSearchAPIResponse(v *AlitripTravelPoiSearchAPIResponse) {
	v.Reset()
	poolAlitripTravelPoiSearchAPIResponse.Put(v)
}
