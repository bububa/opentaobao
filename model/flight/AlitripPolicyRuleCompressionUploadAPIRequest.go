package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyRuleCompressionUploadAPIRequest 大批量上传规则类型的单程/往返政策 API请求
// alitrip.policy.rule.compression.upload
//
// 大批量上传规则类型的单程/往返政策
type AlitripPolicyRuleCompressionUploadAPIRequest struct {
	model.Params
	// alitrip.policy.rule.upload接口的参数经过压缩后得到的字节数组，生成的文件
	_file *model.File
}

// NewAlitripPolicyRuleCompressionUploadRequest 初始化AlitripPolicyRuleCompressionUploadAPIRequest对象
func NewAlitripPolicyRuleCompressionUploadRequest() *AlitripPolicyRuleCompressionUploadAPIRequest {
	return &AlitripPolicyRuleCompressionUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripPolicyRuleCompressionUploadAPIRequest) Reset() {
	r._file = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPolicyRuleCompressionUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.rule.compression.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPolicyRuleCompressionUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPolicyRuleCompressionUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFile is File Setter
// alitrip.policy.rule.upload接口的参数经过压缩后得到的字节数组，生成的文件
func (r *AlitripPolicyRuleCompressionUploadAPIRequest) SetFile(_file *model.File) error {
	r._file = _file
	r.Set("file", _file)
	return nil
}

// GetFile File Getter
func (r AlitripPolicyRuleCompressionUploadAPIRequest) GetFile() *model.File {
	return r._file
}

var poolAlitripPolicyRuleCompressionUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripPolicyRuleCompressionUploadRequest()
	},
}

// GetAlitripPolicyRuleCompressionUploadRequest 从 sync.Pool 获取 AlitripPolicyRuleCompressionUploadAPIRequest
func GetAlitripPolicyRuleCompressionUploadAPIRequest() *AlitripPolicyRuleCompressionUploadAPIRequest {
	return poolAlitripPolicyRuleCompressionUploadAPIRequest.Get().(*AlitripPolicyRuleCompressionUploadAPIRequest)
}

// ReleaseAlitripPolicyRuleCompressionUploadAPIRequest 将 AlitripPolicyRuleCompressionUploadAPIRequest 放入 sync.Pool
func ReleaseAlitripPolicyRuleCompressionUploadAPIRequest(v *AlitripPolicyRuleCompressionUploadAPIRequest) {
	v.Reset()
	poolAlitripPolicyRuleCompressionUploadAPIRequest.Put(v)
}
