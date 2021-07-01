package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsQscoreSplitGetAPIResponse
新质量分服务 API返回值
taobao.simba.keywords.qscore.split.get

获取关键词新的质量分 */
type TaobaoSimbaKeywordsQscoreSplitGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordsQscoreSplitGetAPIResponseModel
}

// TaobaoSimbaKeywordsQscoreSplitGetAPIResponseModel is 新质量分服务 成功返回结果
type TaobaoSimbaKeywordsQscoreSplitGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywords_qscore_split_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoSimbaKeywordsQscoreSplitGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
