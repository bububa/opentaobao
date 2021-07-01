package idleitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单查询V2 API请求
alibaba.idle.recycle.order.get

闲鱼回收业务中,外部回收商作为交易上买家,闲鱼用户下单后,需要回收商主动拉取交易订单
*/
type AlibabaIdleRecycleOrderGetAPIRequest struct {
    model.Params
    // 订单号
    _bizOrderId   int64
}

// 初始化AlibabaIdleRecycleOrderGetAPIRequest对象
func NewAlibabaIdleRecycleOrderGetRequest() *AlibabaIdleRecycleOrderGetAPIRequest{
    return &AlibabaIdleRecycleOrderGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderGetAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单号
func (r *AlibabaIdleRecycleOrderGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaIdleRecycleOrderGetAPIRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
