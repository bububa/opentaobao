package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
逆向申请接口 
alibaba.wdk.reverse.applyrefund

逆向渲染
*/
func AlibabaWdkReverseApplyrefund(clt *core.SDKClient, req *wdk.AlibabaWdkReverseApplyrefundAPIRequest, session string) (*wdk.AlibabaWdkReverseApplyrefundAPIResponse, error) {
    var resp wdk.AlibabaWdkReverseApplyrefundAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
