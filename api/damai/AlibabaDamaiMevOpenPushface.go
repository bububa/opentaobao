package damai

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/damai"
)

/* 
大麦换验平台-第三方对外开放-票面接口pushFace 
alibaba.damai.mev.open.pushface

pushFace
*/
func AlibabaDamaiMevOpenPushface(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushfaceRequest, session string) (*damai.AlibabaDamaiMevOpenPushfaceAPIResponse, error) {
    var resp damai.AlibabaDamaiMevOpenPushfaceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
