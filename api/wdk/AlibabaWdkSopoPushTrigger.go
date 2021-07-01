package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
猫超共享库存寄售sopo推送触发 
alibaba.wdk.sopo.push.trigger

猫超共享库存寄售sopo触发推送给商家
*/
func AlibabaWdkSopoPushTrigger(clt *core.SDKClient, req *wdk.AlibabaWdkSopoPushTriggerAPIRequest, session string) (*wdk.AlibabaWdkSopoPushTriggerAPIResponse, error) {
    var resp wdk.AlibabaWdkSopoPushTriggerAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
