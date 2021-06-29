package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链渠道销售单创建接口 APIRequest
alibaba.ascp.channel.sales.order.create

阿里巴巴供应链渠道销售订单创建接口
*/
type AlibabaAscpChannelSalesOrderCreateRequest struct {
    model.Params

    // 请求参数
    createOrderRequest   *ExternalCreateSalesOrderRequest 

}

func NewAlibabaAscpChannelSalesOrderCreateRequest() *AlibabaAscpChannelSalesOrderCreateRequest{
    return &AlibabaAscpChannelSalesOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpChannelSalesOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.ascp.channel.sales.order.create"
}

func (r AlibabaAscpChannelSalesOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpChannelSalesOrderCreateRequest) SetCreateOrderRequest(createOrderRequest *ExternalCreateSalesOrderRequest) error {
    r.createOrderRequest = createOrderRequest
    r.Set("create_order_request", createOrderRequest)
    return nil
}

func (r AlibabaAscpChannelSalesOrderCreateRequest) GetCreateOrderRequest() *ExternalCreateSalesOrderRequest {
    return r.createOrderRequest
}

