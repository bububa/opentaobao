package mei

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCrmMemberFrontUnbindAPIRequest 品牌会员解绑 API请求
// tmall.crm.member.front.unbind
//
// 品牌会员解绑功能
type TmallCrmMemberFrontUnbindAPIRequest struct {
	model.Params
	// 会员昵称
	_userNick string
}

// NewTmallCrmMemberFrontUnbindRequest 初始化TmallCrmMemberFrontUnbindAPIRequest对象
func NewTmallCrmMemberFrontUnbindRequest() *TmallCrmMemberFrontUnbindAPIRequest {
	return &TmallCrmMemberFrontUnbindAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCrmMemberFrontUnbindAPIRequest) Reset() {
	r._userNick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCrmMemberFrontUnbindAPIRequest) GetApiMethodName() string {
	return "tmall.crm.member.front.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCrmMemberFrontUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCrmMemberFrontUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 会员昵称
func (r *TmallCrmMemberFrontUnbindAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TmallCrmMemberFrontUnbindAPIRequest) GetUserNick() string {
	return r._userNick
}

var poolTmallCrmMemberFrontUnbindAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCrmMemberFrontUnbindRequest()
	},
}

// GetTmallCrmMemberFrontUnbindRequest 从 sync.Pool 获取 TmallCrmMemberFrontUnbindAPIRequest
func GetTmallCrmMemberFrontUnbindAPIRequest() *TmallCrmMemberFrontUnbindAPIRequest {
	return poolTmallCrmMemberFrontUnbindAPIRequest.Get().(*TmallCrmMemberFrontUnbindAPIRequest)
}

// ReleaseTmallCrmMemberFrontUnbindAPIRequest 将 TmallCrmMemberFrontUnbindAPIRequest 放入 sync.Pool
func ReleaseTmallCrmMemberFrontUnbindAPIRequest(v *TmallCrmMemberFrontUnbindAPIRequest) {
	v.Reset()
	poolTmallCrmMemberFrontUnbindAPIRequest.Put(v)
}
