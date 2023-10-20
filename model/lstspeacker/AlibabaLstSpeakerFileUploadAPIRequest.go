package lstspeacker

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerFileUploadAPIRequest 如意音箱音频文件长传 API请求
// alibaba.lst.speaker.file.upload
//
// 如意音箱音频文件长传
type AlibabaLstSpeakerFileUploadAPIRequest struct {
	model.Params
	// 文件类型,audio:音频，advert:广告
	_fileType string
	// 文件ID
	_fileId string
	// md5直
	_md5 string
	// 数据流
	_fileBytes *model.File
}

// NewAlibabaLstSpeakerFileUploadRequest 初始化AlibabaLstSpeakerFileUploadAPIRequest对象
func NewAlibabaLstSpeakerFileUploadRequest() *AlibabaLstSpeakerFileUploadAPIRequest {
	return &AlibabaLstSpeakerFileUploadAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstSpeakerFileUploadAPIRequest) Reset() {
	r._fileType = ""
	r._fileId = ""
	r._md5 = ""
	r._fileBytes = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.speaker.file.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileType is FileType Setter
// 文件类型,audio:音频，advert:广告
func (r *AlibabaLstSpeakerFileUploadAPIRequest) SetFileType(_fileType string) error {
	r._fileType = _fileType
	r.Set("file_type", _fileType)
	return nil
}

// GetFileType FileType Getter
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetFileType() string {
	return r._fileType
}

// SetFileId is FileId Setter
// 文件ID
func (r *AlibabaLstSpeakerFileUploadAPIRequest) SetFileId(_fileId string) error {
	r._fileId = _fileId
	r.Set("file_id", _fileId)
	return nil
}

// GetFileId FileId Getter
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetFileId() string {
	return r._fileId
}

// SetMd5 is Md5 Setter
// md5直
func (r *AlibabaLstSpeakerFileUploadAPIRequest) SetMd5(_md5 string) error {
	r._md5 = _md5
	r.Set("md5", _md5)
	return nil
}

// GetMd5 Md5 Getter
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetMd5() string {
	return r._md5
}

// SetFileBytes is FileBytes Setter
// 数据流
func (r *AlibabaLstSpeakerFileUploadAPIRequest) SetFileBytes(_fileBytes *model.File) error {
	r._fileBytes = _fileBytes
	r.Set("file_bytes", _fileBytes)
	return nil
}

// GetFileBytes FileBytes Getter
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetFileBytes() *model.File {
	return r._fileBytes
}

var poolAlibabaLstSpeakerFileUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstSpeakerFileUploadRequest()
	},
}

// GetAlibabaLstSpeakerFileUploadRequest 从 sync.Pool 获取 AlibabaLstSpeakerFileUploadAPIRequest
func GetAlibabaLstSpeakerFileUploadAPIRequest() *AlibabaLstSpeakerFileUploadAPIRequest {
	return poolAlibabaLstSpeakerFileUploadAPIRequest.Get().(*AlibabaLstSpeakerFileUploadAPIRequest)
}

// ReleaseAlibabaLstSpeakerFileUploadAPIRequest 将 AlibabaLstSpeakerFileUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstSpeakerFileUploadAPIRequest(v *AlibabaLstSpeakerFileUploadAPIRequest) {
	v.Reset()
	poolAlibabaLstSpeakerFileUploadAPIRequest.Put(v)
}
