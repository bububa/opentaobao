package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链渠道销售单创建接口 API请求
alibaba.ascp.channel.sales.order.create

阿里巴巴供应链渠道销售订单创建接口
*/
type AlibabaAscpChannelSalesOrderCreateRequest struct {
    model.Params
    // 请求参数
    _createOrderRequest   *ExternalCreateSalesOrderRequest
}

// 初始化AlibabaAscpChannelSalesOrderCreateRequest对象
func NewAlibabaAscpChannelSalesOrderCreateRequest() *AlibabaAscpChannelSalesOrderCreateRequest{
    return &AlibabaAscpChannelSalesOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSalesOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.sales.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSalesOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateOrderRequest Setter
// 请求参数
func (r *AlibabaAscpChannelSalesOrderCreateRequest) SetCreateOrderRequest(_createOrderRequest *ExternalCreateSalesOrderRequest) error {
    r._createOrderRequest = _createOrderRequest
    r.Set("create_order_request", _createOrderRequest)
    return nil
}

// CreateOrderRequest Getter
func (r AlibabaAscpChannelSalesOrderCreateRequest) GetCreateOrderRequest() *ExternalCreateSalesOrderRequest {
    return r._createOrderRequest
}
