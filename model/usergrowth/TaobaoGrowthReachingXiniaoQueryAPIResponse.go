package usergrowth

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGrowthReachingXiniaoQueryAPIResponse 查询溪鸟推荐信息数据 API返回值
// taobao.growth.reaching.xiniao.query
//
// 查询溪鸟推荐信息数据
type TaobaoGrowthReachingXiniaoQueryAPIResponse struct {
	model.CommonResponse
	TaobaoGrowthReachingXiniaoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoGrowthReachingXiniaoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoGrowthReachingXiniaoQueryAPIResponseModel).Reset()
}

// TaobaoGrowthReachingXiniaoQueryAPIResponseModel is 查询溪鸟推荐信息数据 成功返回结果
type TaobaoGrowthReachingXiniaoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"growth_reaching_xiniao_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推荐信息列表
	Suggestions []XiNiaoSuggestionDto `json:"suggestions,omitempty" xml:"suggestions>xi_niao_suggestion_dto,omitempty"`
	// 用户类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 是否展现
	IsOffering bool `json:"is_offering,omitempty" xml:"is_offering,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoGrowthReachingXiniaoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Suggestions = m.Suggestions[:0]
	m.Type = ""
	m.IsOffering = false
}

var poolTaobaoGrowthReachingXiniaoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoGrowthReachingXiniaoQueryAPIResponse)
	},
}

// GetTaobaoGrowthReachingXiniaoQueryAPIResponse 从 sync.Pool 获取 TaobaoGrowthReachingXiniaoQueryAPIResponse
func GetTaobaoGrowthReachingXiniaoQueryAPIResponse() *TaobaoGrowthReachingXiniaoQueryAPIResponse {
	return poolTaobaoGrowthReachingXiniaoQueryAPIResponse.Get().(*TaobaoGrowthReachingXiniaoQueryAPIResponse)
}

// ReleaseTaobaoGrowthReachingXiniaoQueryAPIResponse 将 TaobaoGrowthReachingXiniaoQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoGrowthReachingXiniaoQueryAPIResponse(v *TaobaoGrowthReachingXiniaoQueryAPIResponse) {
	v.Reset()
	poolTaobaoGrowthReachingXiniaoQueryAPIResponse.Put(v)
}
