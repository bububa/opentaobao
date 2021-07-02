package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest 运营商删除用户的appkey API请求
// account.aliyuncs.com.DeleteAppForBid.2013-07-01
//
// 删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。
type AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest struct {
	model.Params
	// 要删除的appkey的所有者用户的pk
	_ownerId string
	// 要删除的appkey
	_ownerAppkey string
}

// NewAccountAliyuncsComDeleteAppForBid2013_07_01Request 初始化AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest对象
func NewAccountAliyuncsComDeleteAppForBid2013_07_01Request() *AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest {
	return &AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.DeleteAppForBid.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OwnerId Setter
// 要删除的appkey的所有者用户的pk
func (r *AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest) SetOwnerId(_ownerId string) error {
	r._ownerId = _ownerId
	r.Set("OwnerId", _ownerId)
	return nil
}

// Get OwnerId Getter
func (r AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest) GetOwnerId() string {
	return r._ownerId
}

// Set is OwnerAppkey Setter
// 要删除的appkey
func (r *AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest) SetOwnerAppkey(_ownerAppkey string) error {
	r._ownerAppkey = _ownerAppkey
	r.Set("OwnerAppkey", _ownerAppkey)
	return nil
}

// Get OwnerAppkey Getter
func (r AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest) GetOwnerAppkey() string {
	return r._ownerAppkey
}
