package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyNormalCompressionUploadAPIRequest 大批量上传普通类型的单程/往返政策 API请求
// alitrip.policy.normal.compression.upload
//
// 大批量上传普通类型的单程/往返政策
type AlitripPolicyNormalCompressionUploadAPIRequest struct {
	model.Params
	// alitrip.policy.normal.upload接口的参数经过压缩后得到的字节数组，生成的文件
	_file *model.File
}

// NewAlitripPolicyNormalCompressionUploadRequest 初始化AlitripPolicyNormalCompressionUploadAPIRequest对象
func NewAlitripPolicyNormalCompressionUploadRequest() *AlitripPolicyNormalCompressionUploadAPIRequest {
	return &AlitripPolicyNormalCompressionUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPolicyNormalCompressionUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.normal.compression.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPolicyNormalCompressionUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPolicyNormalCompressionUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFile is File Setter
// alitrip.policy.normal.upload接口的参数经过压缩后得到的字节数组，生成的文件
func (r *AlitripPolicyNormalCompressionUploadAPIRequest) SetFile(_file *model.File) error {
	r._file = _file
	r.Set("file", _file)
	return nil
}

// GetFile File Getter
func (r AlitripPolicyNormalCompressionUploadAPIRequest) GetFile() *model.File {
	return r._file
}
