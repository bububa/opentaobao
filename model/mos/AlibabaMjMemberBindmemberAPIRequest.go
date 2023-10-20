package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMemberBindmemberAPIRequest 绑定会员 API请求
// alibaba.mj.member.bindmember
//
// 用于绑定喵街数字化会员
type AlibabaMjMemberBindmemberAPIRequest struct {
	model.Params
	// 渠道
	_channel string
	// open_id
	_openId string
	// 用户号
	_userId int64
	// 商城Id
	_mallId int64
}

// NewAlibabaMjMemberBindmemberRequest 初始化AlibabaMjMemberBindmemberAPIRequest对象
func NewAlibabaMjMemberBindmemberRequest() *AlibabaMjMemberBindmemberAPIRequest {
	return &AlibabaMjMemberBindmemberAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjMemberBindmemberAPIRequest) Reset() {
	r._channel = ""
	r._openId = ""
	r._userId = 0
	r._mallId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMemberBindmemberAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.member.bindmember"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMemberBindmemberAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjMemberBindmemberAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannel is Channel Setter
// 渠道
func (r *AlibabaMjMemberBindmemberAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r AlibabaMjMemberBindmemberAPIRequest) GetChannel() string {
	return r._channel
}

// SetOpenId is OpenId Setter
// open_id
func (r *AlibabaMjMemberBindmemberAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r AlibabaMjMemberBindmemberAPIRequest) GetOpenId() string {
	return r._openId
}

// SetUserId is UserId Setter
// 用户号
func (r *AlibabaMjMemberBindmemberAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaMjMemberBindmemberAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetMallId is MallId Setter
// 商城Id
func (r *AlibabaMjMemberBindmemberAPIRequest) SetMallId(_mallId int64) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// GetMallId MallId Getter
func (r AlibabaMjMemberBindmemberAPIRequest) GetMallId() int64 {
	return r._mallId
}

var poolAlibabaMjMemberBindmemberAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjMemberBindmemberRequest()
	},
}

// GetAlibabaMjMemberBindmemberRequest 从 sync.Pool 获取 AlibabaMjMemberBindmemberAPIRequest
func GetAlibabaMjMemberBindmemberAPIRequest() *AlibabaMjMemberBindmemberAPIRequest {
	return poolAlibabaMjMemberBindmemberAPIRequest.Get().(*AlibabaMjMemberBindmemberAPIRequest)
}

// ReleaseAlibabaMjMemberBindmemberAPIRequest 将 AlibabaMjMemberBindmemberAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjMemberBindmemberAPIRequest(v *AlibabaMjMemberBindmemberAPIRequest) {
	v.Reset()
	poolAlibabaMjMemberBindmemberAPIRequest.Put(v)
}
