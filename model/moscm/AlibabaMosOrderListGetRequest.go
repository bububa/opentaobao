package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询订单交易 APIRequest
alibaba.mos.order.list.get

批量查询交易信息
*/
type AlibabaMosOrderListGetRequest struct {
    model.Params

    // 订单查询条件
    orderCriteria   *OrderCriteria 

    // 分页信息
    paginator   *Paginator 

}

func NewAlibabaMosOrderListGetRequest() *AlibabaMosOrderListGetRequest{
    return &AlibabaMosOrderListGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosOrderListGetRequest) GetApiMethodName() string {
    return "alibaba.mos.order.list.get"
}

func (r AlibabaMosOrderListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosOrderListGetRequest) SetOrderCriteria(orderCriteria *OrderCriteria) error {
    r.orderCriteria = orderCriteria
    r.Set("order_criteria", orderCriteria)
    return nil
}

func (r AlibabaMosOrderListGetRequest) GetOrderCriteria() *OrderCriteria {
    return r.orderCriteria
}

func (r *AlibabaMosOrderListGetRequest) SetPaginator(paginator *Paginator) error {
    r.paginator = paginator
    r.Set("paginator", paginator)
    return nil
}

func (r AlibabaMosOrderListGetRequest) GetPaginator() *Paginator {
    return r.paginator
}

