package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼验货担保订单详情查询V1 API请求
alibaba.idle.appraise.order.query

鉴定商调用该接口获取订单状态
*/
type AlibabaIdleAppraiseOrderQueryRequest struct {
    model.Params
    // orderId
    bizOrderId   int64
}

// 初始化AlibabaIdleAppraiseOrderQueryRequest对象
func NewAlibabaIdleAppraiseOrderQueryRequest() *AlibabaIdleAppraiseOrderQueryRequest{
    return &AlibabaIdleAppraiseOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleAppraiseOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.appraise.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleAppraiseOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// orderId
func (r *AlibabaIdleAppraiseOrderQueryRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaIdleAppraiseOrderQueryRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}
