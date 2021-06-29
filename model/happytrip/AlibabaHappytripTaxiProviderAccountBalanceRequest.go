package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商渠道余额 API请求
alibaba.happytrip.taxi.provider.account.balance

查询不同供应商不同渠道账户余额
*/
type AlibabaHappytripTaxiProviderAccountBalanceRequest struct {
    model.Params
    // 成本中心代码，用于区分不同的分账账号
    _costCenter   string
}

// 初始化AlibabaHappytripTaxiProviderAccountBalanceRequest对象
func NewAlibabaHappytripTaxiProviderAccountBalanceRequest() *AlibabaHappytripTaxiProviderAccountBalanceRequest{
    return &AlibabaHappytripTaxiProviderAccountBalanceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiProviderAccountBalanceRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.provider.account.balance"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiProviderAccountBalanceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CostCenter Setter
// 成本中心代码，用于区分不同的分账账号
func (r *AlibabaHappytripTaxiProviderAccountBalanceRequest) SetCostCenter(_costCenter string) error {
    r._costCenter = _costCenter
    r.Set("cost_center", _costCenter)
    return nil
}

// CostCenter Getter
func (r AlibabaHappytripTaxiProviderAccountBalanceRequest) GetCostCenter() string {
    return r._costCenter
}
