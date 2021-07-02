package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformAccountQueryInfoAPIRequest 查询账户信息 API请求
// alibaba.fundplatform.account.query.info
//
// 外部查询资金平台用户账户信息
type AlibabaFundplatformAccountQueryInfoAPIRequest struct {
	model.Params
	// 账户ID
	_accountId int64
}

// NewAlibabaFundplatformAccountQueryInfoRequest 初始化AlibabaFundplatformAccountQueryInfoAPIRequest对象
func NewAlibabaFundplatformAccountQueryInfoRequest() *AlibabaFundplatformAccountQueryInfoAPIRequest {
	return &AlibabaFundplatformAccountQueryInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformAccountQueryInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.account.query.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformAccountQueryInfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AccountId Setter
// 账户ID
func (r *AlibabaFundplatformAccountQueryInfoAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// Get AccountId Getter
func (r AlibabaFundplatformAccountQueryInfoAPIRequest) GetAccountId() int64 {
	return r._accountId
}
