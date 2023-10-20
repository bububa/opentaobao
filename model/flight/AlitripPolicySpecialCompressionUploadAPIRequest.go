package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicyspecialcompressionuploadAPIRequest 大批量上传特殊类型的单程/往返政策 API请求
// alitrip.policy.special.compression.upload
//
// 大批量上传特殊类型的单程/往返政策
type AlitrippolicyspecialcompressionuploadAPIRequest struct {
	model.Params
	// alitrip.policy.special.upload接口的参数经过压缩后得到的字节数组，生成的文件
	_file *model.File
}

// NewAlitrippolicyspecialcompressionuploadRequest 初始化AlitrippolicyspecialcompressionuploadAPIRequest对象
func NewAlitrippolicyspecialcompressionuploadRequest() *AlitrippolicyspecialcompressionuploadAPIRequest {
	return &AlitrippolicyspecialcompressionuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitrippolicyspecialcompressionuploadAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.special.compression.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitrippolicyspecialcompressionuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitrippolicyspecialcompressionuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFile is File Setter
// alitrip.policy.special.upload接口的参数经过压缩后得到的字节数组，生成的文件
func (r *AlitrippolicyspecialcompressionuploadAPIRequest) SetFile(_file *model.File) error {
	r._file = _file
	r.Set("file", _file)
	return nil
}

// GetFile File Getter
func (r AlitrippolicyspecialcompressionuploadAPIRequest) GetFile() *model.File {
	return r._file
}
