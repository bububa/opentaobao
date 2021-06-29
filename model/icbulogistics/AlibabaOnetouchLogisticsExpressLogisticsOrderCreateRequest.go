package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
快递下单 APIRequest
alibaba.onetouch.logistics.express.logistics.order.create

快递下单
*/
type AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest struct {
    model.Params

    // 请求参数对象
    paramnQuery   *PlaceOrderDTO 

}

func NewAlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest() *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest{
    return &AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.logistics.order.create"
}

func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest) SetParamnQuery(paramnQuery *PlaceOrderDTO) error {
    r.paramnQuery = paramnQuery
    r.Set("paramn_query", paramnQuery)
    return nil
}

func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest) GetParamnQuery() *PlaceOrderDTO {
    return r.paramnQuery
}

