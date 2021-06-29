package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
楼盘测评草稿状态同步 APIResponse
alibaba.alihouse.newhome.review.changestatus

楼盘测评草稿状态更新
*/
type AlibabaAlihouseNewhomeReviewChangestatusAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeReviewChangestatusResponse
}

type AlibabaAlihouseNewhomeReviewChangestatusResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_review_changestatus_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeReviewChangestatusResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
