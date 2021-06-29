package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询订单信息 API请求
alibaba.mos.order.query

查询多笔交易信息
*/
type AlibabaMosOrderQueryRequest struct {
    model.Params
    // 订单查询
    _orderCriteria   *OrderCriteria
    // 分页信息
    _paginator   *Paginator
}

// 初始化AlibabaMosOrderQueryRequest对象
func NewAlibabaMosOrderQueryRequest() *AlibabaMosOrderQueryRequest{
    return &AlibabaMosOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.mos.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCriteria Setter
// 订单查询
func (r *AlibabaMosOrderQueryRequest) SetOrderCriteria(_orderCriteria *OrderCriteria) error {
    r._orderCriteria = _orderCriteria
    r.Set("order_criteria", _orderCriteria)
    return nil
}

// OrderCriteria Getter
func (r AlibabaMosOrderQueryRequest) GetOrderCriteria() *OrderCriteria {
    return r._orderCriteria
}
// Paginator Setter
// 分页信息
func (r *AlibabaMosOrderQueryRequest) SetPaginator(_paginator *Paginator) error {
    r._paginator = _paginator
    r.Set("paginator", _paginator)
    return nil
}

// Paginator Getter
func (r AlibabaMosOrderQueryRequest) GetPaginator() *Paginator {
    return r._paginator
}
