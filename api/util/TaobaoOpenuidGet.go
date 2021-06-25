package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
获取授权账号对应的OpenUid 
taobao.openuid.get

获取授权账号对应的OpenUid
*/
func TaobaoOpenuidGet(clt *core.SDKClient, req *util.TaobaoOpenuidGetRequest, session string) (*util.TaobaoOpenuidGetResponse, error) {
    var resp util.TaobaoOpenuidGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
