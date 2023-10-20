package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AccountaliyuncscomcreateAliyunAccountForBid20130701APIRequest 为bid用户创建账号 API请求
// account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01
//
// 给指定的bid创建账号，同时账号打上owner bid的标记
type AccountaliyuncscomcreateAliyunAccountForBid20130701APIRequest struct {
	model.Params
}

// NewAccountaliyuncscomcreateAliyunAccountForBid20130701Request 初始化AccountaliyuncscomcreateAliyunAccountForBid20130701APIRequest对象
func NewAccountaliyuncscomcreateAliyunAccountForBid20130701Request() *AccountaliyuncscomcreateAliyunAccountForBid20130701APIRequest {
	return &AccountaliyuncscomcreateAliyunAccountForBid20130701APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountaliyuncscomcreateAliyunAccountForBid20130701APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountaliyuncscomcreateAliyunAccountForBid20130701APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AccountaliyuncscomcreateAliyunAccountForBid20130701APIRequest) GetRawParams() model.Params {
	return r.Params
}
