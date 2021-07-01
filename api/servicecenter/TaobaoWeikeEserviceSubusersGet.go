package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
客服外包订单分配的商家子账号列表 
taobao.weike.eservice.subusers.get

获取客服外包订单分配的商家子账号列表，以及授权状态
*/
func TaobaoWeikeEserviceSubusersGet(clt *core.SDKClient, req *servicecenter.TaobaoWeikeEserviceSubusersGetAPIRequest, session string) (*servicecenter.TaobaoWeikeEserviceSubusersGetAPIResponse, error) {
    var resp servicecenter.TaobaoWeikeEserviceSubusersGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
