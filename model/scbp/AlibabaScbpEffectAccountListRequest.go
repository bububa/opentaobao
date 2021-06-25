package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
账户-报表 APIRequest
alibaba.scbp.effect.account.list

账户-报表,支持最近7天，最近30天，以及180天内时间区间。
*/
type AlibabaScbpEffectAccountListRequest struct {
    model.Params

    // AccountQuery
    p4pAccountReportQuery   *AccountQuery 

}

func NewAlibabaScbpEffectAccountListRequest() *AlibabaScbpEffectAccountListRequest{
    return &AlibabaScbpEffectAccountListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpEffectAccountListRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.account.list"
}

func (r AlibabaScbpEffectAccountListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpEffectAccountListRequest) SetP4pAccountReportQuery(p4pAccountReportQuery *AccountQuery) error {
    r.p4pAccountReportQuery = p4pAccountReportQuery
    r.Set("p4p_account_report_query", p4pAccountReportQuery)
    return nil
}

func (r AlibabaScbpEffectAccountListRequest) GetP4pAccountReportQuery() *AccountQuery {
    return r.p4pAccountReportQuery
}

