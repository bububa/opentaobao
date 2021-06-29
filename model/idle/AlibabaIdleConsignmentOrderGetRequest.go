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
type AlibabaIdleConsignmentOrderGetRequest struct {
    model.Params
    // 闲鱼订单ID
    _bizOrderId   int64
}

// 初始化AlibabaIdleConsignmentOrderGetRequest对象
func NewAlibabaIdleConsignmentOrderGetRequest() *AlibabaIdleConsignmentOrderGetRequest{
    return &AlibabaIdleConsignmentOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentOrderGetRequest) GetApiMethodName() string {
    return "alibaba.idle.consignment.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 闲鱼订单ID
func (r *AlibabaIdleConsignmentOrderGetRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaIdleConsignmentOrderGetRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
