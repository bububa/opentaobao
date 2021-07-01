package mei

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员积分变更 API请求
tmall.crm.member.point.change

品牌CRM项目中，会员积分变更接口。
*/
type TmallCrmMemberPointChangeAPIRequest struct {
    model.Params
    // 更改积分数值
    _point   int64
    // minus:扣减;add:累加
    _type   string
    // 业务代码
    _bizCode   string
    // 业务描述
    _bizDetail   string
    // 用户昵称
    _userNick   string
}

// 初始化TmallCrmMemberPointChangeAPIRequest对象
func NewTmallCrmMemberPointChangeRequest() *TmallCrmMemberPointChangeAPIRequest{
    return &TmallCrmMemberPointChangeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCrmMemberPointChangeAPIRequest) GetApiMethodName() string {
    return "tmall.crm.member.point.change"
}

// IRequest interface 方法, 获取API参数
func (r TmallCrmMemberPointChangeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Point Setter
// 更改积分数值
func (r *TmallCrmMemberPointChangeAPIRequest) SetPoint(_point int64) error {
    r._point = _point
    r.Set("point", _point)
    return nil
}

// Point Getter
func (r TmallCrmMemberPointChangeAPIRequest) GetPoint() int64 {
    return r._point
}
// Type Setter
// minus:扣减;add:累加
func (r *TmallCrmMemberPointChangeAPIRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TmallCrmMemberPointChangeAPIRequest) GetType() string {
    return r._type
}
// BizCode Setter
// 业务代码
func (r *TmallCrmMemberPointChangeAPIRequest) SetBizCode(_bizCode string) error {
    r._bizCode = _bizCode
    r.Set("biz_code", _bizCode)
    return nil
}

// BizCode Getter
func (r TmallCrmMemberPointChangeAPIRequest) GetBizCode() string {
    return r._bizCode
}
// BizDetail Setter
// 业务描述
func (r *TmallCrmMemberPointChangeAPIRequest) SetBizDetail(_bizDetail string) error {
    r._bizDetail = _bizDetail
    r.Set("biz_detail", _bizDetail)
    return nil
}

// BizDetail Getter
func (r TmallCrmMemberPointChangeAPIRequest) GetBizDetail() string {
    return r._bizDetail
}
// UserNick Setter
// 用户昵称
func (r *TmallCrmMemberPointChangeAPIRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r TmallCrmMemberPointChangeAPIRequest) GetUserNick() string {
    return r._userNick
}
