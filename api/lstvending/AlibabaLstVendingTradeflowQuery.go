package lstvending

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstvending"
)

/* 
自动售卖机交易流水查询 
alibaba.lst.vending.tradeflow.query

零售通自动售卖机交易流水查询接口，品牌商通过此接口同步商品交易数据。
*/
func AlibabaLstVendingTradeflowQuery(clt *core.SDKClient, req *lstvending.AlibabaLstVendingTradeflowQueryAPIRequest, session string) (*lstvending.AlibabaLstVendingTradeflowQueryAPIResponse, error) {
    var resp lstvending.AlibabaLstVendingTradeflowQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
