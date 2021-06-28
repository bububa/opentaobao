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
    // Response *AlibabaWdkFulfillConfigReadLimitOrderResponse `json:"alibaba_wdk_fulfill_config_read_limit_order_response,omitempty"` 
    AlibabaWdkFulfillConfigReadLimitOrderResponse
}

/* model for simplify = false
type AlibabaWdkFulfillConfigReadLimitOrderResponse struct {

    // result
    
    Results  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type AlibabaWdkFulfillConfigReadLimitOrderResponse struct {

    // result
    Results   []string `json:"results,omitempty"`

}
