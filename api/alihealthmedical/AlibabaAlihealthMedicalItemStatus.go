package alihealthmedical

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthmedical"
)

/* 
商品上下架 
alibaba.alihealth.medical.item.status

医生通三方机构平台进行服务商品上下架操作
*/
func AlibabaAlihealthMedicalItemStatus(clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalItemStatusRequest, session string) (*alihealthmedical.AlibabaAlihealthMedicalItemStatusAPIResponse, error) {
    var resp alihealthmedical.AlibabaAlihealthMedicalItemStatusAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
