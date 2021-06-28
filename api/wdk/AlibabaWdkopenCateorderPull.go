package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商户回传餐饮加工单状态 
alibaba.wdkopen.cateorder.pull

商户回传餐饮加工单状态
*/
func AlibabaWdkopenCateorderPull(clt *core.SDKClient, req *wdk.AlibabaWdkopenCateorderPullRequest, session string) (*wdk.AlibabaWdkopenCateorderPullAPIResponse, error) {
    var resp wdk.AlibabaWdkopenCateorderPullAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
