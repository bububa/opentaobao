package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
RT收箱回传 APIResponse
alibaba.wdk.fulfill.box.post.back.box

RT收箱后，信息同步履约，履约同通知UMS 容器管理
*/
type AlibabaWdkFulfillBoxPostBackBoxAPIResponse struct {
    model.CommonResponse
    AlibabaWdkFulfillBoxPostBackBoxResponse
}

type AlibabaWdkFulfillBoxPostBackBoxResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_fulfill_box_post_back_box_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // fulfillLogisticSingleResult
    
    FulfillLogisticSingleResult   *FulfillLogisticDefaultResult `json:"fulfill_logistic_single_result,omitempty" xml:"fulfill_logistic_single_result,omitempty"`

    
}
