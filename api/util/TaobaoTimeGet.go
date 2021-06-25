package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
获取淘宝系统当前时间 
taobao.time.get

获取淘宝系统当前时间
*/
func TaobaoTimeGet(clt *core.SDKClient, req *util.TaobaoTimeGetRequest, session string) (*util.TaobaoTimeGetResponse, error) {
    var resp util.TaobaoTimeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
