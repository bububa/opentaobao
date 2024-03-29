package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateUploadAPIRequest 协同单附件凭证上传 API请求
// alitrip.agent.coordinate.upload
//
// 协同单附件凭证上传
type AlitripAgentCoordinateUploadAPIRequest struct {
	model.Params
	// 附件Byte[]
	_fileContent *model.File
}

// NewAlitripAgentCoordinateUploadRequest 初始化AlitripAgentCoordinateUploadAPIRequest对象
func NewAlitripAgentCoordinateUploadRequest() *AlitripAgentCoordinateUploadAPIRequest {
	return &AlitripAgentCoordinateUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentCoordinateUploadAPIRequest) Reset() {
	r._fileContent = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentCoordinateUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentCoordinateUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentCoordinateUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileContent is FileContent Setter
// 附件Byte[]
func (r *AlitripAgentCoordinateUploadAPIRequest) SetFileContent(_fileContent *model.File) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// GetFileContent FileContent Getter
func (r AlitripAgentCoordinateUploadAPIRequest) GetFileContent() *model.File {
	return r._fileContent
}

var poolAlitripAgentCoordinateUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentCoordinateUploadRequest()
	},
}

// GetAlitripAgentCoordinateUploadRequest 从 sync.Pool 获取 AlitripAgentCoordinateUploadAPIRequest
func GetAlitripAgentCoordinateUploadAPIRequest() *AlitripAgentCoordinateUploadAPIRequest {
	return poolAlitripAgentCoordinateUploadAPIRequest.Get().(*AlitripAgentCoordinateUploadAPIRequest)
}

// ReleaseAlitripAgentCoordinateUploadAPIRequest 将 AlitripAgentCoordinateUploadAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentCoordinateUploadAPIRequest(v *AlitripAgentCoordinateUploadAPIRequest) {
	v.Reset()
	poolAlitripAgentCoordinateUploadAPIRequest.Put(v)
}
