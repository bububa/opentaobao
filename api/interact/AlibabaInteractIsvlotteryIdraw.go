package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
互动到店抽奖 
alibaba.interact.isvlottery.idraw

互动到店抽奖
*/
func AlibabaInteractIsvlotteryIdraw(clt *core.SDKClient, req *interact.AlibabaInteractIsvlotteryIdrawAPIRequest, session string) (*interact.AlibabaInteractIsvlotteryIdrawAPIResponse, error) {
    var resp interact.AlibabaInteractIsvlotteryIdrawAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
