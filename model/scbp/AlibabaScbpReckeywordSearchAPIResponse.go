package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpreckeywordsearchAPIResponse 推荐词-词推词 API返回值
// alibaba.scbp.reckeyword.search
//
// 推荐词-词推词
type AlibabascbpreckeywordsearchAPIResponse struct {
	model.CommonResponse
	AlibabascbpreckeywordsearchAPIResponseModel
}

// AlibabascbpreckeywordsearchAPIResponseModel is 推荐词-词推词 成功返回结果
type AlibabascbpreckeywordsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_reckeyword_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 词推词结果列表
	ResultList []RecKeywordDto `json:"result_list,omitempty" xml:"result_list>rec_keyword_dto,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}
