package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商服务开通状态 APIRequest
alibaba.happytrip.taxi.servicestatus.get

获取服务供应商在每个地区的服务开通状态、支持的车型等
*/
type AlibabaHappytripTaxiServicestatusGetRequest struct {
    model.Params

    // 成本中心代码，用于区分不同的分账账号
    costCenter   string 

}

func NewAlibabaHappytripTaxiServicestatusGetRequest() *AlibabaHappytripTaxiServicestatusGetRequest{
    return &AlibabaHappytripTaxiServicestatusGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiServicestatusGetRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.servicestatus.get"
}

func (r AlibabaHappytripTaxiServicestatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiServicestatusGetRequest) SetCostCenter(costCenter string) error {
    r.costCenter = costCenter
    r.Set("cost_center", costCenter)
    return nil
}

func (r AlibabaHappytripTaxiServicestatusGetRequest) GetCostCenter() string {
    return r.costCenter
}

