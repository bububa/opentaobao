package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
AXB短信托收推送接口 
alibaba.aliqin.axb.vendor.sms.intercept

用于给供应商推送需要托收的短信
*/
func AlibabaAliqinAxbVendorSmsIntercept(clt *core.SDKClient, req *alicom.AlibabaAliqinAxbVendorSmsInterceptRequest, session string) (*alicom.AlibabaAliqinAxbVendorSmsInterceptResponse, error) {
    var resp alicom.AlibabaAliqinAxbVendorSmsInterceptAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
