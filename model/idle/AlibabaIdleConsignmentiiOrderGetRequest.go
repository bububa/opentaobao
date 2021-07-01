package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼寄卖V2订单查询 API请求
alibaba.idle.consignmentii.order.get

闲鱼寄卖V2服务商以闲鱼交易买家身份查询订单信息
*/
type AlibabaIdleConsignmentiiOrderGetAPIRequest struct {
    model.Params
    // 闲鱼订单ID
    _bizOrderId   int64
}

// 初始化AlibabaIdleConsignmentiiOrderGetAPIRequest对象
func NewAlibabaIdleConsignmentiiOrderGetRequest() *AlibabaIdleConsignmentiiOrderGetAPIRequest{
    return &AlibabaIdleConsignmentiiOrderGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentiiOrderGetAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.consignmentii.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentiiOrderGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 闲鱼订单ID
func (r *AlibabaIdleConsignmentiiOrderGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaIdleConsignmentiiOrderGetAPIRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
