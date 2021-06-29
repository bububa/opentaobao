package lstvending

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstvending"
)

/* 
自动售卖机交易信息回流 
alibaba.lst.vending.tradeflow.save

自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
*/
func AlibabaLstVendingTradeflowSave(clt *core.SDKClient, req *lstvending.AlibabaLstVendingTradeflowSaveRequest, session string) (*lstvending.AlibabaLstVendingTradeflowSaveAPIResponse, error) {
    var resp lstvending.AlibabaLstVendingTradeflowSaveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
