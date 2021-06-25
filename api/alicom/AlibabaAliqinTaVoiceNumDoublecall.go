package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
聚石塔语音双呼接口 
alibaba.aliqin.ta.voice.num.doublecall

根据传入的主叫号码与被叫号码，由系统依次向主叫号码与被叫号码发起呼叫，双方都应答后，开始一对一通话并开始计费。使用前需要在阿里大于管理中心添加呼叫双方的显示号码。
*/
func AlibabaAliqinTaVoiceNumDoublecall(clt *core.SDKClient, req *alicom.AlibabaAliqinTaVoiceNumDoublecallRequest, session string) (*alicom.AlibabaAliqinTaVoiceNumDoublecallResponse, error) {
    var resp alicom.AlibabaAliqinTaVoiceNumDoublecallAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
