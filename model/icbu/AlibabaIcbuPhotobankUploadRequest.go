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
    fileName   string
    // 上传图片所在分组
    groupId   string
    // 图片字节数组
    imageBytes   []*model.File
    // 扩展参数信息,如ICVID
    extraContext   string
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
func (r *AlibabaIcbuPhotobankUploadRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

// FileName Getter
func (r AlibabaIcbuPhotobankUploadRequest) GetFileName() string {
    return r.fileName
}
// GroupId Setter
// 上传图片所在分组
func (r *AlibabaIcbuPhotobankUploadRequest) SetGroupId(groupId string) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r AlibabaIcbuPhotobankUploadRequest) GetGroupId() string {
    return r.groupId
}
// ImageBytes Setter
// 图片字节数组
func (r *AlibabaIcbuPhotobankUploadRequest) SetImageBytes(imageBytes []*model.File) error {
    r.imageBytes = imageBytes
    r.Set("image_bytes", imageBytes)
    return nil
}

// ImageBytes Getter
func (r AlibabaIcbuPhotobankUploadRequest) GetImageBytes() []*model.File {
    return r.imageBytes
}
// ExtraContext Setter
// 扩展参数信息,如ICVID
func (r *AlibabaIcbuPhotobankUploadRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

// ExtraContext Getter
func (r AlibabaIcbuPhotobankUploadRequest) GetExtraContext() string {
    return r.extraContext
}
