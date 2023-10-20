package aliyun

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComCreateAliyunAccount20130701APIRequest 创建阿里云账号 API请求
// account.aliyuncs.com.CreateAliyunAccount.2013-07-01
//
// 根据给定的阿里云账号，密码以及手机号创建阿里云账号
type AccountAliyuncsComCreateAliyunAccount20130701APIRequest struct {
	model.Params
}

// NewAccountAliyuncsComCreateAliyunAccount20130701Request 初始化AccountAliyuncsComCreateAliyunAccount20130701APIRequest对象
func NewAccountAliyuncsComCreateAliyunAccount20130701Request() *AccountAliyuncsComCreateAliyunAccount20130701APIRequest {
	return &AccountAliyuncsComCreateAliyunAccount20130701APIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AccountAliyuncsComCreateAliyunAccount20130701APIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComCreateAliyunAccount20130701APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.CreateAliyunAccount.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComCreateAliyunAccount20130701APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AccountAliyuncsComCreateAliyunAccount20130701APIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAccountAliyuncsComCreateAliyunAccount20130701APIRequest = sync.Pool{
	New: func() any {
		return NewAccountAliyuncsComCreateAliyunAccount20130701Request()
	},
}

// GetAccountAliyuncsComCreateAliyunAccount20130701Request 从 sync.Pool 获取 AccountAliyuncsComCreateAliyunAccount20130701APIRequest
func GetAccountAliyuncsComCreateAliyunAccount20130701APIRequest() *AccountAliyuncsComCreateAliyunAccount20130701APIRequest {
	return poolAccountAliyuncsComCreateAliyunAccount20130701APIRequest.Get().(*AccountAliyuncsComCreateAliyunAccount20130701APIRequest)
}

// ReleaseAccountAliyuncsComCreateAliyunAccount20130701APIRequest 将 AccountAliyuncsComCreateAliyunAccount20130701APIRequest 放入 sync.Pool
func ReleaseAccountAliyuncsComCreateAliyunAccount20130701APIRequest(v *AccountAliyuncsComCreateAliyunAccount20130701APIRequest) {
	v.Reset()
	poolAccountAliyuncsComCreateAliyunAccount20130701APIRequest.Put(v)
}
