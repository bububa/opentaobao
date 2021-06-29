package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计算快递运费&下单参数校验 API请求
alibaba.onetouch.logistics.express.charge.calculate

计算快递运费、下单参数校验
*/
type AlibabaOnetouchLogisticsExpressChargeCalculateRequest struct {
    model.Params
    // 请求参数对象
    paramnQuery   *PlaceOrderDTO
}

// 初始化AlibabaOnetouchLogisticsExpressChargeCalculateRequest对象
func NewAlibabaOnetouchLogisticsExpressChargeCalculateRequest() *AlibabaOnetouchLogisticsExpressChargeCalculateRequest{
    return &AlibabaOnetouchLogisticsExpressChargeCalculateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressChargeCalculateRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.charge.calculate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressChargeCalculateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamnQuery Setter
// 请求参数对象
func (r *AlibabaOnetouchLogisticsExpressChargeCalculateRequest) SetParamnQuery(paramnQuery *PlaceOrderDTO) error {
    r.paramnQuery = paramnQuery
    r.Set("paramn_query", paramnQuery)
    return nil
}

// ParamnQuery Getter
func (r AlibabaOnetouchLogisticsExpressChargeCalculateRequest) GetParamnQuery() *PlaceOrderDTO {
    return r.paramnQuery
}
