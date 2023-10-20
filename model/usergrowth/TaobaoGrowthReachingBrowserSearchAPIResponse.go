package usergrowth

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGrowthReachingBrowserSearchAPIResponse 查询搜索关联 API返回值
// taobao.growth.reaching.browser.search
//
// 查询搜索关联
type TaobaoGrowthReachingBrowserSearchAPIResponse struct {
	model.CommonResponse
	TaobaoGrowthReachingBrowserSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoGrowthReachingBrowserSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoGrowthReachingBrowserSearchAPIResponseModel).Reset()
}

// TaobaoGrowthReachingBrowserSearchAPIResponseModel is 查询搜索关联 成功返回结果
type TaobaoGrowthReachingBrowserSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"growth_reaching_browser_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// []
	Products []ProductDto `json:"products,omitempty" xml:"products>product_dto,omitempty"`
	// 榜单曝光链接
	ExposureUrl string `json:"exposure_url,omitempty" xml:"exposure_url,omitempty"`
	// 是否参竞
	Offering bool `json:"offering,omitempty" xml:"offering,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoGrowthReachingBrowserSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Products = m.Products[:0]
	m.ExposureUrl = ""
	m.Offering = false
}

var poolTaobaoGrowthReachingBrowserSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoGrowthReachingBrowserSearchAPIResponse)
	},
}

// GetTaobaoGrowthReachingBrowserSearchAPIResponse 从 sync.Pool 获取 TaobaoGrowthReachingBrowserSearchAPIResponse
func GetTaobaoGrowthReachingBrowserSearchAPIResponse() *TaobaoGrowthReachingBrowserSearchAPIResponse {
	return poolTaobaoGrowthReachingBrowserSearchAPIResponse.Get().(*TaobaoGrowthReachingBrowserSearchAPIResponse)
}

// ReleaseTaobaoGrowthReachingBrowserSearchAPIResponse 将 TaobaoGrowthReachingBrowserSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoGrowthReachingBrowserSearchAPIResponse(v *TaobaoGrowthReachingBrowserSearchAPIResponse) {
	v.Reset()
	poolTaobaoGrowthReachingBrowserSearchAPIResponse.Put(v)
}
