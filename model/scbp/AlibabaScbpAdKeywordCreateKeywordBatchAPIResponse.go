package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordCreateKeywordBatchAPIResponse 关键词添加 API返回值
// alibaba.scbp.ad.keyword.create.keyword.batch
//
// 关键词添加
type AlibabaScbpAdKeywordCreateKeywordBatchAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordCreateKeywordBatchAPIResponseModel
}

// AlibabaScbpAdKeywordCreateKeywordBatchAPIResponseModel is 关键词添加 成功返回结果
type AlibabaScbpAdKeywordCreateKeywordBatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_create_keyword_batch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回错误集合
	ResultList []ErrorKeyword `json:"result_list,omitempty" xml:"result_list>error_keyword,omitempty"`
}
