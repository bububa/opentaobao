package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServiceBillingQueryAPIRequest
服务平台结算出账信息 API请求
alibaba.service.billing.query

服务平台结算单明细查询服务 */
type AlibabaServiceBillingQueryAPIRequest struct {
	model.Params
	// 账单查询开始时间。格式示例 2019-03-26 17:15:28
	_gmtCreateStart string
	// 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
	_gmtCreateEnd string
}

// New
