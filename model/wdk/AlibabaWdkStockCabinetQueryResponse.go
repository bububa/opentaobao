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
    Response *AlibabaWdkStockCabinetQueryResponse `json:"alibaba_wdk_stock_cabinet_query_response,omitempty"`
}

type AlibabaWdkStockCabinetQueryResponse struct {

    // result
    Result   *AlibabaWdkStockCabinetQueryResultDo `json:"result,omitempty"`

}
