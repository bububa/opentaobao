package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCrmMemberPointChangeAPIRequest 会员积分变更 API请求
// tmall.crm.member.point.change
//
// 品牌CRM项目中，会员积分变更接口。
type TmallCrmMemberPointChangeAPIRequest struct {
	model.Params
	// 更改积分数值
	_point int64
	// minus:扣减;add:累加
	_type string
	// 业务代码
	_bizCode string
	// 业务描述
	_bizDetail string
	// 用户昵称
	_userNick string
}

// NewTmallCrmMemberPointChangeRequest 初始化TmallCrmMemberPointChangeAPIRequest对象
func NewTmallCrmMemberPointChangeRequest() *TmallCrmMemberPointChangeAPIRequest {
	return &TmallCrmMemberPointChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCrmMemberPointChangeAPIRequest) GetApiMethodName() string {
	return "tmall.crm.member.point.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCrmMemberPointChangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Point Setter
// 更改积分数值
func (r *TmallCrmMemberPointChangeAPIRequest) SetPoint(_point int64) error {
	r._point = _point
	r.Set("point", _point)
	return nil
}

// Get Point Getter
func (r TmallCrmMemberPointChangeAPIRequest) GetPoint() int64 {
	return r._point
}

// Set is Type Setter
// minus:扣减;add:累加
func (r *TmallCrmMemberPointChangeAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TmallCrmMemberPointChangeAPIRequest) GetType() string {
	return r._type
}

// Set is BizCode Setter
// 业务代码
func (r *TmallCrmMemberPointChangeAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// Get BizCode Getter
func (r TmallCrmMemberPointChangeAPIRequest) GetBizCode() string {
	return r._bizCode
}

// Set is BizDetail Setter
// 业务描述
func (r *TmallCrmMemberPointChangeAPIRequest) SetBizDetail(_bizDetail string) error {
	r._bizDetail = _bizDetail
	r.Set("biz_detail", _bizDetail)
	return nil
}

// Get BizDetail Getter
func (r TmallCrmMemberPointChangeAPIRequest) GetBizDetail() string {
	return r._bizDetail
}

// Set is UserNick Setter
// 用户昵称
func (r *TmallCrmMemberPointChangeAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// Get UserNick Getter
func (r TmallCrmMemberPointChangeAPIRequest) GetUserNick() string {
	return r._userNick
}
