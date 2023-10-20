package taotv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuTvDesktopToyouRecommendAPIResponse TV桌面为你推荐接口 API返回值
// youku.tv.desktop.toyou.recommend
//
// 提供为你推荐数据
type YoukuTvDesktopToyouRecommendAPIResponse struct {
	model.CommonResponse
	YoukuTvDesktopToyouRecommendAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuTvDesktopToyouRecommendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuTvDesktopToyouRecommendAPIResponseModel).Reset()
}

// YoukuTvDesktopToyouRecommendAPIResponseModel is TV桌面为你推荐接口 成功返回结果
type YoukuTvDesktopToyouRecommendAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_tv_desktop_toyou_recommend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应的结果列表
	Results []V5BaseItemRbo `json:"results,omitempty" xml:"results>v5base_item_rbo,omitempty"`
}

// Reset 清空结构体
func (m *YoukuTvDesktopToyouRecommendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolYoukuTvDesktopToyouRecommendAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuTvDesktopToyouRecommendAPIResponse)
	},
}

// GetYoukuTvDesktopToyouRecommendAPIResponse 从 sync.Pool 获取 YoukuTvDesktopToyouRecommendAPIResponse
func GetYoukuTvDesktopToyouRecommendAPIResponse() *YoukuTvDesktopToyouRecommendAPIResponse {
	return poolYoukuTvDesktopToyouRecommendAPIResponse.Get().(*YoukuTvDesktopToyouRecommendAPIResponse)
}

// ReleaseYoukuTvDesktopToyouRecommendAPIResponse 将 YoukuTvDesktopToyouRecommendAPIResponse 保存到 sync.Pool
func ReleaseYoukuTvDesktopToyouRecommendAPIResponse(v *YoukuTvDesktopToyouRecommendAPIResponse) {
	v.Reset()
	poolYoukuTvDesktopToyouRecommendAPIResponse.Put(v)
}
