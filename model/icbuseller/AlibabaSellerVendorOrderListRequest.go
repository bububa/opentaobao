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
type AlibabaSellerVendorOrderListAPIRequest struct {
    model.Params
    // 查询参数
    _queryTradeDto   *QueryTradeDTO
}

// 初始化AlibabaSellerVendorOrderListAPIRequest对象
func NewAlibabaSellerVendorOrderListRequest() *AlibabaSellerVendorOrderListAPIRequest{
    return &AlibabaSellerVendorOrderListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSellerVendorOrderListAPIRequest) GetApiMethodName() string {
    return "alibaba.seller.vendor.order.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSellerVendorOrderListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryTradeDto Setter
// 查询参数
func (r *AlibabaSellerVendorOrderListAPIRequest) SetQueryTradeDto(_queryTradeDto *QueryTradeDTO) error {
    r._queryTradeDto = _queryTradeDto
    r.Set("query_trade_dto", _queryTradeDto)
    return nil
}

// QueryTradeDto Getter
func (r AlibabaSellerVendorOrderListAPIRequest) GetQueryTradeDto() *QueryTradeDTO {
    return r._queryTradeDto
}
