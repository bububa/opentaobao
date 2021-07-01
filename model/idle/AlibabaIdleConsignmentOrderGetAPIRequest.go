package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼帮卖订单查询 API请求
alibaba.idle.consignment.order.get

闲鱼帮卖服务商以闲鱼交易买家身份查询订单信息
*/
type AlibabaIdleConsignmentOrderGetAPIRequest struct {
    model.Params
    // 闲鱼订单ID
    _bizOrderId   int64
}

// 初始化AlibabaIdleConsignmentOrderGetAPIRequest对象
func NewAlibabaIdleConsignmentOrderGetRequest() *AlibabaIdleConsignmentOrderGetAPIRequest{
    return &AlibabaIdleConsignmentOrderGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentOrderGetAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.consignment.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentOrderGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 闲鱼订单ID
func (r *AlibabaIdleConsignmentOrderGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaIdleConsignmentOrderGetAPIRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
