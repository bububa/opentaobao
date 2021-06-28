package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
财务订单回流 APIResponse
alibaba.wdk.finance.order.backflow

星巴克拉取财务订单回流数据
*/
type AlibabaWdkFinanceOrderBackflowAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_finance_order_backflow_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 根据站点名称查询产品
    
    ApiResult   *AlibabaWdkFinanceOrderBackflowApiResult `json:"api_result,omitempty" xml:"