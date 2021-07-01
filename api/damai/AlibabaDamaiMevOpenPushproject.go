package damai

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/damai"
)

/* 
大麦换验平台-第三方对外开放-项目接口pushProject 
alibaba.damai.mev.open.pushproject

pushProject
*/
func AlibabaDamaiMevOpenPushproject(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushprojectAPIRequest, session string) (*damai.AlibabaDamaiMevOpenPushprojectAPIResponse, error) {
    var resp damai.AlibabaDamaiMevOpenPushprojectAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
