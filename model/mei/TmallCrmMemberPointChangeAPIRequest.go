package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcrmmemberpointchangeAPIRequest 会员积分变更 API请求
// tmall.crm.member.point.change
//
// 品牌CRM项目中，会员积分变更接口。
type TmallcrmmemberpointchangeAPIRequest struct {
	model.Params
	// minus:扣减;add:累加
	_type string
	// 业务代码
	_bizCode string
	// 业务描述
	_bizDetail string
	// 用户昵称
	_userNick string
	// 更改积分数值
	_point int64
}

// NewTmallcrmmemberpointchangeRequest 初始化TmallcrmmemberpointchangeAPIRequest对象
func NewTmallcrmmemberpointchangeRequest() *TmallcrmmemberpointchangeAPIRequest {
	return &TmallcrmmemberpointchangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcrmmemberpointchangeAPIRequest) GetApiMethodName() string {
	return "tmall.crm.member.point.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcrmmemberpointchangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcrmmemberpointchangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// minus:扣减;add:累加
func (r *TmallcrmmemberpointchangeAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TmallcrmmemberpointchangeAPIRequest) GetType() string {
	return r._type
}

// SetBizCode is BizCode Setter
// 业务代码
func (r *TmallcrmmemberpointchangeAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallcrmmemberpointchangeAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetBizDetail is BizDetail Setter
// 业务描述
func (r *TmallcrmmemberpointchangeAPIRequest) SetBizDetail(_bizDetail string) error {
	r._bizDetail = _bizDetail
	r.Set("biz_detail", _bizDetail)
	return nil
}

// GetBizDetail BizDetail Getter
func (r TmallcrmmemberpointchangeAPIRequest) GetBizDetail() string {
	return r._bizDetail
}

// SetUserNick is UserNick Setter
// 用户昵称
func (r *TmallcrmmemberpointchangeAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TmallcrmmemberpointchangeAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetPoint is Point Setter
// 更改积分数值
func (r *TmallcrmmemberpointchangeAPIRequest) SetPoint(_point int64) error {
	r._point = _point
	r.Set("point", _point)
	return nil
}

// GetPoint Point Getter
func (r TmallcrmmemberpointchangeAPIRequest) GetPoint() int64 {
	return r._point
}
