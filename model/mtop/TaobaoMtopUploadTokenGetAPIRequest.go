package mtop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomtopuploadtokengetAPIRequest 获取文件上传授权 API请求
// taobao.mtop.upload.token.get
//
// 获取mtop文件上传授权
type TaobaomtopuploadtokengetAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramUploadTokenRequest *UploadTokenRequestV
}

// NewTaobaomtopuploadtokengetRequest 初始化TaobaomtopuploadtokengetAPIRequest对象
func NewTaobaomtopuploadtokengetRequest() *TaobaomtopuploadtokengetAPIRequest {
	return &TaobaomtopuploadtokengetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomtopuploadtokengetAPIRequest) GetApiMethodName() string {
	return "taobao.mtop.upload.token.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomtopuploadtokengetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomtopuploadtokengetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUploadTokenRequest is ParamUploadTokenRequest Setter
// 系统自动生成
func (r *TaobaomtopuploadtokengetAPIRequest) SetParamUploadTokenRequest(_paramUploadTokenRequest *UploadTokenRequestV) error {
	r._paramUploadTokenRequest = _paramUploadTokenRequest
	r.Set("param_upload_token_request", _paramUploadTokenRequest)
	return nil
}

// GetParamUploadTokenRequest ParamUploadTokenRequest Getter
func (r TaobaomtopuploadtokengetAPIRequest) GetParamUploadTokenRequest() *UploadTokenRequestV {
	return r._paramUploadTokenRequest
}
