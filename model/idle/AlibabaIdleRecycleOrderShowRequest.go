package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单查询V1.1 API请求
alibaba.idle.recycle.order.show

查询回收订单
*/
type AlibabaIdleRecycleOrderShowRequest struct {
    model.Params
    // 订单号
    bizOrderId   int64
}

// 初始化AlibabaIdleRecycleOrderShowRequest对象
func NewAlibabaIdleRecycleOrderShowRequest() *AlibabaIdleRecycleOrderShowRequest{
    return &AlibabaIdleRecycleOrderShowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderShowRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.order.show"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderShowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单号
func (r *AlibabaIdleRecycleOrderShowRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaIdleRecycleOrderShowRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}
