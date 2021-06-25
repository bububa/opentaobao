package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
客服外包订单查询 
taobao.weike.eservice.order.get

用于客服外包中服务商查询订单列表
*/
func TaobaoWeikeEserviceOrderGet(clt *core.SDKClient, req *servicecenter.TaobaoWeikeEserviceOrderGetRequest, session string) (*servicecenter.TaobaoWeikeEserviceOrderGetResponse, error) {
    var resp servicecenter.TaobaoWeikeEserviceOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
