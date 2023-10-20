package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjmemberhasbindAPIRequest 喵街会员是否绑定 API请求
// alibaba.mj.member.hasbind
//
// 喵街检测用户是否为数字化会员
type AlibabamjmemberhasbindAPIRequest struct {
	model.Params
	// open_id
	_openId string
	// user_id
	_userId int64
}

// NewAlibabamjmemberhasbindRequest 初始化AlibabamjmemberhasbindAPIRequest对象
func NewAlibabamjmemberhasbindRequest() *AlibabamjmemberhasbindAPIRequest {
	return &AlibabamjmemberhasbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjmemberhasbindAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.member.hasbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjmemberhasbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjmemberhasbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenId is OpenId Setter
// open_id
func (r *AlibabamjmemberhasbindAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r AlibabamjmemberhasbindAPIRequest) GetOpenId() string {
	return r._openId
}

// SetUserId is UserId Setter
// user_id
func (r *AlibabamjmemberhasbindAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabamjmemberhasbindAPIRequest) GetUserId() int64 {
	return r._userId
}
