package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightCatsinfoGetAPIResponse 类目信息获取 API返回值
// taobao.simba.insight.catsinfo.get
//
// 获取类目信息，此接口既提供所有顶级类目的查询，又提供给定类目id自身信息和子类目信息的查询，所以可以根据此接口逐层获取所有的类目信息
type TaobaoSimbaInsightCatsinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaInsightCatsinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaInsightCatsinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaInsightCatsinfoGetAPIResponseModel).Reset()
}

// TaobaoSimbaInsightCatsinfoGetAPIResponseModel is 类目信息获取 成功返回结果
type TaobaoSimbaInsightCatsinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_insight_catsinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目详细信息
	CategoryInfoList []InsightCategoryInfoDto `json:"category_info_list,omitempty" xml:"category_info_list>insight_category_info_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaInsightCatsinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CategoryInfoList = m.CategoryInfoList[:0]
}

var poolTaobaoSimbaInsightCatsinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaInsightCatsinfoGetAPIResponse)
	},
}

// GetTaobaoSimbaInsightCatsinfoGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaInsightCatsinfoGetAPIResponse
func GetTaobaoSimbaInsightCatsinfoGetAPIResponse() *TaobaoSimbaInsightCatsinfoGetAPIResponse {
	return poolTaobaoSimbaInsightCatsinfoGetAPIResponse.Get().(*TaobaoSimbaInsightCatsinfoGetAPIResponse)
}

// ReleaseTaobaoSimbaInsightCatsinfoGetAPIResponse 将 TaobaoSimbaInsightCatsinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaInsightCatsinfoGetAPIResponse(v *TaobaoSimbaInsightCatsinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaInsightCatsinfoGetAPIResponse.Put(v)
}
