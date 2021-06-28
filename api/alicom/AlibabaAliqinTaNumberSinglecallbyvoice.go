package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
根据号码tts单呼 
alibaba.aliqin.ta.number.singlecallbyvoice

根据号码语音单呼
*/
func AlibabaAliqinTaNumberSinglecallbyvoice(clt *core.SDKClient, req *alicom.AlibabaAliqinTaNumberSinglecallbyvoiceRequest, session string) (*alicom.AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse, error) {
    var resp alicom.AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
