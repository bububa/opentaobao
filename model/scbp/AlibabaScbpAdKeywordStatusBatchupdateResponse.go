package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量启动暂停推广词状态 APIResponse
alibaba.scbp.ad.keyword.status.batchupdate

批量启动暂停关键词推广状态
*/
type AlibabaScbpAdKeywordStatusBatchupdateAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordStatusBatchupdateResponse
}

type AlibabaScbpAdKeywordStatusBatchupdateResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_status_batchupdate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改失败关键词列表
    
    KeywordErrorResultList   []KeywordErrorResultDto `json:"keyword_error_result_list,omitempty" xml:"keyword_error_result_list>keyword_error_result_dto,omitempty"`
    
    
}
