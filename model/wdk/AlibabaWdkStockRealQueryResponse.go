package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内实时库存查询 APIResponse
alibaba.wdk.stock.real.query

查询仓内实时库存信息
*/
type AlibabaWdkStockRealQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_stock_real_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkStockRealQueryResultDo `json:"result,omitempty" xml:"