package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫好房楼盘评测同步 APIResponse
alibaba.alihouse.newhome.review.sync

接受楼盘测评图文信息
*/
type AlibabaAlihouseNewhomeReviewSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeReviewSyncResponse
}

type AlibabaAlihouseNewhomeReviewSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_review_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeReviewSyncResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
