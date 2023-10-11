package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtMiaolingThirdLoginAPIRequest 喵零第三方免登 API请求
// tmall.nrt.miaoling.third.login
//
// 喵零第三方免登
type TmallNrtMiaolingThirdLoginAPIRequest struct {
	model.Params
	// 参数对象
	_nrtEaLoginDto *NrtEaLoginDto
}

// NewTmallNrtMiaolingThirdLoginRequest 初始化TmallNrtMiaolingThirdLoginAPIRequest对象
func NewTmallNrtMiaolingThirdLoginRequest() *TmallNrtMiaolingThirdLoginAPIRequest {
	return &TmallNrtMiaolingThirdLoginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtMiaolingThirdLoginAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.miaoling.third.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtMiaolingThirdLoginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtMiaolingThirdLoginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtEaLoginDto is NrtEaLoginDto Setter
// 参数对象
func (r *TmallNrtMiaolingThirdLoginAPIRequest) SetNrtEaLoginDto(_nrtEaLoginDto *NrtEaLoginDto) error {
	r._nrtEaLoginDto = _nrtEaLoginDto
	r.Set("nrt_ea_login_dto", _nrtEaLoginDto)
	return nil
}

// GetNrtEaLoginDto NrtEaLoginDto Getter
func (r TmallNrtMiaolingThirdLoginAPIRequest) GetNrtEaLoginDto() *NrtEaLoginDto {
	return r._nrtEaLoginDto
}
