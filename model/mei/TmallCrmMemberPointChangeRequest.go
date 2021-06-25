package mei

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员积分变更 APIRequest
tmall.crm.member.point.change

品牌CRM项目中，会员积分变更接口。
*/
type TmallCrmMemberPointChangeRequest struct {
    model.Params

    // 更改积分数值
    point   int64 

    // minus:扣减;add:累加
    type   string 

    // 业务代码
    bizCode   string 

    // 业务描述
    bizDetail   string 

    // 用户昵称
    userNick   string 

}

func NewTmallCrmMemberPointChangeRequest() *TmallCrmMemberPointChangeRequest{
    return &TmallCrmMemberPointChangeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCrmMemberPointChangeRequest) GetApiMethodName() string {
    return "tmall.crm.member.point.change"
}

func (r TmallCrmMemberPointChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCrmMemberPointChangeRequest) SetPoint(point int64) error {
    r.point = point
    r.Set("point", point)
    return nil
}

func (r TmallCrmMemberPointChangeRequest) GetPoint() int64 {
    return r.point
}

func (r *TmallCrmMemberPointChangeRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TmallCrmMemberPointChangeRequest) GetType() string {
    return r.type
}

func (r *TmallCrmMemberPointChangeRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r TmallCrmMemberPointChangeRequest) GetBizCode() string {
    return r.bizCode
}

func (r *TmallCrmMemberPointChangeRequest) SetBizDetail(bizDetail string) error {
    r.bizDetail = bizDetail
    r.Set("biz_detail", bizDetail)
    return nil
}

func (r TmallCrmMemberPointChangeRequest) GetBizDetail() string {
    return r.bizDetail
}

func (r *TmallCrmMemberPointChangeRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TmallCrmMemberPointChangeRequest) GetUserNick() string {
    return r.userNick
}

