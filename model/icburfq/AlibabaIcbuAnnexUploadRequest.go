package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传附件获取附件files_str API请求
alibaba.icbu.annex.upload

上传附件获取附件files_str
*/
type AlibabaIcbuAnnexUploadRequest struct {
    model.Params
    // 文件名
    _fileName   string
    // 文件字节流
    _fileInputStreamBytes   []*model.File
    // 来源
    _source   string
}

// 初始化AlibabaIcbuAnnexUploadRequest对象
func NewAlibabaIcbuAnnexUploadRequest() *AlibabaIcbuAnnexUploadRequest{
    return &AlibabaIcbuAnnexUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuAnnexUploadRequest) GetApiMethodName() string {
    return "alibaba.icbu.annex.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuAnnexUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FileName Setter
// 文件名
func (r *AlibabaIcbuAnnexUploadRequest) SetFileName(_fileName string) error {
    r._fileName = _fileName
    r.Set("file_name", _fileName)
    return nil
}

// FileName Getter
func (r AlibabaIcbuAnnexUploadRequest) GetFileName() string {
    return r._fileName
}
// FileInputStreamBytes Setter
// 文件字节流
func (r *AlibabaIcbuAnnexUploadRequest) SetFileInputStreamBytes(_fileInputStreamBytes []*model.File) error {
    r._fileInputStreamBytes = _fileInputStreamBytes
    r.Set("file_input_stream_bytes", _fileInputStreamBytes)
    return nil
}

// FileInputStreamBytes Getter
func (r AlibabaIcbuAnnexUploadRequest) GetFileInputStreamBytes() []*model.File {
    return r._fileInputStreamBytes
}
// Source Setter
// 来源
func (r *AlibabaIcbuAnnexUploadRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaIcbuAnnexUploadRequest) GetSource() string {
    return r._source
}