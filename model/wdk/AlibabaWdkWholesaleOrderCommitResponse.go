package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
盒马帮采购确认订单接口 APIResponse
alibaba.wdk.wholesale.order.commit

盒马帮采购确认订单接口
*/
type AlibabaWdkWholesaleOrderCommitAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkWholesaleOrderCommitResponse `json:"alibaba_wdk_wholesale_order_commit_response,omitempty"`
}

type AlibabaWdkWholesaleOrderCommitResponse struct {

    // result
    Result   *AlibabaWdkWholesaleOrderCommitApiResult `json:"result,omitempty"`

}
