package usergrowth

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaogrowthreachingsuggestionqueryAPIResponse 厂商生态推荐信息查询 API返回值
// taobao.growth.reaching.suggestion.query
//
// 厂商生态推荐信息查询
type TaobaogrowthreachingsuggestionqueryAPIResponse struct {
	model.CommonResponse
	TaobaogrowthreachingsuggestionqueryAPIResponseModel
}

// TaobaogrowthreachingsuggestionqueryAPIResponseModel is 厂商生态推荐信息查询 成功返回结果
type TaobaogrowthreachingsuggestionqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"growth_reaching_suggestion_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推荐信息列表
	Suggestions []SuggestionDto `json:"suggestions,omitempty" xml:"suggestions>suggestion_dto,omitempty"`
	// 列表维度曝光上报链接
	ExposureUrl string `json:"exposure_url,omitempty" xml:"exposure_url,omitempty"`
	// 是否参竞
	IsOffering bool `json:"is_offering,omitempty" xml:"is_offering,omitempty"`
}
