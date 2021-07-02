package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerFileUploadAPIRequest 如意音箱音频文件长传 API请求
// alibaba.lst.speaker.file.upload
//
// 如意音箱音频文件长传
type AlibabaLstSpeakerFileUploadAPIRequest struct {
	model.Params
	// 数据流
	_fileBytes *model.File
	// 文件类型,audio:音频，advert:广告
	_fileType string
	// 文件ID
	_fileId string
	// md5直
	_md5 string
}

// NewAlibabaLstSpeakerFileUploadRequest 初始化AlibabaLstSpeakerFileUploadAPIRequest对象
func NewAlibabaLstSpeakerFileUploadRequest() *AlibabaLstSpeakerFileUploadAPIRequest {
	return &AlibabaLstSpeakerFileUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.speaker.file.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is FileBytes Setter
// 数据流
func (r *AlibabaLstSpeakerFileUploadAPIRequest) SetFileBytes(_fileBytes *model.File) error {
	r._fileBytes = _fileBytes
	r.Set("file_bytes", _fileBytes)
	return nil
}

// Get FileBytes Getter
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetFileBytes() *model.File {
	return r._fileBytes
}

// Set is FileType Setter
// 文件类型,audio:音频，advert:广告
func (r *AlibabaLstSpeakerFileUploadAPIRequest) SetFileType(_fileType string) error {
	r._fileType = _fileType
	r.Set("file_type", _fileType)
	return nil
}

// Get FileType Getter
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetFileType() string {
	return r._fileType
}

// Set is FileId Setter
// 文件ID
func (r *AlibabaLstSpeakerFileUploadAPIRequest) SetFileId(_fileId string) error {
	r._fileId = _fileId
	r.Set("file_id", _fileId)
	return nil
}

// Get FileId Getter
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetFileId() string {
	return r._fileId
}

// Set is Md5 Setter
// md5直
func (r *AlibabaLstSpeakerFileUploadAPIRequest) SetMd5(_md5 string) error {
	r._md5 = _md5
	r.Set("md5", _md5)
	return nil
}

// Get Md5 Getter
func (r AlibabaLstSpeakerFileUploadAPIRequest) GetMd5() string {
	return r._md5
}
