package mc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mc"
)

/* 
云码充电宝投放链路 
tmall.mc.task.charge.launch

云码充电宝投放链路，用于判断用户是否有匹配的投放计划
*/
func TmallMcTaskChargeLaunch(clt *core.SDKClient, req *mc.TmallMcTaskChargeLaunchAPIRequest, session string) (*mc.TmallMcTaskChargeLaunchAPIResponse, error) {
    var resp mc.TmallMcTaskChargeLaunchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
