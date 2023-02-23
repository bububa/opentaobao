package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberExitAPIRequest 退会 API请求
// alibaba.member.exit
//
// 商家会员解绑
type AlibabaMemberExitAPIRequest struct {
	model.Params
	// 退会对象
	_exitMember *ExitMemberDto
}

// NewAlibabaMemberExitRequest 初始化AlibabaMemberExitAPIRequest对象
func NewAlibabaMemberExitRequest() *AlibabaMemberExitAPIRequest {
	return &AlibabaMemberExitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberExitAPIRequest) GetApiMethodName() string {
	return "alibaba.member.exit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberExitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMemberExitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExitMember is ExitMember Setter
// 退会对象
func (r *AlibabaMemberExitAPIRequest) SetExitMember(_exitMember *ExitMemberDto) error {
	r._exitMember = _exitMember
	r.Set("exit_member", _exitMember)
	return nil
}

// GetExitMember ExitMember Getter
func (r AlibabaMemberExitAPIRequest) GetExitMember() *ExitMemberDto {
	return r._exitMember
}
