package alihealthmedical

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalImDataUploadAPIRequest 三方IM图片音频消息上传 API请求
// alibaba.alihealth.medical.im.data.upload
//
// 三方IM图片音频消息上传
type AlibabaAlihealthMedicalImDataUploadAPIRequest struct {
	model.Params
	// request
	_uploadDataRequest *UploadDataRequest
	// 文件字节流
	_file *model.File
}

// NewAlibabaAlihealthMedicalImDataUploadRequest 初始化AlibabaAlihealthMedicalImDataUploadAPIRequest对象
func NewAlibabaAlihealthMedicalImDataUploadRequest() *AlibabaAlihealthMedicalImDataUploadAPIRequest {
	return &AlibabaAlihealthMedicalImDataUploadAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalImDataUploadAPIRequest) Reset() {
	r._uploadDataRequest = nil
	r._file = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalImDataUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.im.data.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalImDataUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalImDataUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUploadDataRequest is UploadDataRequest Setter
// request
func (r *AlibabaAlihealthMedicalImDataUploadAPIRequest) SetUploadDataRequest(_uploadDataRequest *UploadDataRequest) error {
	r._uploadDataRequest = _uploadDataRequest
	r.Set("upload_data_request", _uploadDataRequest)
	return nil
}

// GetUploadDataRequest UploadDataRequest Getter
func (r AlibabaAlihealthMedicalImDataUploadAPIRequest) GetUploadDataRequest() *UploadDataRequest {
	return r._uploadDataRequest
}

// SetFile is File Setter
// 文件字节流
func (r *AlibabaAlihealthMedicalImDataUploadAPIRequest) SetFile(_file *model.File) error {
	r._file = _file
	r.Set("file", _file)
	return nil
}

// GetFile File Getter
func (r AlibabaAlihealthMedicalImDataUploadAPIRequest) GetFile() *model.File {
	return r._file
}

var poolAlibabaAlihealthMedicalImDataUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalImDataUploadRequest()
	},
}

// GetAlibabaAlihealthMedicalImDataUploadRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalImDataUploadAPIRequest
func GetAlibabaAlihealthMedicalImDataUploadAPIRequest() *AlibabaAlihealthMedicalImDataUploadAPIRequest {
	return poolAlibabaAlihealthMedicalImDataUploadAPIRequest.Get().(*AlibabaAlihealthMedicalImDataUploadAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalImDataUploadAPIRequest 将 AlibabaAlihealthMedicalImDataUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalImDataUploadAPIRequest(v *AlibabaAlihealthMedicalImDataUploadAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalImDataUploadAPIRequest.Put(v)
}
