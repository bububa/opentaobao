package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamemberexitAPIRequest 退会 API请求
// alibaba.member.exit
//
// 商家会员解绑
type AlibabamemberexitAPIRequest struct {
	model.Params
	// 退会对象
	_exitMember *ExitMemberDto
}

// NewAlibabamemberexitRequest 初始化AlibabamemberexitAPIRequest对象
func NewAlibabamemberexitRequest() *AlibabamemberexitAPIRequest {
	return &AlibabamemberexitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamemberexitAPIRequest) GetApiMethodName() string {
	return "alibaba.member.exit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamemberexitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamemberexitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExitMember is ExitMember Setter
// 退会对象
func (r *AlibabamemberexitAPIRequest) SetExitMember(_exitMember *ExitMemberDto) error {
	r._exitMember = _exitMember
	r.Set("exit_member", _exitMember)
	return nil
}

// GetExitMember ExitMember Getter
func (r AlibabamemberexitAPIRequest) GetExitMember() *ExitMemberDto {
	return r._exitMember
}
