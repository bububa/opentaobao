package nrt

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtMiaolingThirdLoginAPIRequest) Reset() {
	r._nrtEaLoginDto = nil
	r.Params.ToZero()
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

var poolTmallNrtMiaolingThirdLoginAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtMiaolingThirdLoginRequest()
	},
}

// GetTmallNrtMiaolingThirdLoginRequest 从 sync.Pool 获取 TmallNrtMiaolingThirdLoginAPIRequest
func GetTmallNrtMiaolingThirdLoginAPIRequest() *TmallNrtMiaolingThirdLoginAPIRequest {
	return poolTmallNrtMiaolingThirdLoginAPIRequest.Get().(*TmallNrtMiaolingThirdLoginAPIRequest)
}

// ReleaseTmallNrtMiaolingThirdLoginAPIRequest 将 TmallNrtMiaolingThirdLoginAPIRequest 放入 sync.Pool
func ReleaseTmallNrtMiaolingThirdLoginAPIRequest(v *TmallNrtMiaolingThirdLoginAPIRequest) {
	v.Reset()
	poolTmallNrtMiaolingThirdLoginAPIRequest.Put(v)
}
