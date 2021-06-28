package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片银行图片上传开放接口 APIRequest
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

func NewAlibabaIcbuPhotobankUploadRequest() *AlibabaIcbuPhotobankUploadRequest{
    return &AlibabaIcbuPhotobankUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuPhotobankUploadRequest) GetApiMethodName() string {
    return "alibaba.icbu.photobank.upload"
}

func (r AlibabaIcbuPhotobankUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuPhotobankUploadRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

func (r AlibabaIcbuPhotobankUploadRequest) GetFileName() string {
    return r.fileName
}

func (r *AlibabaIcbuPhotobankUploadRequest) SetGroupId(groupId string) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaIcbuPhotobankUploadRequest) GetGroupId() string {
    return r.groupId
}

func (r *AlibabaIcbuPhotobankUploadRequest) SetImageBytes(imageBytes []*model.File) error {
    r.imageBytes = imageBytes
    r.Set("image_bytes", imageBytes)
    return nil
}

func (r AlibabaIcbuPhotobankUploadRequest) GetImageBytes() []*model.File {
    return r.imageBytes
}

func (r *AlibabaIcbuPhotobankUploadRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

func (r AlibabaIcbuPhotobankUploadRequest) GetExtraContext() string {
    return r.extraContext
}

