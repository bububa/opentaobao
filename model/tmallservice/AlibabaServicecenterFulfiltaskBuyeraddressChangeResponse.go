package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改消费者服务地址 APIResponse
alibaba.servicecenter.fulfiltask.buyeraddress.change

当消费者反馈自己的服务地址错误时，可以电话联系服务商修改为正确地址，服务商只能修改派给自己的单子
*/
type AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_servicecenter_fulfiltask_buyeraddress_change_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaServicecenterFulfiltaskBuyeraddressChangeResult `json:"result,omitempty" xml:"