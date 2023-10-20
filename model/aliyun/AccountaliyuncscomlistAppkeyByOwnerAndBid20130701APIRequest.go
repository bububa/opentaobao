package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIRequest 根据渠道id和状态列出appkey API请求
// account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01
//
// 根据渠道id和状态列出appkey
type AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIRequest struct {
	model.Params
}

// NewAccountAliyuncsComListAppkeyByOwnerAndBid20130701Request 初始化AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIRequest对象
func NewAccountAliyuncsComListAppkeyByOwnerAndBid20130701Request() *AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIRequest {
	return &AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AccountAliyuncsComListAppkeyByOwnerAndBid20130701APIRequest) GetRawParams() model.Params {
	return r.Params
}
