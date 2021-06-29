package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定担保资金订单履约 API请求
alibaba.idle.apprize.order.fulfillment

服务商针对自己的服务订单进行履约
*/
type AlibabaIdleApprizeOrderFulfillmentRequest struct {
    model.Params
    // deal：服务商收取费用、refund：退款给付款方
    action   string
    // 天猫服务工单Id
    workCardId   int64
}

// 初始化AlibabaIdleApprizeOrderFulfillmentRequest对象
func NewAlibabaIdleApprizeOrderFulfillmentRequest() *AlibabaIdleApprizeOrderFulfillmentRequest{
    return &AlibabaIdleApprizeOrderFulfillmentRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleApprizeOrderFulfillmentRequest) GetApiMethodName() string {
    return "alibaba.idle.apprize.order.fulfillment"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleApprizeOrderFulfillmentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Action Setter
// deal：服务商收取费用、refund：退款给付款方
func (r *AlibabaIdleApprizeOrderFulfillmentRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

// Action Getter
func (r AlibabaIdleApprizeOrderFulfillmentRequest) GetAction() string {
    return r.action
}
// WorkCardId Setter
// 天猫服务工单Id
func (r *AlibabaIdleApprizeOrderFulfillmentRequest) SetWorkCardId(workCardId int64) error {
    r.workCardId = workCardId
    r.Set("work_card_id", workCardId)
    return nil
}

// WorkCardId Getter
func (r AlibabaIdleApprizeOrderFulfillmentRequest) GetWorkCardId() int64 {
    return r.workCardId
}
