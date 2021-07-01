package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量启动暂停推广词状态 API返回值 
alibaba.scbp.ad.keyword.status.batchupdate

批量启动暂停关键词推广状态
*/
type AlibabaScbpAdKeywordStatusBatchupdateAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordStatusBatchupdateAPIResponseModel
}

// 批量启动暂停推广词状态 成功返回结果
type AlibabaScbpAdKeywordStatusBatchupdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_status_batchupdate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 修改失败关键词列表
    KeywordErrorResultList   []KeywordErrorResultDto `json:"keyword_error_result_list,omitempty" xml:"keyword_error_result_list>keyword_error_result_dto,omitempty"`
}
