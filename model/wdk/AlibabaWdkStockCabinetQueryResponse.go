package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内实时库位库存查询 APIResponse
alibaba.wdk.stock.cabinet.query

查询仓内实时库位库存信息
*/
type AlibabaWdkStockCabinetQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_stock_cabinet_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaWdkStockCabinetQueryResultDo `json:"result,omitempty" xml:"