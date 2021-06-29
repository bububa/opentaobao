package icbuseller

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际站服务市场订单列表接口 APIRequest
alibaba.seller.vendor.order.list

返回服务商在服务市场的客户订单
*/
type AlibabaSellerVendorOrderListRequest struct {
    model.Params

    // 查询参数
    queryTradeDto   *QueryTradeDto 

}

func NewAlibabaSellerVendorOrderListRequest() *AlibabaSellerVendorOrderListRequest{
    return &AlibabaSellerVendorOrderListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSellerVendorOrderListRequest) GetApiMethodName() string {
    return "alibaba.seller.vendor.order.list"
}

func (r AlibabaSellerVendorOrderListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSellerVendorOrderListRequest) SetQueryTradeDto(queryTradeDto *QueryTradeDto) error {
    r.queryTradeDto = queryTradeDto
    r.Set("query_trade_dto", queryTradeDto)
    return nil
}

func (r AlibabaSellerVendorOrderListRequest) GetQueryTradeDto() *QueryTradeDto {
    return r.queryTradeDto
}

