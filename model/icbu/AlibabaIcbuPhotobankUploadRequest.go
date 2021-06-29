package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片银行图片上传开放接口 API请求
alibaba.icbu.photobank.upload

图片银行图片上传开放接口
*/
type AlibabaIcbuPhotobankUploadRequest struct {
    model.Params
    // 上传图片名称
    _fileName   string
    // 上传图片所在分组
    _groupId   string
    // 图片字节数组
    _imageBytes   []*model.File
    // 扩展参数信息,如ICVID
    _extraContext   string
}

// 初始化AlibabaIcbuPhotobankUploadRequest对象
func NewAlibabaIcbuPhotobankUploadRequest() *AlibabaIcbuPhotobankUploadRequest{
    return &AlibabaIcbuPhotobankUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuPhotobankUploadRequest) GetApiMethodName() string {
    return "alibaba.icbu.photobank.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuPhotobankUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FileName Setter
// 上传图片名称
func (r *AlibabaIcbuPhotobankUploadRequest) SetFileName(_fileName string) error {
    r._fileName = _fileName
    r.Set("file_name", _fileName)
    return nil
}

// FileName Getter
func (r AlibabaIcbuPhotobankUploadRequest) GetFileName() string {
    return r._fileName
}
// GroupId Setter
// 上传图片所在分组
func (r *AlibabaIcbuPhotobankUploadRequest) SetGroupId(_groupId string) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r AlibabaIcbuPhotobankUploadRequest) GetGroupId() string {
    return r._groupId
}
// ImageBytes Setter
// 图片字节数组
func (r *AlibabaIcbuPhotobankUploadRequest) SetImageBytes(_imageBytes []*model.File) error {
    r._imageBytes = _imageBytes
    r.Set("image_bytes", _imageBytes)
    return nil
}

// ImageBytes Getter
func (r AlibabaIcbuPhotobankUploadRequest) GetImageBytes() []*model.File {
    return r._imageBytes
}
// ExtraContext Setter
// 扩展参数信息,如ICVID
func (r *AlibabaIcbuPhotobankUploadRequest) SetExtraContext(_extraContext string) error {
    r._extraContext = _extraContext
    r.Set("extra_context", _extraContext)
    return nil
}

// ExtraContext Getter
func (r AlibabaIcbuPhotobankUploadRequest) GetExtraContext() string {
    return r._extraContext
}
