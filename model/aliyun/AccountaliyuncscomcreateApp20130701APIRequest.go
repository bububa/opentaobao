package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AccountaliyuncscomcreateApp20130701APIRequest 给指定用户创建appkey API请求
// account.aliyuncs.com.CreateApp.2013-07-01
//
// 为某个用户创建appkey
type AccountaliyuncscomcreateApp20130701APIRequest struct {
	model.Params
}

// NewAccountaliyuncscomcreateApp20130701Request 初始化AccountaliyuncscomcreateApp20130701APIRequest对象
func NewAccountaliyuncscomcreateApp20130701Request() *AccountaliyuncscomcreateApp20130701APIRequest {
	return &AccountaliyuncscomcreateApp20130701APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountaliyuncscomcreateApp20130701APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.CreateApp.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountaliyuncscomcreateApp20130701APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AccountaliyuncscomcreateApp20130701APIRequest) GetRawParams() model.Params {
	return r.Params
}
