package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询订单交易 API请求
alibaba.mos.order.list.get

批量查询交易信息
*/
type AlibabaMosOrderListGetRequest struct {
    model.Params
    // 订单查询条件
    _orderCriteria   *OrderCriteria
    // 分页信息
    _paginator   *Paginator
}

// 初始化AlibabaMosOrderListGetRequest对象
func NewAlibabaMosOrderListGetRequest() *AlibabaMosOrderListGetRequest{
    return &AlibabaMosOrderListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosOrderListGetRequest) GetApiMethodName() string {
    return "alibaba.mos.order.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosOrderListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCriteria Setter
// 订单查询条件
func (r *AlibabaMosOrderListGetRequest) SetOrderCriteria(_orderCriteria *OrderCriteria) error {
    r._orderCriteria = _orderCriteria
    r.Set("order_criteria", _orderCriteria)
    return nil
}

// OrderCriteria Getter
func (r AlibabaMosOrderListGetRequest) GetOrderCriteria() *OrderCriteria {
    return r._orderCriteria
}
// Paginator Setter
// 分页信息
func (r *AlibabaMosOrderListGetRequest) SetPaginator(_paginator *Paginator) error {
    r._paginator = _paginator
    r.Set("paginator", _paginator)
    return nil
}

// Paginator Getter
func (r AlibabaMosOrderListGetRequest) GetPaginator() *Paginator {
    return r._paginator
}