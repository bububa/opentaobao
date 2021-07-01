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
type TmallServiceSettleadjustmentGetAPIRequest struct {
    model.Params
    // 结算调整单ID
    _id   int64
}

// 初始化TmallServiceSettleadjustmentGetAPIRequest对象
func NewTmallServiceSettleadjustmentGetRequest() *TmallServiceSettleadjustmentGetAPIRequest{
    return &TmallServiceSettleadjustmentGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentGetAPIRequest) GetApiMethodName() string {
    return "tmall.service.settleadjustment.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 结算调整单ID
func (r *TmallServiceSettleadjustmentGetAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TmallServiceSettleadjustmentGetAPIRequest) GetId() int64 {
    return r._id
}
