package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单查询V1 API请求
alibaba.idle.recycle.order.query

查询回收订单
*/
type AlibabaIdleRecycleOrderQueryRequest struct {
    model.Params
    // 订单号
    _bizOrderId   int64
}

// 初始化AlibabaIdleRecycleOrderQueryRequest对象
func NewAlibabaIdleRecycleOrderQueryRequest() *AlibabaIdleRecycleOrderQueryRequest{
    return &AlibabaIdleRecycleOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单号
func (r *AlibabaIdleRecycleOrderQueryRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaIdleRecycleOrderQueryRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
