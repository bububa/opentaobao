package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
周期送配送明细查询 APIRequest
tmall.nr.zqs.plan.query

周期送配送明细查询
*/
type TmallNrZqsPlanQueryRequest struct {
    model.Params

    // 交易子订单id
    detailOrderId   int64 

}

func NewTmallNrZqsPlanQueryRequest() *TmallNrZqsPlanQueryRequest{
    return &TmallNrZqsPlanQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrZqsPlanQueryRequest) GetApiMethodName() string {
    return "tmall.nr.zqs.plan.query"
}

func (r TmallNrZqsPlanQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrZqsPlanQueryRequest) SetDetailOrderId(detailOrderId int64) error {
    r.detailOrderId = detailOrderId
    r.Set("detail_order_id", detailOrderId)
    return nil
}

func (r TmallNrZqsPlanQueryRequest) GetDetailOrderId() int64 {
    return r.detailOrderId
}

