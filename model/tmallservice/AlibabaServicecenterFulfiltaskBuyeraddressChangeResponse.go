package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改消费者服务地址 APIResponse
alibaba.servicecenter.fulfiltask.buyeraddress.change

当消费者反馈自己的服务地址错误时，可以电话联系服务商修改为正确地址，服务商只能修改派给自己的单子
*/
type AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaServicecenterFulfiltaskBuyeraddressChangeResponse `json:"alibaba_servicecenter_fulfiltask_buyeraddress_change_response,omitempty"`
}

type AlibabaServicecenterFulfiltaskBuyeraddressChangeResponse struct {

    // 接口返回model
    Result   *AlibabaServicecenterFulfiltaskBuyeraddressChangeResult `json:"result,omitempty"`

}
