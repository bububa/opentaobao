package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
供应商推送通话结束事件 
alibaba.aliqin.axb.vendor.push.call.release

通话结束挂断的时候，供应商推送通话结束事件给阿里侧
*/
func AlibabaAliqinAxbVendorPushCallRelease(clt *core.SDKClient, req *alicom.AlibabaAliqinAxbVendorPushCallReleaseAPIRequest, session string) (*alicom.AlibabaAliqinAxbVendorPushCallReleaseAPIResponse, error) {
    var resp alicom.AlibabaAliqinAxbVendorPushCallReleaseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
