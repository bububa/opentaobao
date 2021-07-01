package damai

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/damai"
)

/* 
大麦换验平台-第三方对外开放-看台接口deleteStand 
alibaba.damai.mev.open.deletestand

deleteStand
*/
func AlibabaDamaiMevOpenDeletestand(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletestandAPIRequest, session string) (*damai.AlibabaDamaiMevOpenDeletestandAPIResponse, error) {
    var resp damai.AlibabaDamaiMevOpenDeletestandAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
