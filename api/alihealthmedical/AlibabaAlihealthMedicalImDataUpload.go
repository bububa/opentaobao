package alihealthmedical

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthmedical"
)

/* 
三方IM图片音频消息上传 
alibaba.alihealth.medical.im.data.upload

三方IM图片音频消息上传
*/
func AlibabaAlihealthMedicalImDataUpload(clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalImDataUploadAPIRequest, session string) (*alihealthmedical.AlibabaAlihealthMedicalImDataUploadAPIResponse, error) {
    var resp alihealthmedical.AlibabaAlihealthMedicalImDataUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
