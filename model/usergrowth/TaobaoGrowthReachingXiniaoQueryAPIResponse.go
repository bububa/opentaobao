package usergrowth

import (
	"encoding/xml"

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
