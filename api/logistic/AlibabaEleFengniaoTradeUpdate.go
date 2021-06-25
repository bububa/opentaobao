package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
更新蜂鸟扣费状态 
alibaba.ele.fengniao.trade.update

汇金扣费成功后，回调该接口更新扣费状态
*/
func AlibabaEleFengniaoTradeUpdate(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoTradeUpdateRequest, session string) (*logistic.AlibabaEleFengniaoTradeUpdateResponse, error) {
    var resp logistic.AlibabaEleFengniaoTradeUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
