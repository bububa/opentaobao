package aliyun

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComCreateAppForBid20130701APIRequest 给当前渠道下的用户创建app API请求
// account.aliyuncs.com.CreateAppForBid.2013-07-01
//
// 给自己渠道下的用户创建app
type AccountAliyuncsComCreateAppForBid20130701APIRequest struct {
	model.Params
}

// NewAccountAliyuncsComCreateAppForBid20130701Request 初始化AccountAliyuncsComCreateAppForBid20130701APIRequest对象
func NewAccountAliyuncsComCreateAppForBid20130701Request() *AccountAliyuncsComCreateAppForBid20130701APIRequest {
	return &AccountAliyuncsComCreateAppForBid20130701APIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AccountAliyuncsComCreateAppForBid20130701APIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComCreateAppForBid20130701APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.CreateAppForBid.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComCreateAppForBid20130701APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AccountAliyuncsComCreateAppForBid20130701APIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAccountAliyuncsComCreateAppForBid20130701APIRequest = sync.Pool{
	New: func() any {
		return NewAccountAliyuncsComCreateAppForBid20130701Request()
	},
}

// GetAccountAliyuncsComCreateAppForBid20130701Request 从 sync.Pool 获取 AccountAliyuncsComCreateAppForBid20130701APIRequest
func GetAccountAliyuncsComCreateAppForBid20130701APIRequest() *AccountAliyuncsComCreateAppForBid20130701APIRequest {
	return poolAccountAliyuncsComCreateAppForBid20130701APIRequest.Get().(*AccountAliyuncsComCreateAppForBid20130701APIRequest)
}

// ReleaseAccountAliyuncsComCreateAppForBid20130701APIRequest 将 AccountAliyuncsComCreateAppForBid20130701APIRequest 放入 sync.Pool
func ReleaseAccountAliyuncsComCreateAppForBid20130701APIRequest(v *AccountAliyuncsComCreateAppForBid20130701APIRequest) {
	v.Reset()
	poolAccountAliyuncsComCreateAppForBid20130701APIRequest.Put(v)
}
