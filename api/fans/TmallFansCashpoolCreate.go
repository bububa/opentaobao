package fans

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fans"
)

/* 
创建资金池 
tmall.fans.cashpool.create

商家创建资金池接口
*/
func TmallFansCashpoolCreate(clt *core.SDKClient, req *fans.TmallFansCashpoolCreateRequest, session string) (*fans.TmallFansCashpoolCreateAPIResponse, error) {
    var resp fans.TmallFansCashpoolCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
