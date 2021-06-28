package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
需求订单查询接口 
taobao.weike.subscinfo.get

需求订单查询接口
*/
func TaobaoWeikeSubscinfoGet(clt *core.SDKClient, req *servicecenter.TaobaoWeikeSubscinfoGetRequest, session string) (*servicecenter.TaobaoWeikeSubscinfoGetAPIResponse, error) {
    var resp servicecenter.TaobaoWeikeSubscinfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
