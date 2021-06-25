package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
财务订单回流 APIResponse
alibaba.wdk.finance.order.backflow

星巴克拉取财务订单回流数据
*/
type AlibabaWdkFinanceOrderBackflowAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkFinanceOrderBackflowResponse `json:"alibaba_wdk_finance_order_backflow_response,omitempty"`
}

type AlibabaWdkFinanceOrderBackflowResponse struct {

    // 根据站点名称查询产品
    ApiResult   *AlibabaWdkFinanceOrderBackflowApiResult `json:"api_result,omitempty"`

}
