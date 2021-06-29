package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询订单信息 APIRequest
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

func NewAlibabaMosOrderQueryRequest() *AlibabaMosOrderQueryRequest{
    return &AlibabaMosOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.mos.order.query"
}

func (r AlibabaMosOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosOrderQueryRequest) SetOrderCriteria(orderCriteria *OrderCriteria) error {
    r.orderCriteria = orderCriteria
    r.Set("order_criteria", orderCriteria)
    return nil
}

func (r AlibabaMosOrderQueryRequest) GetOrderCriteria() *OrderCriteria {
    return r.orderCriteria
}

func (r *AlibabaMosOrderQueryRequest) SetPaginator(paginator *Paginator) error {
    r.paginator = paginator
    r.Set("paginator", paginator)
    return nil
}

func (r AlibabaMosOrderQueryRequest) GetPaginator() *Paginator {
    return r.paginator
}

