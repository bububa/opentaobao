package mei

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌会员解绑 APIRequest
tmall.crm.member.front.unbind

品牌会员解绑功能
*/
type TmallCrmMemberFrontUnbindRequest struct {
    model.Params

    // 会员昵称
    userNick   string 

}

func NewTmallCrmMemberFrontUnbindRequest() *TmallCrmMemberFrontUnbindRequest{
    return &TmallCrmMemberFrontUnbindRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCrmMemberFrontUnbindRequest) GetApiMethodName() string {
    return "tmall.crm.member.front.unbind"
}

func (r TmallCrmMemberFrontUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCrmMemberFrontUnbindRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TmallCrmMemberFrontUnbindRequest) GetUserNick() string {
    return r.userNick
}

