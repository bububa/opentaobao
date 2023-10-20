package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordrecommendwordAPIResponse 推词 API返回值
// alibaba.scbp.ad.keyword.recommend.word
//
// 推词
type AlibabascbpadkeywordrecommendwordAPIResponse struct {
	model.CommonResponse
	AlibabascbpadkeywordrecommendwordAPIResponseModel
}

// AlibabascbpadkeywordrecommendwordAPIResponseModel is 推词 成功返回结果
type AlibabascbpadkeywordrecommendwordAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_recommend_word_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推词返回实体类
	ResultList []AdRecommendWordDto `json:"result_list,omitempty" xml:"result_list>ad_recommend_word_dto,omitempty"`
}
