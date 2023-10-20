package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordbatchquerykeywordrankpriceAPIResponse 批量查询关键词前五名排价 API返回值
// alibaba.scbp.ad.keyword.batch.query.keyword.rank.price
//
// 批量查询关键词前五名排价
type AlibabascbpadkeywordbatchquerykeywordrankpriceAPIResponse struct {
	model.CommonResponse
	AlibabascbpadkeywordbatchquerykeywordrankpriceAPIResponseModel
}

// AlibabascbpadkeywordbatchquerykeywordrankpriceAPIResponseModel is 批量查询关键词前五名排价 成功返回结果
type AlibabascbpadkeywordbatchquerykeywordrankpriceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_batch_query_keyword_rank_price_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词前五名排价详细信息返回
	ResultList []KeywordRankPriceDto `json:"result_list,omitempty" xml:"result_list>keyword_rank_price_dto,omitempty"`
}
