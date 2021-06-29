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
    fileName   string
    // 文件字节流
    fileInputStreamBytes   []*model.File
    // 来源
    source   string
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
func (r *AlibabaIcbuAnnexUploadRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

// FileName Getter
func (r AlibabaIcbuAnnexUploadRequest) GetFileName() string {
    return r.fileName
}
// FileInputStreamBytes Setter
// 文件字节流
func (r *AlibabaIcbuAnnexUploadRequest) SetFileInputStreamBytes(fileInputStreamBytes []*model.File) error {
    r.fileInputStreamBytes = fileInputStreamBytes
    r.Set("file_input_stream_bytes", fileInputStreamBytes)
    return nil
}

// FileInputStreamBytes Getter
func (r AlibabaIcbuAnnexUploadRequest) GetFileInputStreamBytes() []*model.File {
    return r.fileInputStreamBytes
}
// Source Setter
// 来源
func (r *AlibabaIcbuAnnexUploadRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaIcbuAnnexUploadRequest) GetSource() string {
    return r.source
}
