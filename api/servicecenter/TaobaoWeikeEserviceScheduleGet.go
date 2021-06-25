package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
客服排班信息查询接口 
taobao.weike.eservice.schedule.get

客服排班信息查询接口
*/
func TaobaoWeikeEserviceScheduleGet(clt *core.SDKClient, req *servicecenter.TaobaoWeikeEserviceScheduleGetRequest, session string) (*servicecenter.TaobaoWeikeEserviceScheduleGetResponse, error) {
    var resp servicecenter.TaobaoWeikeEserviceScheduleGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
