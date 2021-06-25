package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
交易信息回流 
alibaba.lst.vas.tradeflow.save

自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
*/
func AlibabaLstVasTradeflowSave(clt *core.SDKClient, req *trade.AlibabaLstVasTradeflowSaveRequest, session string) (*trade.AlibabaLstVasTradeflowSaveResponse, error) {
    var resp trade.AlibabaLstVasTradeflowSaveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
