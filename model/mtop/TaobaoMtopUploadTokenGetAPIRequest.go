package mtop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMtopUploadTokenGetAPIRequest 获取文件上传授权 API请求
// taobao.mtop.upload.token.get
//
// 获取mtop文件上传授权
type TaobaoMtopUploadTokenGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramUploadTokenRequest *UploadTokenRequestV
}

// NewTaobaoMtopUploadTokenGetRequest 初始化TaobaoMtopUploadTokenGetAPIRequest对象
func NewTaobaoMtopUploadTokenGetRequest() *TaobaoMtopUploadTokenGetAPIRequest {
	return &TaobaoMtopUploadTokenGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMtopUploadTokenGetAPIRequest) Reset() {
	r._paramUploadTokenRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMtopUploadTokenGetAPIRequest) GetApiMethodName() string {
	return "taobao.mtop.upload.token.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMtopUploadTokenGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMtopUploadTokenGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUploadTokenRequest is ParamUploadTokenRequest Setter
// 系统自动生成
func (r *TaobaoMtopUploadTokenGetAPIRequest) SetParamUploadTokenRequest(_paramUploadTokenRequest *UploadTokenRequestV) error {
	r._paramUploadTokenRequest = _paramUploadTokenRequest
	r.Set("param_upload_token_request", _paramUploadTokenRequest)
	return nil
}

// GetParamUploadTokenRequest ParamUploadTokenRequest Getter
func (r TaobaoMtopUploadTokenGetAPIRequest) GetParamUploadTokenRequest() *UploadTokenRequestV {
	return r._paramUploadTokenRequest
}

var poolTaobaoMtopUploadTokenGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMtopUploadTokenGetRequest()
	},
}

// GetTaobaoMtopUploadTokenGetRequest 从 sync.Pool 获取 TaobaoMtopUploadTokenGetAPIRequest
func GetTaobaoMtopUploadTokenGetAPIRequest() *TaobaoMtopUploadTokenGetAPIRequest {
	return poolTaobaoMtopUploadTokenGetAPIRequest.Get().(*TaobaoMtopUploadTokenGetAPIRequest)
}

// ReleaseTaobaoMtopUploadTokenGetAPIRequest 将 TaobaoMtopUploadTokenGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoMtopUploadTokenGetAPIRequest(v *TaobaoMtopUploadTokenGetAPIRequest) {
	v.Reset()
	poolTaobaoMtopUploadTokenGetAPIRequest.Put(v)
}
