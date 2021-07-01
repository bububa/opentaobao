package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformAccountQueryInfoAPIRequest
查询账户信息 API请求
alibaba.fundplatform.account.query.info

外部查询资金平台用户账户信息 */
type AlibabaFundplatformAccountQueryInfoAPIRequest struct {
	model.Params
	// 账户ID
	_accountId int64
}

// New
