package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
聚石塔新短信发送接口 
taobao.jst.sms.message.direct.batchsend

聚石塔所见即所得的短信发送接口
*/
func TaobaoJstSmsMessageDirectBatchsend(clt *core.SDKClient, req *jst.TaobaoJstSmsMessageDirectBatchsendAPIRequest, session string) (*jst.TaobaoJstSmsMessageDirectBatchsendAPIResponse, error) {
    var resp jst.TaobaoJstSmsMessageDirectBatchsendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
