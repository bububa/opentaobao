package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台结算出账信息 API请求
alibaba.service.billing.query

服务平台结算单明细查询服务
*/
type AlibabaServiceBillingQueryRequest struct {
    model.Params
    // 账单查询开始时间。格式示例 2019-03-26 17:15:28
    _gmtCreateStart   string
    // 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
    _gmtCreateEnd   string
}

// 初始化AlibabaServiceBillingQueryRequest对象
func NewAlibabaServiceBillingQueryRequest() *AlibabaServiceBillingQueryRequest{
    return &AlibabaServiceBillingQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaServiceBillingQueryRequest) GetApiMethodName() string {
    return "alibaba.service.billing.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaServiceBillingQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GmtCreateStart Setter
// 账单查询开始时间。格式示例 2019-03-26 17:15:28
func (r *AlibabaServiceBillingQueryRequest) SetGmtCreateStart(_gmtCreateStart string) error {
    r._gmtCreateStart = _gmtCreateStart
    r.Set("gmt_create_start", _gmtCreateStart)
    return nil
}

// GmtCreateStart Getter
func (r AlibabaServiceBillingQueryRequest) GetGmtCreateStart() string {
    return r._gmtCreateStart
}
// GmtCreateEnd Setter
// 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
func (r *AlibabaServiceBillingQueryRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
    r._gmtCreateEnd = _gmtCreateEnd
    r.Set("gmt_create_end", _gmtCreateEnd)
    return nil
}

// GmtCreateEnd Getter
func (r AlibabaServiceBillingQueryRequest) GetGmtCreateEnd() string {
    return r._gmtCreateEnd
}
