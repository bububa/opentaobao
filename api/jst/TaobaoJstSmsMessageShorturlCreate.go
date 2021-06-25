package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
聚石塔营销效果短链生成 
taobao.jst.sms.message.shorturl.create

聚石塔生成淘短链接口
*/
func TaobaoJstSmsMessageShorturlCreate(clt *core.SDKClient, req *jst.TaobaoJstSmsMessageShorturlCreateRequest, session string) (*jst.TaobaoJstSmsMessageShorturlCreateResponse, error) {
    var resp jst.TaobaoJstSmsMessageShorturlCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
