package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
获取开放平台出口IP段 
taobao.top.ipout.get

获取开放平台出口IP段
*/
func TaobaoTopIpoutGet(clt *core.SDKClient, req *util.TaobaoTopIpoutGetRequest, session string) (*util.TaobaoTopIpoutGetResponse, error) {
    var resp util.TaobaoTopIpoutGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
