package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除推广单元 APIResponse
alibaba.scbp.ad.group.delete.ad.group.batch

删除推广单元
*/
type AlibabaScbpAdGroupDeleteAdGroupBatchAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdGroupDeleteAdGroupBatchResponse
}

type AlibabaScbpAdGroupDeleteAdGroupBatchResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_group_delete_ad_group_batch_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
