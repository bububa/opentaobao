package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据仓code查询仓限单配置 APIResponse
alibaba.wdk.fulfill.config.read.limit.order

根据仓code查询仓限单配置
*/
type AlibabaWdkFulfillConfigReadLimitOrderAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkFulfillConfigReadLimitOrderResponse `json:"alibaba_wdk_fulfill_config_read_limit_order_response,omitempty"`
}

type AlibabaWdkFulfillConfigReadLimitOrderResponse struct {

    // result
    Results   []Json `json:"results,omitempty"`

}
