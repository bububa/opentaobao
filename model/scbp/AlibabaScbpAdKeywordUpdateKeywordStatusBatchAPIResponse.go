package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIResponse
修改关键词状态 API返回值
alibaba.scbp.ad.keyword.update.keyword.status.batch

修改关键词状态 */
type AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIResponseModel
}

// AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIResponseModel is 修改关键词状态 成功返回结果
type AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_update_keyword_status_batch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回错误集合
	ResultList []ErrorKeyword `json:"result_list,omitempty" xml:"result_list>error_keyword,omitempty"`
}
