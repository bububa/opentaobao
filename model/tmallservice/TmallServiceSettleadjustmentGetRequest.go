package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询结算调整单单条记录 APIRequest
tmall.service.settleadjustment.get

提供给服务商通过结算调整单id获取结算调整单信息
*/
type TmallServiceSettleadjustmentGetRequest struct {
    model.Params

    // 结算调整单ID
    id   int64 

}

func NewTmallServiceSettleadjustmentGetRequest() *TmallServiceSettleadjustmentGetRequest{
    return &TmallServiceSettleadjustmentGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServiceSettleadjustmentGetRequest) GetApiMethodName() string {
    return "tmall.service.settleadjustment.get"
}

func (r TmallServiceSettleadjustmentGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServiceSettleadjustmentGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TmallServiceSettleadjustmentGetRequest) GetId() int64 {
    return r.id
}

