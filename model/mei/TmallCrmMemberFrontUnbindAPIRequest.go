package mei

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌会员解绑 API请求
tmall.crm.member.front.unbind

品牌会员解绑功能
*/
type TmallCrmMemberFrontUnbindAPIRequest struct {
    model.Params
    // 会员昵称
    _userNick   string
}

// 初始化TmallCrmMemberFrontUnbindAPIRequest对象
func NewTmallCrmMemberFrontUnbindRequest() *TmallCrmMemberFrontUnbindAPIRequest{
    return &TmallCrmMemberFrontUnbindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCrmMemberFrontUnbindAPIRequest) GetApiMethodName() string {
    return "tmall.crm.member.front.unbind"
}

// IRequest interface 方法, 获取API参数
func (r TmallCrmMemberFrontUnbindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserNick Setter
// 会员昵称
func (r *TmallCrmMemberFrontUnbindAPIRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r TmallCrmMemberFrontUnbindAPIRequest) GetUserNick() string {
    return r._userNick
}
