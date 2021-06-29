package icbuassurance

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
icbu信保账户信息 API请求
alibaba.icbu.trade.assurance.account.get

icbu交易信用保障开通状态&额度信息查询
*/
type AlibabaIcbuTradeAssuranceAccountGetRequest struct {
    model.Params
}

// 初始化AlibabaIcbuTradeAssuranceAccountGetRequest对象
func NewAlibabaIcbuTradeAssuranceAccountGetRequest() *AlibabaIcbuTradeAssuranceAccountGetRequest{
    return &AlibabaIcbuTradeAssuranceAccountGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuTradeAssuranceAccountGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.trade.assurance.account.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuTradeAssuranceAccountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
