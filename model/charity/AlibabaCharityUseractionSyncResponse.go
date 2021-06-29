package charity

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户公益行为同步 APIResponse
alibaba.charity.useraction.sync

外部公益活动，用户公益行为同步
*/
type AlibabaCharityUseractionSyncAPIResponse struct {
    model.CommonResponse
    AlibabaCharityUseractionSyncResponse
}

type AlibabaCharityUseractionSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_charity_useraction_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *ThreehoursResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
