package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
聚石塔公众号状态查询 
taobao.jst.sms.status.query

聚石塔公众号状态查询
*/
func TaobaoJstSmsStatusQuery(clt *core.SDKClient, req *jst.TaobaoJstSmsStatusQueryAPIRequest, session string) (*jst.TaobaoJstSmsStatusQueryAPIResponse, error) {
    var resp jst.TaobaoJstSmsStatusQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
