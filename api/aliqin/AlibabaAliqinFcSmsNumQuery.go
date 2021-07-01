package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
短信发送记录查询 
alibaba.aliqin.fc.sms.num.query

短信发送记录查询。
*/
func AlibabaAliqinFcSmsNumQuery(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcSmsNumQueryAPIRequest, session string) (*aliqin.AlibabaAliqinFcSmsNumQueryAPIResponse, error) {
    var resp aliqin.AlibabaAliqinFcSmsNumQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
