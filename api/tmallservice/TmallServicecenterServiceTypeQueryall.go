package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
服务供应链服务类型 
tmall.servicecenter.service.type.queryall

查询天猫服务类型列表
*/
func TmallServicecenterServiceTypeQueryall(clt *core.SDKClient, req *tmallservice.TmallServicecenterServiceTypeQueryallRequest, session string) (*tmallservice.TmallServicecenterServiceTypeQueryallAPIResponse, error) {
    var resp tmallservice.TmallServicecenterServiceTypeQueryallAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
