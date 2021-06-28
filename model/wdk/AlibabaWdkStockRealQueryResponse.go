package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
仓内实时库存查询 APIResponse
alibaba.wdk.stock.real.query

查询仓内实时库存信息
*/
type AlibabaWdkStockRealQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkStockRealQueryResponse `json:"alibaba_wdk_stock_real_query_response,omitempty"` 
    AlibabaWdkStockRealQueryResponse
}

/* model for simplify = false
type AlibabaWdkStockRealQueryResponse struct {

    // 结果
    
    Result  *struct {
        AlibabaWdkStockRealQueryResultDo  *AlibabaWdkStockRealQueryResultDo `json:"alibaba_wdk_stock_real_query_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkStockRealQueryResponse struct {

    // 结果
    Result   *AlibabaWdkStockRealQueryResultDo `json:"result,omitempty"`

}
