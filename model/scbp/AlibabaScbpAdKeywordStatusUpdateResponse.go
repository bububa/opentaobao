package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词启动暂停推广 APIResponse
alibaba.scbp.ad.keyword.status.update

关键词启动暂停推广
*/
type AlibabaScbpAdKeywordStatusUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordStatusUpdateResponse
}

type AlibabaScbpAdKeywordStatusUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_status_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新关键词推广状态是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
