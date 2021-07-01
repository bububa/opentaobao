package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
快递下单 API请求
alibaba.onetouch.logistics.express.logistics.order.create

快递下单
*/
type AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest struct {
    model.Params
    // 请求参数对象
    _paramnQuery   *PlaceOrderDTO
}

// 初始化AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest() *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest{
    return &AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.logistics.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamnQuery Setter
// 请求参数对象
func (r *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest) SetParamnQuery(_paramnQuery *PlaceOrderDTO) error {
    r._paramnQuery = _paramnQuery
    r.Set("paramn_query", _paramnQuery)
    return nil
}

// ParamnQuery Getter
func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest) GetParamnQuery() *PlaceOrderDTO {
    return r._paramnQuery
}
