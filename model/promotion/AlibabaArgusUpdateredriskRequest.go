package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新红线价格 APIRequest
alibaba.argus.updateredrisk

商品健康中心新增红线价格规则
*/
type AlibabaArgusUpdateredriskRequest struct {
    model.Params

    // 红线价格参数
    redRiskUpdateFactor   *RedRiskUpdateFactor 

}

func NewAlibabaArgusUpdateredriskRequest() *AlibabaArgusUpdateredriskRequest{
    return &AlibabaArgusUpdateredriskRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaArgusUpdateredriskRequest) GetApiMethodName() string {
    return "alibaba.argus.updateredrisk"
}

func (r AlibabaArgusUpdateredriskRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaArgusUpdateredriskRequest) SetRedRiskUpdateFactor(redRiskUpdateFactor *RedRiskUpdateFactor) error {
    r.redRiskUpdateFactor = redRiskUpdateFactor
    r.Set("red_risk_update_factor", redRiskUpdateFactor)
    return nil
}

func (r AlibabaArgusUpdateredriskRequest) GetRedRiskUpdateFactor() *RedRiskUpdateFactor {
    return r.redRiskUpdateFactor
}

