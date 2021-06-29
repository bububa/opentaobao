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
    id   int64
    // 取消原因说明
    comments   string
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
func (r *TmallServiceSettleadjustmentCancelRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TmallServiceSettleadjustmentCancelRequest) GetId() int64 {
    return r.id
}
// Comments Setter
// 取消原因说明
func (r *TmallServiceSettleadjustmentCancelRequest) SetComments(comments string) error {
    r.comments = comments
    r.Set("comments", comments)
    return nil
}

// Comments Getter
func (r TmallServiceSettleadjustmentCancelRequest) GetComments() string {
    return r.comments
}
