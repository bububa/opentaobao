package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjMemberHasbindAPIRequest
喵街会员是否绑定 API请求
alibaba.mj.member.hasbind

喵街检测用户是否为数字化会员 */
type AlibabaMjMemberHasbindAPIRequest struct {
	model.Params
	// user_id
	_userId int64
	// open_id
	_openId string
}

// NewAlibabaMjMemberHasbindRequest 初始化AlibabaMjMemberHasbindAPIRequest对象
func NewAlibabaMjMemberHasbindRequest() *AlibabaMjMemberHasbindAPIRequest {
	return &AlibabaMjMemberHasbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMemberHasbindAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.member.hasbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMemberHasbindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserId Setter
// user_id
func (r *AlibabaMjMemberHasbindAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaMjMemberHasbindAPIRequest) GetUserId() int64 {
	return r._userId
}

// Set is OpenId Setter
// open_id
func (r *AlibabaMjMemberHasbindAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// Get OpenId Getter
func (r AlibabaMjMemberHasbindAPIRequest) GetOpenId() string {
	return r._openId
}
