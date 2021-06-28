package mtopopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mtopopen"
)

/* 
资源位数据推送接口 
alibaba.interact.aopdata.register

提供给isv，查询以及推送浮层资源位的三方活动数据
*/
func AlibabaInteractAopdataRegister(clt *core.SDKClient, req *mtopopen.AlibabaInteractAopdataRegisterRequest, session string) (*mtopopen.AlibabaInteractAopdataRegisterAPIResponse, error) {
    var resp mtopopen.AlibabaInteractAopdataRegisterAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
