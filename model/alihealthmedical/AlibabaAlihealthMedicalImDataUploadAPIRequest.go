package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalimdatauploadAPIRequest 三方IM图片音频消息上传 API请求
// alibaba.alihealth.medical.im.data.upload
//
// 三方IM图片音频消息上传
type AlibabaalihealthmedicalimdatauploadAPIRequest struct {
	model.Params
	// request
	_uploadDataRequest *UploadDataRequest
	// 文件字节流
	_file *model.File
}

// NewAlibabaalihealthmedicalimdatauploadRequest 初始化AlibabaalihealthmedicalimdatauploadAPIRequest对象
func NewAlibabaalihealthmedicalimdatauploadRequest() *AlibabaalihealthmedicalimdatauploadAPIRequest {
	return &AlibabaalihealthmedicalimdatauploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalimdatauploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.im.data.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalimdatauploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalimdatauploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUploadDataRequest is UploadDataRequest Setter
// request
func (r *AlibabaalihealthmedicalimdatauploadAPIRequest) SetUploadDataRequest(_uploadDataRequest *UploadDataRequest) error {
	r._uploadDataRequest = _uploadDataRequest
	r.Set("upload_data_request", _uploadDataRequest)
	return nil
}

// GetUploadDataRequest UploadDataRequest Getter
func (r AlibabaalihealthmedicalimdatauploadAPIRequest) GetUploadDataRequest() *UploadDataRequest {
	return r._uploadDataRequest
}

// SetFile is File Setter
// 文件字节流
func (r *AlibabaalihealthmedicalimdatauploadAPIRequest) SetFile(_file *model.File) error {
	r._file = _file
	r.Set("file", _file)
	return nil
}

// GetFile File Getter
func (r AlibabaalihealthmedicalimdatauploadAPIRequest) GetFile() *model.File {
	return r._file
}
