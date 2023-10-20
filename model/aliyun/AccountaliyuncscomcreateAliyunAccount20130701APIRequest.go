package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AccountaliyuncscomcreateAliyunAccount20130701APIRequest 创建阿里云账号 API请求
// account.aliyuncs.com.CreateAliyunAccount.2013-07-01
//
// 根据给定的阿里云账号，密码以及手机号创建阿里云账号
type AccountaliyuncscomcreateAliyunAccount20130701APIRequest struct {
	model.Params
}

// NewAccountaliyuncscomcreateAliyunAccount20130701Request 初始化AccountaliyuncscomcreateAliyunAccount20130701APIRequest对象
func NewAccountaliyuncscomcreateAliyunAccount20130701Request() *AccountaliyuncscomcreateAliyunAccount20130701APIRequest {
	return &AccountaliyuncscomcreateAliyunAccount20130701APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountaliyuncscomcreateAliyunAccount20130701APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.CreateAliyunAccount.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountaliyuncscomcreateAliyunAccount20130701APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AccountaliyuncscomcreateAliyunAccount20130701APIRequest) GetRawParams() model.Params {
	return r.Params
}
