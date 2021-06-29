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
type AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest struct {
    model.Params
    // 请求参数对象
    _paramnQuery   *PlaceOrderDTO
}

// 初始化AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest对象
func NewAlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest() *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest{
    return &AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.logistics.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamnQuery Setter
// 请求参数对象
func (r *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest) SetParamnQuery(_paramnQuery *PlaceOrderDTO) error {
    r._paramnQuery = _paramnQuery
    r.Set("paramn_query", _paramnQuery)
    return nil
}

// ParamnQuery Getter
func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest) GetParamnQuery() *PlaceOrderDTO {
    return r._paramnQuery
}
