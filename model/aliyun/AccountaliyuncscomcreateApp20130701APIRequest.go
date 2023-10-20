package aliyun

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComCreateApp20130701APIRequest 给指定用户创建appkey API请求
// account.aliyuncs.com.CreateApp.2013-07-01
//
// 为某个用户创建appkey
type AccountAliyuncsComCreateApp20130701APIRequest struct {
	model.Params
}

// NewAccountAliyuncsComCreateApp20130701Request 初始化AccountAliyuncsComCreateApp20130701APIRequest对象
func NewAccountAliyuncsComCreateApp20130701Request() *AccountAliyuncsComCreateApp20130701APIRequest {
	return &AccountAliyuncsComCreateApp20130701APIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AccountAliyuncsComCreateApp20130701APIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComCreateApp20130701APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.CreateApp.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComCreateApp20130701APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AccountAliyuncsComCreateApp20130701APIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAccountAliyuncsComCreateApp20130701APIRequest = sync.Pool{
	New: func() any {
		return NewAccountAliyuncsComCreateApp20130701Request()
	},
}

// GetAccountAliyuncsComCreateApp20130701Request 从 sync.Pool 获取 AccountAliyuncsComCreateApp20130701APIRequest
func GetAccountAliyuncsComCreateApp20130701APIRequest() *AccountAliyuncsComCreateApp20130701APIRequest {
	return poolAccountAliyuncsComCreateApp20130701APIRequest.Get().(*AccountAliyuncsComCreateApp20130701APIRequest)
}

// ReleaseAccountAliyuncsComCreateApp20130701APIRequest 将 AccountAliyuncsComCreateApp20130701APIRequest 放入 sync.Pool
func ReleaseAccountAliyuncsComCreateApp20130701APIRequest(v *AccountAliyuncsComCreateApp20130701APIRequest) {
	v.Reset()
	poolAccountAliyuncsComCreateApp20130701APIRequest.Put(v)
}
