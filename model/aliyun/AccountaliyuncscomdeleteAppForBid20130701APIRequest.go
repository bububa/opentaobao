package aliyun

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComDeleteAppForBid20130701APIRequest 运营商删除用户的appkey API请求
// account.aliyuncs.com.DeleteAppForBid.2013-07-01
//
// 删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。
type AccountAliyuncsComDeleteAppForBid20130701APIRequest struct {
	model.Params
	// 要删除的appkey的所有者用户的pk
	_ownerId string
	// 要删除的appkey
	_ownerAppkey string
}

// NewAccountAliyuncsComDeleteAppForBid20130701Request 初始化AccountAliyuncsComDeleteAppForBid20130701APIRequest对象
func NewAccountAliyuncsComDeleteAppForBid20130701Request() *AccountAliyuncsComDeleteAppForBid20130701APIRequest {
	return &AccountAliyuncsComDeleteAppForBid20130701APIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AccountAliyuncsComDeleteAppForBid20130701APIRequest) Reset() {
	r._ownerId = ""
	r._ownerAppkey = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComDeleteAppForBid20130701APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.DeleteAppForBid.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComDeleteAppForBid20130701APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AccountAliyuncsComDeleteAppForBid20130701APIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOwnerId is OwnerId Setter
// 要删除的appkey的所有者用户的pk
func (r *AccountAliyuncsComDeleteAppForBid20130701APIRequest) SetOwnerId(_ownerId string) error {
	r._ownerId = _ownerId
	r.Set("OwnerId", _ownerId)
	return nil
}

// GetOwnerId OwnerId Getter
func (r AccountAliyuncsComDeleteAppForBid20130701APIRequest) GetOwnerId() string {
	return r._ownerId
}

// SetOwnerAppkey is OwnerAppkey Setter
// 要删除的appkey
func (r *AccountAliyuncsComDeleteAppForBid20130701APIRequest) SetOwnerAppkey(_ownerAppkey string) error {
	r._ownerAppkey = _ownerAppkey
	r.Set("OwnerAppkey", _ownerAppkey)
	return nil
}

// GetOwnerAppkey OwnerAppkey Getter
func (r AccountAliyuncsComDeleteAppForBid20130701APIRequest) GetOwnerAppkey() string {
	return r._ownerAppkey
}

var poolAccountAliyuncsComDeleteAppForBid20130701APIRequest = sync.Pool{
	New: func() any {
		return NewAccountAliyuncsComDeleteAppForBid20130701Request()
	},
}

// GetAccountAliyuncsComDeleteAppForBid20130701Request 从 sync.Pool 获取 AccountAliyuncsComDeleteAppForBid20130701APIRequest
func GetAccountAliyuncsComDeleteAppForBid20130701APIRequest() *AccountAliyuncsComDeleteAppForBid20130701APIRequest {
	return poolAccountAliyuncsComDeleteAppForBid20130701APIRequest.Get().(*AccountAliyuncsComDeleteAppForBid20130701APIRequest)
}

// ReleaseAccountAliyuncsComDeleteAppForBid20130701APIRequest 将 AccountAliyuncsComDeleteAppForBid20130701APIRequest 放入 sync.Pool
func ReleaseAccountAliyuncsComDeleteAppForBid20130701APIRequest(v *AccountAliyuncsComDeleteAppForBid20130701APIRequest) {
	v.Reset()
	poolAccountAliyuncsComDeleteAppForBid20130701APIRequest.Put(v)
}
