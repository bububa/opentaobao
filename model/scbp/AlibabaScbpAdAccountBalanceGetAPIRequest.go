package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdAccountBalanceGetAPIRequest 查询账户余额 API请求
// alibaba.scbp.ad.account.balance.get
//
// 查询推广账户余额
type AlibabaScbpAdAccountBalanceGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpAdAccountBalanceGetRequest 初始化AlibabaScbpAdAccountBalanceGetAPIRequest对象
func NewAlibabaScbpAdAccountBalanceGetRequest() *AlibabaScbpAdAccountBalanceGetAPIRequest {
	return &AlibabaScbpAdAccountBalanceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdAccountBalanceGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.account.balance.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdAccountBalanceGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
