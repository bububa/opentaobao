package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
聚石塔短链信息查询 
taobao.jst.sms.message.shorturl.query

聚石塔短链信息查询
*/
func TaobaoJstSmsMessageShorturlQuery(clt *core.SDKClient, req *jst.TaobaoJstSmsMessageShorturlQueryRequest, session string) (*jst.TaobaoJstSmsMessageShorturlQueryResponse, error) {
    var resp jst.TaobaoJstSmsMessageShorturlQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
