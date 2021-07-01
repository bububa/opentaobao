package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformAccountJourQueryInfoAPIRequest
查询账户流水信息 API请求
alibaba.fundplatform.account.jour.query.info

外部查询账户流水信息 */
type AlibabaFundplatformAccountJourQueryInfoAPIRequest struct {
	model.Params
	// 入参对象
	_paramFundAccountJournalQueryReq *FundAccountJournalQueryReq
}

// New
