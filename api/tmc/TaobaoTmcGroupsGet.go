package tmc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmc"
)

/* 
获取自定义用户分组列表 
taobao.tmc.groups.get

获取自定义用户分组列表
*/
func TaobaoTmcGroupsGet(clt *core.SDKClient, req *tmc.TaobaoTmcGroupsGetAPIRequest, session string) (*tmc.TaobaoTmcGroupsGetAPIResponse, error) {
    var resp tmc.TaobaoTmcGroupsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
