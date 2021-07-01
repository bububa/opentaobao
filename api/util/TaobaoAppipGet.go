package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
获取ISV发起请求服务器IP 
taobao.appip.get

获取ISV发起请求服务器IP
*/
func TaobaoAppipGet(clt *core.SDKClient, req *util.TaobaoAppipGetAPIRequest, session string) (*util.TaobaoAppipGetAPIResponse, error) {
    var resp util.TaobaoAppipGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
