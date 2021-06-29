package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单揽件 API请求
tmall.servicecenter.workcard.collect

服务工单揽件接口
*/
type TmallServicecenterWorkcardCollectRequest struct {
    model.Params
    // 工单id
    workcardId   int64
    // 买家id
    buyerId   int64
    // 扩展信息
    attributes   string
}

// 初始化TmallServicecenterWorkcardCollectRequest对象
func NewTmallServicecenterWorkcardCollectRequest() *TmallServicecenterWorkcardCollectRequest{
    return &TmallServicecenterWorkcardCollectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardCollectRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.collect"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardCollectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardCollectRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardCollectRequest) GetWorkcardId() int64 {
    return r.workcardId
}
// BuyerId Setter
// 买家id
func (r *TmallServicecenterWorkcardCollectRequest) SetBuyerId(buyerId int64) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

// BuyerId Getter
func (r TmallServicecenterWorkcardCollectRequest) GetBuyerId() int64 {
    return r.buyerId
}
// Attributes Setter
// 扩展信息
func (r *TmallServicecenterWorkcardCollectRequest) SetAttributes(attributes string) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

// Attributes Getter
func (r TmallServicecenterWorkcardCollectRequest) GetAttributes() string {
    return r.attributes
}
