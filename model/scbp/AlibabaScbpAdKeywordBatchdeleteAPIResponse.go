package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordBatchdeleteAPIResponse 外贸直通车批量删除关键词 API返回值
// alibaba.scbp.ad.keyword.batchdelete
//
// 外贸直通车批量删除关键词
type AlibabaScbpAdKeywordBatchdeleteAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordBatchdeleteAPIResponseModel
}

// AlibabaScbpAdKeywordBatchdeleteAPIResponseModel is 外贸直通车批量删除关键词 成功返回结果
type AlibabaScbpAdKeywordBatchdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_batchdelete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除关键词是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
