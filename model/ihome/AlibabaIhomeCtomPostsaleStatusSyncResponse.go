package ihome

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
C2M售后状态同步 APIResponse
alibaba.ihome.ctom.postsale.status.sync

供给三维家同步定制、成品商品售后进度状态
*/
type AlibabaIhomeCtomPostsaleStatusSyncAPIResponse struct {
    model.CommonResponse
    AlibabaIhomeCtomPostsaleStatusSyncResponse
}

type AlibabaIhomeCtomPostsaleStatusSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_ihome_ctom_postsale_status_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // true
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
