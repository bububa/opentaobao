package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消结算调整单 API请求
tmall.service.settleadjustment.cancel

提供给服务商在对取消已经发起的结算调整单。
通过说明调整单ID进行结算调整单取消。
*/
type TmallServiceSettleadjustmentCancelRequest struct {
    model.Params
    // 结算调整单ID
    _id   int64
    // 取消原因说明
    _comments   string
}

// 初始化TmallServiceSettleadjustmentCancelRequest对象
func NewTmallServiceSettleadjustmentCancelRequest() *TmallServiceSettleadjustmentCancelRequest{
    return &TmallServiceSettleadjustmentCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentCancelRequest) GetApiMethodName() string {
    return "tmall.service.settleadjustment.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 结算调整单ID
func (r *TmallServiceSettleadjustmentCancelRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TmallServiceSettleadjustmentCancelRequest) GetId() int64 {
    return r._id
}
// Comments Setter
// 取消原因说明
func (r *TmallServiceSettleadjustmentCancelRequest) SetComments(_comments string) error {
    r._comments = _comments
    r.Set("comments", _comments)
    return nil
}

// Comments Getter
func (r TmallServiceSettleadjustmentCancelRequest) GetComments() string {
    return r._comments
}
