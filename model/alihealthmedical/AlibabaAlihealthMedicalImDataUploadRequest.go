package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方IM图片音频消息上传 APIRequest
alibaba.alihealth.medical.im.data.upload

三方IM图片音频消息上传
*/
type AlibabaAlihealthMedicalImDataUploadRequest struct {
    model.Params

    // request
    uploadDataRequest   *UploadDataRequest 

    // 文件字节流
    file   []*model.File 

}

func NewAlibabaAlihealthMedicalImDataUploadRequest() *AlibabaAlihealthMedicalImDataUploadRequest{
    return &AlibabaAlihealthMedicalImDataUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalImDataUploadRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.im.data.upload"
}

func (r AlibabaAlihealthMedicalImDataUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalImDataUploadRequest) SetUploadDataRequest(uploadDataRequest *UploadDataRequest) error {
    r.uploadDataRequest = uploadDataRequest
    r.Set("upload_data_request", uploadDataRequest)
    return nil
}

func (r AlibabaAlihealthMedicalImDataUploadRequest) GetUploadDataRequest() *UploadDataRequest {
    return r.uploadDataRequest
}

func (r *AlibabaAlihealthMedicalImDataUploadRequest) SetFile(file []*model.File) error {
    r.file = file
    r.Set("file", file)
    return nil
}

func (r AlibabaAlihealthMedicalImDataUploadRequest) GetFile() []*model.File {
    return r.file
}

