package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
删除数据推送用户 
taobao.jushita.jdp.user.delete

删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
*/
func TaobaoJushitaJdpUserDelete(clt *core.SDKClient, req *jst.TaobaoJushitaJdpUserDeleteAPIRequest, session string) (*jst.TaobaoJushitaJdpUserDeleteAPIResponse, error) {
    var resp jst.TaobaoJushitaJdpUserDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
