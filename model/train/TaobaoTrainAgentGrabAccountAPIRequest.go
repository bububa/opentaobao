package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentGrabAccountAPIRequest 代购抢代理商回传12306账号 API请求
// taobao.train.agent.grab.account
//
// 火车票业务代购抢功能，代理商回传12306账号，用于自营抢票链路出票
type TaobaoTrainAgentGrabAccountAPIRequest struct {
	model.Params
	// 12306账户信息
	_accountParam *AccountParam
}

// NewTaobaoTrainAgentGrabAccountRequest 初始化TaobaoTrainAgentGrabAccountAPIRequest对象
func NewTaobaoTrainAgentGrabAccountRequest() *TaobaoTrainAgentGrabAccountAPIRequest {
	return &TaobaoTrainAgentGrabAccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentGrabAccountAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.grab.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentGrabAccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentGrabAccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountParam is AccountParam Setter
// 12306账户信息
func (r *TaobaoTrainAgentGrabAccountAPIRequest) SetAccountParam(_accountParam *AccountParam) error {
	r._accountParam = _accountParam
	r.Set("account_param", _accountParam)
	return nil
}

// GetAccountParam AccountParam Getter
func (r TaobaoTrainAgentGrabAccountAPIRequest) GetAccountParam() *AccountParam {
	return r._accountParam
}
