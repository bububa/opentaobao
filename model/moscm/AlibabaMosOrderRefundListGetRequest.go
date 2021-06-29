package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询交易退货信息 API请求
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

// 初始化AlibabaMosOrderRefundListGetRequest对象
func NewAlibabaMosOrderRefundListGetRequest() *AlibabaMosOrderRefundListGetRequest{
    return &AlibabaMosOrderRefundListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosOrderRefundListGetRequest) GetApiMethodName() string {
    return "alibaba.mos.order.refund.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosOrderRefundListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RmaCriteria Setter
// 退换货查询条件
func (r *AlibabaMosOrderRefundListGetRequest) SetRmaCriteria(rmaCriteria *RmaCriteria) error {
    r.rmaCriteria = rmaCriteria
    r.Set("rma_criteria", rmaCriteria)
    return nil
}

// RmaCriteria Getter
func (r AlibabaMosOrderRefundListGetRequest) GetRmaCriteria() *RmaCriteria {
    return r.rmaCriteria
}
// Paginator Setter
// 分页信息
func (r *AlibabaMosOrderRefundListGetRequest) SetPaginator(paginator *Paginator) error {
    r.paginator = paginator
    r.Set("paginator", paginator)
    return nil
}

// Paginator Getter
func (r AlibabaMosOrderRefundListGetRequest) GetPaginator() *Paginator {
    return r.paginator
}
