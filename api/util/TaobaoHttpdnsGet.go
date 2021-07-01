package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
TOPDNS配置 
taobao.httpdns.get

获取TOP DNS配置
*/
func TaobaoHttpdnsGet(clt *core.SDKClient, req *util.TaobaoHttpdnsGetAPIRequest, session string) (*util.TaobaoHttpdnsGetAPIResponse, error) {
    var resp util.TaobaoHttpdnsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
