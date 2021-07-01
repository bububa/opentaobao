package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowWalletSignAPIRequest
流量平台用户签约情况查询 API请求
alibaba.aliqin.flow.wallet.sign

流量平台用户签约情况查询 */
type AlibabaAliqinFlowWalletSignAPIRequest struct {
	model.Params
	// 用户昵称
	_userNick string
}

// NewAlibabaAliqinFlowWalletSignRequest 初始化AlibabaAliqinFlowWalletSignAPIRequest对象
func NewAlibabaAliqinFlowWalletSignRequest() *AlibabaAliqinFlowWalletSignAPIRequest {
	return &AlibabaAliqinFlowWalletSignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletSignAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletSignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserNick Setter
// 用户昵称
func (r *AlibabaAliqinFlowWalletSignAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// Get UserNick Getter
func (r AlibabaAliqinFlowWalletSignAPIRequest) GetUserNick() string {
	return r._userNick
}
