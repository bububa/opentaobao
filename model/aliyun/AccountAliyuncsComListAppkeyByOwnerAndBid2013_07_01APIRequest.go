package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIRequest
根据渠道id和状态列出appkey API请求
account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01

根据渠道id和状态列出appkey */
type AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIRequest struct {
	model.Params
}

// NewAccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01Request 初始化AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIRequest对象
func NewAccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01Request() *AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIRequest {
	return &AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
