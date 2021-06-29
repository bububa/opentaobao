package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方IM图片音频消息上传 API请求
alibaba.alihealth.medical.im.data.upload

三方IM图片音频消息上传
*/
type AlibabaAlihealthMedicalImDataUploadRequest struct {
    model.Params
    // request
    _uploadDataRequest   *UploadDataRequest
    // 文件字节流
    _file   []*model.File
}

// 初始化AlibabaAlihealthMedicalImDataUploadRequest对象
func NewAlibabaAlihealthMedicalImDataUploadRequest() *AlibabaAlihealthMedicalImDataUploadRequest{
    return &AlibabaAlihealthMedicalImDataUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalImDataUploadRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.im.data.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalImDataUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UploadDataRequest Setter
// request
func (r *AlibabaAlihealthMedicalImDataUploadRequest) SetUploadDataRequest(_uploadDataRequest *UploadDataRequest) error {
    r._uploadDataRequest = _uploadDataRequest
    r.Set("upload_data_request", _uploadDataRequest)
    return nil
}

// UploadDataRequest Getter
func (r AlibabaAlihealthMedicalImDataUploadRequest) GetUploadDataRequest() *UploadDataRequest {
    return r._uploadDataRequest
}
// File Setter
// 文件字节流
func (r *AlibabaAlihealthMedicalImDataUploadRequest) SetFile(_file []*model.File) error {
    r._file = _file
    r.Set("file", _file)
    return nil
}

// File Getter
func (r AlibabaAlihealthMedicalImDataUploadRequest) GetFile() []*model.File {
    return r._file
}
