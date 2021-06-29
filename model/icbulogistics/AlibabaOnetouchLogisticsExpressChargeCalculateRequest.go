package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计算快递运费&下单参数校验 APIRequest
alibaba.onetouch.logistics.express.charge.calculate

计算快递运费、下单参数校验
*/
type AlibabaOnetouchLogisticsExpressChargeCalculateRequest struct {
    model.Params

    // 请求参数对象
    paramnQuery   *PlaceOrderDTO 

}

func NewAlibabaOnetouchLogisticsExpressChargeCalculateRequest() *AlibabaOnetouchLogisticsExpressChargeCalculateRequest{
    return &AlibabaOnetouchLogisticsExpressChargeCalculateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOnetouchLogisticsExpressChargeCalculateRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.charge.calculate"
}

func (r AlibabaOnetouchLogisticsExpressChargeCalculateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaOnetouchLogisticsExpressChargeCalculateRequest) SetParamnQuery(paramnQuery *PlaceOrderDTO) error {
    r.paramnQuery = paramnQuery
    r.Set("paramn_query", paramnQuery)
    return nil
}

func (r AlibabaOnetouchLogisticsExpressChargeCalculateRequest) GetParamnQuery() *PlaceOrderDTO {
    return r.paramnQuery
}

