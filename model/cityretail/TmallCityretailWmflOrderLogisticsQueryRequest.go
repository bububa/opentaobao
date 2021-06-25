package cityretail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
完美履约订单物流状态查询接口 APIRequest
tmall.cityretail.wmfl.order.logistics.query

完美履约订单物流状态查询接口，该接口只能查询未完结的履约单以及完结的3天内订单
*/
type TmallCityretailWmflOrderLogisticsQueryRequest struct {
    model.Params

    // 订单号
    mainOrderId   string 

}

func NewTmallCityretailWmflOrderLogisticsQueryRequest() *TmallCityretailWmflOrderLogisticsQueryRequest{
    return &TmallCityretailWmflOrderLogisticsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCityretailWmflOrderLogisticsQueryRequest) GetApiMethodName() string {
    return "tmall.cityretail.wmfl.order.logistics.query"
}

func (r TmallCityretailWmflOrderLogisticsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCityretailWmflOrderLogisticsQueryRequest) SetMainOrderId(mainOrderId string) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TmallCityretailWmflOrderLogisticsQueryRequest) GetMainOrderId() string {
    return r.mainOrderId
}

