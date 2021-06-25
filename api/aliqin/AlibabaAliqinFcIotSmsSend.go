package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
物联网短信发送 
alibaba.aliqin.fc.iot.sms.send

发送物联网短信，只允许使用物联网短信模板
*/
func AlibabaAliqinFcIotSmsSend(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotSmsSendRequest, session string) (*aliqin.AlibabaAliqinFcIotSmsSendResponse, error) {
    var resp aliqin.AlibabaAliqinFcIotSmsSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
