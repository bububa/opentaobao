package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
小程序数据回传 
alibaba.alihealth.drug.wxinfo.upload

小程序数据回传
*/
func AlibabaAlihealthDrugWxinfoUpload(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugWxinfoUploadAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugWxinfoUploadAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugWxinfoUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
