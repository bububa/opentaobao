package damai

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/damai"
)

/* 
大麦换验平台-第三方对外开放-场次接口deletePerform 
alibaba.damai.mev.open.deleteperform

deletePerform
*/
func AlibabaDamaiMevOpenDeleteperform(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeleteperformAPIRequest, session string) (*damai.AlibabaDamaiMevOpenDeleteperformAPIResponse, error) {
    var resp damai.AlibabaDamaiMevOpenDeleteperformAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
