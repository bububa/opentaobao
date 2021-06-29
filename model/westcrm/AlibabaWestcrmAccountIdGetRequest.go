package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据支付宝id查询IB的id API请求
alibaba.westcrm.account.id.get

根据支付宝id查询IB的id
*/
type AlibabaWestcrmAccountIdGetRequest struct {
    model.Params
    // 支付宝id
    alipayId   string
}

// 初始化AlibabaWestcrmAccountIdGetRequest对象
func NewAlibabaWestcrmAccountIdGetRequest() *AlibabaWestcrmAccountIdGetRequest{
    return &AlibabaWestcrmAccountIdGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmAccountIdGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.account.id.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmAccountIdGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayId Setter
// 支付宝id
func (r *AlibabaWestcrmAccountIdGetRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

// AlipayId Getter
func (r AlibabaWestcrmAccountIdGetRequest) GetAlipayId() string {
    return r.alipayId
}
