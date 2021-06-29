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
    _fileBytes   *model.File
    // 文件类型,audio:音频，advert:广告
    _fileType   string
    // 文件ID
    _fileId   string
    // md5直
    _md5   string
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
func (r *AlibabaLstSpeakerFileUploadRequest) SetFileBytes(_fileBytes *model.File) error {
    r._fileBytes = _fileBytes
    r.Set("file_bytes", _fileBytes)
    return nil
}

// FileBytes Getter
func (r AlibabaLstSpeakerFileUploadRequest) GetFileBytes() *model.File {
    return r._fileBytes
}
// FileType Setter
// 文件类型,audio:音频，advert:广告
func (r *AlibabaLstSpeakerFileUploadRequest) SetFileType(_fileType string) error {
    r._fileType = _fileType
    r.Set("file_type", _fileType)
    return nil
}

// FileType Getter
func (r AlibabaLstSpeakerFileUploadRequest) GetFileType() string {
    return r._fileType
}
// FileId Setter
// 文件ID
func (r *AlibabaLstSpeakerFileUploadRequest) SetFileId(_fileId string) error {
    r._fileId = _fileId
    r.Set("file_id", _fileId)
    return nil
}

// FileId Getter
func (r AlibabaLstSpeakerFileUploadRequest) GetFileId() string {
    return r._fileId
}
// Md5 Setter
// md5直
func (r *AlibabaLstSpeakerFileUploadRequest) SetMd5(_md5 string) error {
    r._md5 = _md5
    r.Set("md5", _md5)
    return nil
}

// Md5 Getter
func (r AlibabaLstSpeakerFileUploadRequest) GetMd5() string {
    return r._md5
}
