package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstspeakerfileuploadAPIRequest 如意音箱音频文件长传 API请求
// alibaba.lst.speaker.file.upload
//
// 如意音箱音频文件长传
type AlibabalstspeakerfileuploadAPIRequest struct {
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

// NewAlibabalstspeakerfileuploadRequest 初始化AlibabalstspeakerfileuploadAPIRequest对象
func NewAlibabalstspeakerfileuploadRequest() *AlibabalstspeakerfileuploadAPIRequest {
	return &AlibabalstspeakerfileuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstspeakerfileuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.speaker.file.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstspeakerfileuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstspeakerfileuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileType is FileType Setter
// 文件类型,audio:音频，advert:广告
func (r *AlibabalstspeakerfileuploadAPIRequest) SetFileType(_fileType string) error {
	r._fileType = _fileType
	r.Set("file_type", _fileType)
	return nil
}

// GetFileType FileType Getter
func (r AlibabalstspeakerfileuploadAPIRequest) GetFileType() string {
	return r._fileType
}

// SetFileId is FileId Setter
// 文件ID
func (r *AlibabalstspeakerfileuploadAPIRequest) SetFileId(_fileId string) error {
	r._fileId = _fileId
	r.Set("file_id", _fileId)
	return nil
}

// GetFileId FileId Getter
func (r AlibabalstspeakerfileuploadAPIRequest) GetFileId() string {
	return r._fileId
}

// SetMd5 is Md5 Setter
// md5直
func (r *AlibabalstspeakerfileuploadAPIRequest) SetMd5(_md5 string) error {
	r._md5 = _md5
	r.Set("md5", _md5)
	return nil
}

// GetMd5 Md5 Getter
func (r AlibabalstspeakerfileuploadAPIRequest) GetMd5() string {
	return r._md5
}

// SetFileBytes is FileBytes Setter
// 数据流
func (r *AlibabalstspeakerfileuploadAPIRequest) SetFileBytes(_fileBytes *model.File) error {
	r._fileBytes = _fileBytes
	r.Set("file_bytes", _fileBytes)
	return nil
}

// GetFileBytes FileBytes Getter
func (r AlibabalstspeakerfileuploadAPIRequest) GetFileBytes() *model.File {
	return r._fileBytes
}
