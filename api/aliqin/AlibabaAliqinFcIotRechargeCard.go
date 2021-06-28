package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
按终端号订购增值业务 
alibaba.aliqin.fc.iot.rechargeCard

按终端号订购增值业务
*/
func AlibabaAliqinFcIotRechargeCard(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotRechargeCardRequest, session string) (*aliqin.AlibabaAliqinFcIotRechargeCardAPIResponse, error) {
    var resp aliqin.AlibabaAliqinFcIotRechargeCardAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
