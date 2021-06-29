package icbuseller

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际站服务市场订单列表接口 API请求
alibaba.seller.vendor.order.list

返回服务商在服务市场的客户订单
*/
type AlibabaSellerVendorOrderListRequest struct {
    model.Params
    // 查询参数
    queryTradeDto   *QueryTradeDto
}

// 初始化AlibabaSellerVendorOrderListRequest对象
func NewAlibabaSellerVendorOrderListRequest() *AlibabaSellerVendorOrderListRequest{
    return &AlibabaSellerVendorOrderListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSellerVendorOrderListRequest) GetApiMethodName() string {
    return "alibaba.seller.vendor.order.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSellerVendorOrderListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryTradeDto Setter
// 查询参数
func (r *AlibabaSellerVendorOrderListRequest) SetQueryTradeDto(queryTradeDto *QueryTradeDto) error {
    r.queryTradeDto = queryTradeDto
    r.Set("query_trade_dto", queryTradeDto)
    return nil
}

// QueryTradeDto Getter
func (r AlibabaSellerVendorOrderListRequest) GetQueryTradeDto() *QueryTradeDto {
    return r.queryTradeDto
}
