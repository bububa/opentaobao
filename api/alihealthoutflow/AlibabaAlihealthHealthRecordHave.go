package alihealthoutflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthoutflow"
)

/* 
判断用户的慢健康健康档案是否建设完成 
alibaba.alihealth.health.record.have

判断用户的慢健康健康档案是否建设完成
*/
func AlibabaAlihealthHealthRecordHave(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthHealthRecordHaveAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthHealthRecordHaveAPIResponse, error) {
    var resp alihealthoutflow.AlibabaAlihealthHealthRecordHaveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
