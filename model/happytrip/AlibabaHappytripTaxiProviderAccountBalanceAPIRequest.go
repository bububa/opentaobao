package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiProviderAccountBalanceAPIRequest
供应商渠道余额 API请求
alibaba.happytrip.taxi.provider.account.balance

查询不同供应商不同渠道账户余额 */
type AlibabaHappytripTaxiProviderAccountBalanceAPIRequest struct {
	model.Params
	// 成本中心代码，用于区分不同的分账账号
	_costCenter string
}

// New
