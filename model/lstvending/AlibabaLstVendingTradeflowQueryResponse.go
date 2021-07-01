package lstvending

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售卖机交易流水查询 API返回值 
alibaba.lst.vending.tradeflow.query

零售通自动售卖机交易流水查询接口，品牌商通过此接口同步商品交易数据。
*/
type AlibabaLstVendingTradeflowQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstVendingTradeflowQueryResponse
}

// 自动售卖机交易流水查询 成功返回结果
type AlibabaLstVendingTradeflowQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_vending_tradeflow_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果集
    Result   *ResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
