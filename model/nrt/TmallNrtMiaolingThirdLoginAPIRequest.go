package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtmiaolingthirdloginAPIRequest 喵零第三方免登 API请求
// tmall.nrt.miaoling.third.login
//
// 喵零第三方免登
type TmallnrtmiaolingthirdloginAPIRequest struct {
	model.Params
	// 参数对象
	_nrtEaLoginDto *NrtEaLoginDto
}

// NewTmallnrtmiaolingthirdloginRequest 初始化TmallnrtmiaolingthirdloginAPIRequest对象
func NewTmallnrtmiaolingthirdloginRequest() *TmallnrtmiaolingthirdloginAPIRequest {
	return &TmallnrtmiaolingthirdloginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtmiaolingthirdloginAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.miaoling.third.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtmiaolingthirdloginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtmiaolingthirdloginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtEaLoginDto is NrtEaLoginDto Setter
// 参数对象
func (r *TmallnrtmiaolingthirdloginAPIRequest) SetNrtEaLoginDto(_nrtEaLoginDto *NrtEaLoginDto) error {
	r._nrtEaLoginDto = _nrtEaLoginDto
	r.Set("nrt_ea_login_dto", _nrtEaLoginDto)
	return nil
}

// GetNrtEaLoginDto NrtEaLoginDto Getter
func (r TmallnrtmiaolingthirdloginAPIRequest) GetNrtEaLoginDto() *NrtEaLoginDto {
	return r._nrtEaLoginDto
}
