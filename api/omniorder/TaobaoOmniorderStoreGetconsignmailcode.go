package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
全渠道门店物流菜鸟裹裹取号 
taobao.omniorder.store.getconsignmailcode

用于ISV获取全渠道门店物流订单菜鸟裹裹门店的物流快递取号
*/
func TaobaoOmniorderStoreGetconsignmailcode(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreGetconsignmailcodeAPIRequest, session string) (*omniorder.TaobaoOmniorderStoreGetconsignmailcodeAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderStoreGetconsignmailcodeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
