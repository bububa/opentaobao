package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcrmmemberfrontunbindAPIRequest 品牌会员解绑 API请求
// tmall.crm.member.front.unbind
//
// 品牌会员解绑功能
type TmallcrmmemberfrontunbindAPIRequest struct {
	model.Params
	// 会员昵称
	_userNick string
}

// NewTmallcrmmemberfrontunbindRequest 初始化TmallcrmmemberfrontunbindAPIRequest对象
func NewTmallcrmmemberfrontunbindRequest() *TmallcrmmemberfrontunbindAPIRequest {
	return &TmallcrmmemberfrontunbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcrmmemberfrontunbindAPIRequest) GetApiMethodName() string {
	return "tmall.crm.member.front.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcrmmemberfrontunbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcrmmemberfrontunbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 会员昵称
func (r *TmallcrmmemberfrontunbindAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TmallcrmmemberfrontunbindAPIRequest) GetUserNick() string {
	return r._userNick
}
