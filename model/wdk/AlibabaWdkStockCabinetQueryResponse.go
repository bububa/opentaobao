package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
仓内实时库位库存查询 APIResponse
alibaba.wdk.stock.cabinet.query

查询仓内实时库位库存信息
*/
type AlibabaWdkStockCabinetQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkStockCabinetQueryResponse `json:"alibaba_wdk_stock_cabinet_query_response,omitempty"` 
    AlibabaWdkStockCabinetQueryResponse
}

/* model for simplify = false
type AlibabaWdkStockCabinetQueryResponse struct {

    // result
    
    Result  *struct {
        AlibabaWdkStockCabinetQueryResultDo  *AlibabaWdkStockCabinetQueryResultDo `json:"alibaba_wdk_stock_cabinet_query_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkStockCabinetQueryResponse struct {

    // result
    Result   *AlibabaWdkStockCabinetQueryResultDo `json:"result,omitempty"`

}
