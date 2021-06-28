package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
取消呼叫 
alibaba.aliqin.fc.voice.num.cancelcall

当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话
*/
func AlibabaAliqinFcVoiceNumCancelcall(clt *core.SDKClient, req *alicom.AlibabaAliqinFcVoiceNumCancelcallRequest, session string) (*alicom.AlibabaAliqinFcVoiceNumCancelcallAPIResponse, error) {
    var resp alicom.AlibabaAliqinFcVoiceNumCancelcallAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
