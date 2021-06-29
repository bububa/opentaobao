package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商渠道余额 APIRequest
alibaba.happytrip.taxi.provider.account.balance

查询不同供应商不同渠道账户余额
*/
type AlibabaHappytripTaxiProviderAccountBalanceRequest struct {
    model.Params

    // 成本中心代码，用于区分不同的分账账号
    costCenter   string 

}

func NewAlibabaHappytripTaxiProviderAccountBalanceRequest() *AlibabaHappytripTaxiProviderAccountBalanceRequest{
    return &AlibabaHappytripTaxiProviderAccountBalanceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiProviderAccountBalanceRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.provider.account.balance"
}

func (r AlibabaHappytripTaxiProviderAccountBalanceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiProviderAccountBalanceRequest) SetCostCenter(costCenter string) error {
    r.costCenter = costCenter
    r.Set("cost_center", costCenter)
    return nil
}

func (r AlibabaHappytripTaxiProviderAccountBalanceRequest) GetCostCenter() string {
    return r.costCenter
}

