package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账户流水信息 APIRequest
alibaba.fundplatform.account.jour.query.info

外部查询账户流水信息
*/
type AlibabaFundplatformAccountJourQueryInfoRequest struct {
    model.Params

    // 入参对象
    paramFundAccountJournalQueryReq   *FundAccountJournalQueryReq 

}

func NewAlibabaFundplatformAccountJourQueryInfoRequest() *AlibabaFundplatformAccountJourQueryInfoRequest{
    return &AlibabaFundplatformAccountJourQueryInfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformAccountJourQueryInfoRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.account.jour.query.info"
}

func (r AlibabaFundplatformAccountJourQueryInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformAccountJourQueryInfoRequest) SetParamFundAccountJournalQueryReq(paramFundAccountJournalQueryReq *FundAccountJournalQueryReq) error {
    r.paramFundAccountJournalQueryReq = paramFundAccountJournalQueryReq
    r.Set("param_fund_account_journal_query_req", paramFundAccountJournalQueryReq)
    return nil
}

func (r AlibabaFundplatformAccountJourQueryInfoRequest) GetParamFundAccountJournalQueryReq() *FundAccountJournalQueryReq {
    return r.paramFundAccountJournalQueryReq
}

