package lstspeacker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
如意音箱音频文件长传 API请求
alibaba.lst.speaker.file.upload

如意音箱音频文件长传
*/
type AlibabaLstSpeakerFileUploadRequest struct {
    model.Params
    // 数据流
    fileBytes   []*model.File
    // 文件类型,audio:音频，advert:广告
    fileType   string
    // 文件ID
    fileId   string
    // md5直
    md5   string
}

// 初始化AlibabaLstSpeakerFileUploadRequest对象
func NewAlibabaLstSpeakerFileUploadRequest() *AlibabaLstSpeakerFileUploadRequest{
    return &AlibabaLstSpeakerFileUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerFileUploadRequest) GetApiMethodName() string {
    return "alibaba.lst.speaker.file.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerFileUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FileBytes Setter
// 数据流
func (r *AlibabaLstSpeakerFileUploadRequest) SetFileBytes(fileBytes []*model.File) error {
    r.fileBytes = fileBytes
    r.Set("file_bytes", fileBytes)
    return nil
}

// FileBytes Getter
func (r AlibabaLstSpeakerFileUploadRequest) GetFileBytes() []*model.File {
    return r.fileBytes
}
// FileType Setter
// 文件类型,audio:音频，advert:广告
func (r *AlibabaLstSpeakerFileUploadRequest) SetFileType(fileType string) error {
    r.fileType = fileType
    r.Set("file_type", fileType)
    return nil
}

// FileType Getter
func (r AlibabaLstSpeakerFileUploadRequest) GetFileType() string {
    return r.fileType
}
// FileId Setter
// 文件ID
func (r *AlibabaLstSpeakerFileUploadRequest) SetFileId(fileId string) error {
    r.fileId = fileId
    r.Set("file_id", fileId)
    return nil
}

// FileId Getter
func (r AlibabaLstSpeakerFileUploadRequest) GetFileId() string {
    return r.fileId
}
// Md5 Setter
// md5直
func (r *AlibabaLstSpeakerFileUploadRequest) SetMd5(md5 string) error {
    r.md5 = md5
    r.Set("md5", md5)
    return nil
}

// Md5 Getter
func (r AlibabaLstSpeakerFileUploadRequest) GetMd5() string {
    return r.md5
}
