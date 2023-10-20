package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicySpecialCompressionUploadAPIRequest 大批量上传特殊类型的单程/往返政策 API请求
// alitrip.policy.special.compression.upload
//
// 大批量上传特殊类型的单程/往返政策
type AlitripPolicySpecialCompressionUploadAPIRequest struct {
	model.Params
	// alitrip.policy.special.upload接口的参数经过压缩后得到的字节数组，生成的文件
	_file *model.File
}

// NewAlitripPolicySpecialCompressionUploadRequest 初始化AlitripPolicySpecialCompressionUploadAPIRequest对象
func NewAlitripPolicySpecialCompressionUploadRequest() *AlitripPolicySpecialCompressionUploadAPIRequest {
	return &AlitripPolicySpecialCompressionUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripPolicySpecialCompressionUploadAPIRequest) Reset() {
	r._file = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPolicySpecialCompressionUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.special.compression.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPolicySpecialCompressionUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPolicySpecialCompressionUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFile is File Setter
// alitrip.policy.special.upload接口的参数经过压缩后得到的字节数组，生成的文件
func (r *AlitripPolicySpecialCompressionUploadAPIRequest) SetFile(_file *model.File) error {
	r._file = _file
	r.Set("file", _file)
	return nil
}

// GetFile File Getter
func (r AlitripPolicySpecialCompressionUploadAPIRequest) GetFile() *model.File {
	return r._file
}

var poolAlitripPolicySpecialCompressionUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripPolicySpecialCompressionUploadRequest()
	},
}

// GetAlitripPolicySpecialCompressionUploadRequest 从 sync.Pool 获取 AlitripPolicySpecialCompressionUploadAPIRequest
func GetAlitripPolicySpecialCompressionUploadAPIRequest() *AlitripPolicySpecialCompressionUploadAPIRequest {
	return poolAlitripPolicySpecialCompressionUploadAPIRequest.Get().(*AlitripPolicySpecialCompressionUploadAPIRequest)
}

// ReleaseAlitripPolicySpecialCompressionUploadAPIRequest 将 AlitripPolicySpecialCompressionUploadAPIRequest 放入 sync.Pool
func ReleaseAlitripPolicySpecialCompressionUploadAPIRequest(v *AlitripPolicySpecialCompressionUploadAPIRequest) {
	v.Reset()
	poolAlitripPolicySpecialCompressionUploadAPIRequest.Put(v)
}
