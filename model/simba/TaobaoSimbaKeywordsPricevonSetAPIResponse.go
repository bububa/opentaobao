package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsPricevonSetAPIResponse 设置一批关键词的信息 API返回值
// taobao.simba.keywords.pricevon.set
//
// 设置一批关键词的信息，包含无线出价、计算机出价和关键词匹配方式
type TaobaoSimbaKeywordsPricevonSetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordsPricevonSetAPIResponseModel
}

// TaobaoSimbaKeywordsPricevonSetAPIResponseModel is 设置一批关键词的信息 成功返回结果
type TaobaoSimbaKeywordsPricevonSetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywords_pricevon_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功设置关键词价格的关键词列表
	Keywords []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
}
