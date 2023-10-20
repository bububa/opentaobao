package aliyun

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest 为bid用户创建账号 API请求
// account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01
//
// 给指定的bid创建账号，同时账号打上owner bid的标记
type AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest struct {
	model.Params
}

// NewAccountAliyuncsComCreateAliyunAccountForBid20130701Request 初始化AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest对象
func NewAccountAliyuncsComCreateAliyunAccountForBid20130701Request() *AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest {
	return &AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest = sync.Pool{
	New: func() any {
		return NewAccountAliyuncsComCreateAliyunAccountForBid20130701Request()
	},
}

// GetAccountAliyuncsComCreateAliyunAccountForBid20130701Request 从 sync.Pool 获取 AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest
func GetAccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest() *AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest {
	return poolAccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest.Get().(*AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest)
}

// ReleaseAccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest 将 AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest 放入 sync.Pool
func ReleaseAccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest(v *AccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest) {
	v.Reset()
	poolAccountAliyuncsComCreateAliyunAccountForBid20130701APIRequest.Put(v)
}
