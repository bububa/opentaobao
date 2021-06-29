package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询结算调整单单条记录 API请求
tmall.service.settleadjustment.get

提供给服务商通过结算调整单id获取结算调整单信息
*/
type TmallServiceSettleadjustmentGetRequest struct {
    model.Params
    // 结算调整单ID
    id   int64
}

// 初始化TmallServiceSettleadjustmentGetRequest对象
func NewTmallServiceSettleadjustmentGetRequest() *TmallServiceSettleadjustmentGetRequest{
    return &TmallServiceSettleadjustmentGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentGetRequest) GetApiMethodName() string {
    return "tmall.service.settleadjustment.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 结算调整单ID
func (r *TmallServiceSettleadjustmentGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TmallServiceSettleadjustmentGetRequest) GetId() int64 {
    return r.id
}
