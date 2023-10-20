package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaVideoTokenGetAPIRequest 获取上传token API请求
// alibaba.video.token.get
//
// 获取上传token
type AlibabaVideoTokenGetAPIRequest struct {
	model.Params
	// 上传文件类型限制，可空，有默认文件类型
	_mimeLimit string
	// 上传大小限制(单位：字节)，最大不超过2G
	_sizeLimit int64
	// token超时时间，毫秒，比如1ms超时，参数就填1
	_expiration int64
}

// NewAlibabaVideoTokenGetRequest 初始化AlibabaVideoTokenGetAPIRequest对象
func NewAlibabaVideoTokenGetRequest() *AlibabaVideoTokenGetAPIRequest {
	return &AlibabaVideoTokenGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaVideoTokenGetAPIRequest) GetApiMethodName() string {
	return "alibaba.video.token.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaVideoTokenGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaVideoTokenGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMimeLimit is MimeLimit Setter
// 上传文件类型限制，可空，有默认文件类型
func (r *AlibabaVideoTokenGetAPIRequest) SetMimeLimit(_mimeLimit string) error {
	r._mimeLimit = _mimeLimit
	r.Set("mime_limit", _mimeLimit)
	return nil
}

// GetMimeLimit MimeLimit Getter
func (r AlibabaVideoTokenGetAPIRequest) GetMimeLimit() string {
	return r._mimeLimit
}

// SetSizeLimit is SizeLimit Setter
// 上传大小限制(单位：字节)，最大不超过2G
func (r *AlibabaVideoTokenGetAPIRequest) SetSizeLimit(_sizeLimit int64) error {
	r._sizeLimit = _sizeLimit
	r.Set("size_limit", _sizeLimit)
	return nil
}

// GetSizeLimit SizeLimit Getter
func (r AlibabaVideoTokenGetAPIRequest) GetSizeLimit() int64 {
	return r._sizeLimit
}

// SetExpiration is Expiration Setter
// token超时时间，毫秒，比如1ms超时，参数就填1
func (r *AlibabaVideoTokenGetAPIRequest) SetExpiration(_expiration int64) error {
	r._expiration = _expiration
	r.Set("expiration", _expiration)
	return nil
}

// GetExpiration Expiration Getter
func (r AlibabaVideoTokenGetAPIRequest) GetExpiration() int64 {
	return r._expiration
}
