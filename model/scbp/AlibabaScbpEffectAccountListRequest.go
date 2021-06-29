package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
账户-报表 API请求
alibaba.scbp.effect.account.list

账户-报表,支持最近7天，最近30天，以及180天内时间区间。
*/
type AlibabaScbpEffectAccountListRequest struct {
    model.Params
    // AccountQuery
    p4pAccountReportQuery   *AccountQuery
}

// 初始化AlibabaScbpEffectAccountListRequest对象
func NewAlibabaScbpEffectAccountListRequest() *AlibabaScbpEffectAccountListRequest{
    return &AlibabaScbpEffectAccountListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectAccountListRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.account.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectAccountListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// P4pAccountReportQuery Setter
// AccountQuery
func (r *AlibabaScbpEffectAccountListRequest) SetP4pAccountReportQuery(p4pAccountReportQuery *AccountQuery) error {
    r.p4pAccountReportQuery = p4pAccountReportQuery
    r.Set("p4p_account_report_query", p4pAccountReportQuery)
    return nil
}

// P4pAccountReportQuery Getter
func (r AlibabaScbpEffectAccountListRequest) GetP4pAccountReportQuery() *AccountQuery {
    return r.p4pAccountReportQuery
}
