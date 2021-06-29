package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询交易退货信息 APIRequest
alibaba.mos.order.refund.list.get

批量查询多个退货单的退货明细
*/
type AlibabaMosOrderRefundListGetRequest struct {
    model.Params

    // 退换货查询条件
    rmaCriteria   *RmaCriteria 

    // 分页信息
    paginator   *Paginator 

}

func NewAlibabaMosOrderRefundListGetRequest() *AlibabaMosOrderRefundListGetRequest{
    return &AlibabaMosOrderRefundListGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosOrderRefundListGetRequest) GetApiMethodName() string {
    return "alibaba.mos.order.refund.list.get"
}

func (r AlibabaMosOrderRefundListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosOrderRefundListGetRequest) SetRmaCriteria(rmaCriteria *RmaCriteria) error {
    r.rmaCriteria = rmaCriteria
    r.Set("rma_criteria", rmaCriteria)
    return nil
}

func (r AlibabaMosOrderRefundListGetRequest) GetRmaCriteria() *RmaCriteria {
    return r.rmaCriteria
}

func (r *AlibabaMosOrderRefundListGetRequest) SetPaginator(paginator *Paginator) error {
    r.paginator = paginator
    r.Set("paginator", paginator)
    return nil
}

func (r AlibabaMosOrderRefundListGetRequest) GetPaginator() *Paginator {
    return r.paginator
}

