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
type AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest struct {
    model.Params
    // 请求参数对象
    _paramnQuery   *PlaceOrderDTO
}

// 初始化AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressChargeCalculateRequest() *AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest{
    return &AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.charge.calculate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamnQuery Setter
// 请求参数对象
func (r *AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) SetParamnQuery(_paramnQuery *PlaceOrderDTO) error {
    r._paramnQuery = _paramnQuery
    r.Set("paramn_query", _paramnQuery)
    return nil
}

// ParamnQuery Getter
func (r AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) GetParamnQuery() *PlaceOrderDTO {
    return r._paramnQuery
}
