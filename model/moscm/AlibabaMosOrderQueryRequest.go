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
    orderCriteria   *OrderCriteria
    // 分页信息
    paginator   *Paginator
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
func (r *AlibabaMosOrderQueryRequest) SetOrderCriteria(orderCriteria *OrderCriteria) error {
    r.orderCriteria = orderCriteria
    r.Set("order_criteria", orderCriteria)
    return nil
}

// OrderCriteria Getter
func (r AlibabaMosOrderQueryRequest) GetOrderCriteria() *OrderCriteria {
    return r.orderCriteria
}
// Paginator Setter
// 分页信息
func (r *AlibabaMosOrderQueryRequest) SetPaginator(paginator *Paginator) error {
    r.paginator = paginator
    r.Set("paginator", paginator)
    return nil
}

// Paginator Getter
func (r AlibabaMosOrderQueryRequest) GetPaginator() *Paginator {
    return r.paginator
}
