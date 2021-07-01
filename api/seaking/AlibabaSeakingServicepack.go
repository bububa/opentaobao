package seaking

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/seaking"
)

/* 
获取海王用户权限包 
alibaba.seaking.servicepack

获取海王用户权限包
*/
func AlibabaSeakingServicepack(clt *core.SDKClient, req *seaking.AlibabaSeakingServicepackAPIRequest, session string) (*seaking.AlibabaSeakingServicepackAPIResponse, error) {
    var resp seaking.AlibabaSeakingServicepackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
