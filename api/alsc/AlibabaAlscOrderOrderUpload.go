package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
订单回流 
alibaba.alsc.order.order.upload

第三方订单回流
*/
func AlibabaAlscOrderOrderUpload(clt *core.SDKClient, req *alsc.AlibabaAlscOrderOrderUploadAPIRequest, session string) (*alsc.AlibabaAlscOrderOrderUploadAPIResponse, error) {
    var resp alsc.AlibabaAlscOrderOrderUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
