package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpEffectAccountListAPIRequest
账户-报表 API请求
alibaba.scbp.effect.account.list

账户-报表,支持最近7天，最近30天，以及180天内时间区间。 */
type AlibabaScbpEffectAccountListAPIRequest struct {
	model.Params
	// AccountQuery
	_p4pAccountReportQuery *AccountQuery
}

// New
