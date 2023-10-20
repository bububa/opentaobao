package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformaccountqueryinfoAPIRequest 查询账户信息 API请求
// alibaba.fundplatform.account.query.info
//
// 外部查询资金平台用户账户信息
type AlibabafundplatformaccountqueryinfoAPIRequest struct {
	model.Params
	// 账户ID
	_accountId int64
}

// NewAlibabafundplatformaccountqueryinfoRequest 初始化AlibabafundplatformaccountqueryinfoAPIRequest对象
func NewAlibabafundplatformaccountqueryinfoRequest() *AlibabafundplatformaccountqueryinfoAPIRequest {
	return &AlibabafundplatformaccountqueryinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformaccountqueryinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.account.query.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformaccountqueryinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformaccountqueryinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// 账户ID
func (r *AlibabafundplatformaccountqueryinfoAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabafundplatformaccountqueryinfoAPIRequest) GetAccountId() int64 {
	return r._accountId
}
