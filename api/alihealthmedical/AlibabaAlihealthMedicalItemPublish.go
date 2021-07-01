package alihealthmedical

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthmedical"
)

/* 
三方入驻-开通服务 
alibaba.alihealth.medical.item.publish

三方入驻-开通服务
*/
func AlibabaAlihealthMedicalItemPublish(clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalItemPublishAPIRequest, session string) (*alihealthmedical.AlibabaAlihealthMedicalItemPublishAPIResponse, error) {
    var resp alihealthmedical.AlibabaAlihealthMedicalItemPublishAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
