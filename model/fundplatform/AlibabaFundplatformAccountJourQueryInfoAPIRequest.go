package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformaccountjourqueryinfoAPIRequest 查询账户流水信息 API请求
// alibaba.fundplatform.account.jour.query.info
//
// 外部查询账户流水信息
type AlibabafundplatformaccountjourqueryinfoAPIRequest struct {
	model.Params
	// 入参对象
	_paramFundAccountJournalQueryReq *FundAccountJournalQueryReq
}

// NewAlibabafundplatformaccountjourqueryinfoRequest 初始化AlibabafundplatformaccountjourqueryinfoAPIRequest对象
func NewAlibabafundplatformaccountjourqueryinfoRequest() *AlibabafundplatformaccountjourqueryinfoAPIRequest {
	return &AlibabafundplatformaccountjourqueryinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformaccountjourqueryinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.account.jour.query.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformaccountjourqueryinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformaccountjourqueryinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamFundAccountJournalQueryReq is ParamFundAccountJournalQueryReq Setter
// 入参对象
func (r *AlibabafundplatformaccountjourqueryinfoAPIRequest) SetParamFundAccountJournalQueryReq(_paramFundAccountJournalQueryReq *FundAccountJournalQueryReq) error {
	r._paramFundAccountJournalQueryReq = _paramFundAccountJournalQueryReq
	r.Set("param_fund_account_journal_query_req", _paramFundAccountJournalQueryReq)
	return nil
}

// GetParamFundAccountJournalQueryReq ParamFundAccountJournalQueryReq Getter
func (r AlibabafundplatformaccountjourqueryinfoAPIRequest) GetParamFundAccountJournalQueryReq() *FundAccountJournalQueryReq {
	return r._paramFundAccountJournalQueryReq
}
