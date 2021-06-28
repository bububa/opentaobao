package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
物联卡状态查询 
alibaba.aliqin.fc.iot.cardStatus

物联卡状态查询
*/
func AlibabaAliqinFcIotCardStatus(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotCardStatusRequest, session string) (*aliqin.AlibabaAliqinFcIotCardStatusAPIResponse, error) {
    var resp aliqin.AlibabaAliqinFcIotCardStatusAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
