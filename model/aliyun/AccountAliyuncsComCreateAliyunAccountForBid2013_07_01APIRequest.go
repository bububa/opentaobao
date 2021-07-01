package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIRequest
为bid用户创建账号 API请求
account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01

给指定的bid创建账号，同时账号打上owner bid的标记 */
type AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIRequest struct {
	model.Params
}

// NewAccountAliyuncsComCreateAliyunAccountForBid2013_07_01Request 初始化AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIRequest对象
func NewAccountAliyuncsComCreateAliyunAccountForBid2013_07_01Request() *AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIRequest {
	return &AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIRequest) GetApiMethodName() string {
	return "account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
