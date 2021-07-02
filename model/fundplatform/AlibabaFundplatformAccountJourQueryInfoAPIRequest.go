package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformAccountJourQueryInfoAPIRequest 查询账户流水信息 API请求
// alibaba.fundplatform.account.jour.query.info
//
// 外部查询账户流水信息
type AlibabaFundplatformAccountJourQueryInfoAPIRequest struct {
	model.Params
	// 入参对象
	_paramFundAccountJournalQueryReq *FundAccountJournalQueryReq
}

// NewAlibabaFundplatformAccountJourQueryInfoRequest 初始化AlibabaFundplatformAccountJourQueryInfoAPIRequest对象
func NewAlibabaFundplatformAccountJourQueryInfoRequest() *AlibabaFundplatformAccountJourQueryInfoAPIRequest {
	return &AlibabaFundplatformAccountJourQueryInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformAccountJourQueryInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.account.jour.query.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformAccountJourQueryInfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamFundAccountJournalQueryReq is ParamFundAccountJournalQueryReq Setter
// 入参对象
func (r *AlibabaFundplatformAccountJourQueryInfoAPIRequest) SetParamFundAccountJournalQueryReq(_paramFundAccountJournalQueryReq *FundAccountJournalQueryReq) error {
	r._paramFundAccountJournalQueryReq = _paramFundAccountJournalQueryReq
	r.Set("param_fund_account_journal_query_req", _paramFundAccountJournalQueryReq)
	return nil
}

// GetParamFundAccountJournalQueryReq ParamFundAccountJournalQueryReq Getter
func (r AlibabaFundplatformAccountJourQueryInfoAPIRequest) GetParamFundAccountJournalQueryReq() *FundAccountJournalQueryReq {
	return r._paramFundAccountJournalQueryReq
}
