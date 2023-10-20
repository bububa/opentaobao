package alimember

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMemberExitAPIRequest) Reset() {
	r._exitMember = nil
	r.Params.ToZero()
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

var poolAlibabaMemberExitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMemberExitRequest()
	},
}

// GetAlibabaMemberExitRequest 从 sync.Pool 获取 AlibabaMemberExitAPIRequest
func GetAlibabaMemberExitAPIRequest() *AlibabaMemberExitAPIRequest {
	return poolAlibabaMemberExitAPIRequest.Get().(*AlibabaMemberExitAPIRequest)
}

// ReleaseAlibabaMemberExitAPIRequest 将 AlibabaMemberExitAPIRequest 放入 sync.Pool
func ReleaseAlibabaMemberExitAPIRequest(v *AlibabaMemberExitAPIRequest) {
	v.Reset()
	poolAlibabaMemberExitAPIRequest.Put(v)
}
