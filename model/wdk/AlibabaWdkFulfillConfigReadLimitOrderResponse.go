package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据仓code查询仓限单配置 APIResponse
alibaba.wdk.fulfill.config.read.limit.order

根据仓code查询仓限单配置
*/
type AlibabaWdkFulfillConfigReadLimitOrderAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_fulfill_config_read_limit_order_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Results   []string `json:"results,omitempty" xml:"