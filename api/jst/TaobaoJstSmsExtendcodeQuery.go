package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
聚石塔扩展码查询 
taobao.jst.sms.extendcode.query

聚石塔扩展码查询
*/
func TaobaoJstSmsExtendcodeQuery(clt *core.SDKClient, req *jst.TaobaoJstSmsExtendcodeQueryRequest, session string) (*jst.TaobaoJstSmsExtendcodeQueryAPIResponse, error) {
    var resp jst.TaobaoJstSmsExtendcodeQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
