package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
关键词启动暂停推广 APIResponse
alibaba.scbp.ad.keyword.status.update

关键词启动暂停推广
*/
type AlibabaScbpAdKeywordStatusUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordStatusUpdateResponse `json:"alibaba_scbp_ad_keyword_status_update_response,omitempty"`
}

type AlibabaScbpAdKeywordStatusUpdateResponse struct {

    // 更新关键词推广状态是否成功
    Result   bool `json:"result,omitempty"`

}
