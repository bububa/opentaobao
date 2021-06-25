package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台结算出账信息 APIRequest
alibaba.service.billing.query

服务平台结算单明细查询服务
*/
type AlibabaServiceBillingQueryRequest struct {
    model.Params

    // 账单查询开始时间。格式示例 2019-03-26 17:15:28
    gmtCreateStart   string 

    // 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
    gmtCreateEnd   string 

}

func NewAlibabaServiceBillingQueryRequest() *AlibabaServiceBillingQueryRequest{
    return &AlibabaServiceBillingQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaServiceBillingQueryRequest) GetApiMethodName() string {
    return "alibaba.service.billing.query"
}

func (r AlibabaServiceBillingQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaServiceBillingQueryRequest) SetGmtCreateStart(gmtCreateStart string) error {
    r.gmtCreateStart = gmtCreateStart
    r.Set("gmt_create_start", gmtCreateStart)
    return nil
}

func (r AlibabaServiceBillingQueryRequest) GetGmtCreateStart() string {
    return r.gmtCreateStart
}

func (r *AlibabaServiceBillingQueryRequest) SetGmtCreateEnd(gmtCreateEnd string) error {
    r.gmtCreateEnd = gmtCreateEnd
    r.Set("gmt_create_end", gmtCreateEnd)
    return nil
}

func (r AlibabaServiceBillingQueryRequest) GetGmtCreateEnd() string {
    return r.gmtCreateEnd
}

