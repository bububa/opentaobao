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
type AlibabaIdleConsignmentiiOrderGetRequest struct {
    model.Params
    // 闲鱼订单ID
    bizOrderId   int64
}

// 初始化AlibabaIdleConsignmentiiOrderGetRequest对象
func NewAlibabaIdleConsignmentiiOrderGetRequest() *AlibabaIdleConsignmentiiOrderGetRequest{
    return &AlibabaIdleConsignmentiiOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentiiOrderGetRequest) GetApiMethodName() string {
    return "alibaba.idle.consignmentii.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentiiOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 闲鱼订单ID
func (r *AlibabaIdleConsignmentiiOrderGetRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaIdleConsignmentiiOrderGetRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}
