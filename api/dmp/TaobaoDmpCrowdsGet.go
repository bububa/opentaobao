package dmp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/dmp"
)

/* 
查询人群服务 
taobao.dmp.crowds.get

查询人群服务
*/
func TaobaoDmpCrowdsGet(clt *core.SDKClient, req *dmp.TaobaoDmpCrowdsGetRequest, session string) (*dmp.TaobaoDmpCrowdsGetAPIResponse, error) {
    var resp dmp.TaobaoDmpCrowdsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
