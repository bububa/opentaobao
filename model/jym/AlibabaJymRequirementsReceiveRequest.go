package jym

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易猫需求接单接口 API请求
alibaba.jym.requirements.receive

交易猫需求接单接口
*/
type AlibabaJymRequirementsReceiveRequest struct {
    model.Params
    // 需求id
    requirementId   string
    // 接单者手机号
    receiverMobile   string
    // 需求订单id
    requirementOrderId   string
}

// 初始化AlibabaJymRequirementsReceiveRequest对象
func NewAlibabaJymRequirementsReceiveRequest() *AlibabaJymRequirementsReceiveRequest{
    return &AlibabaJymRequirementsReceiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaJymRequirementsReceiveRequest) GetApiMethodName() string {
    return "alibaba.jym.requirements.receive"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaJymRequirementsReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequirementId Setter
// 需求id
func (r *AlibabaJymRequirementsReceiveRequest) SetRequirementId(requirementId string) error {
    r.requirementId = requirementId
    r.Set("requirement_id", requirementId)
    return nil
}

// RequirementId Getter
func (r AlibabaJymRequirementsReceiveRequest) GetRequirementId() string {
    return r.requirementId
}
// ReceiverMobile Setter
// 接单者手机号
func (r *AlibabaJymRequirementsReceiveRequest) SetReceiverMobile(receiverMobile string) error {
    r.receiverMobile = receiverMobile
    r.Set("receiver_mobile", receiverMobile)
    return nil
}

// ReceiverMobile Getter
func (r AlibabaJymRequirementsReceiveRequest) GetReceiverMobile() string {
    return r.receiverMobile
}
// RequirementOrderId Setter
// 需求订单id
func (r *AlibabaJymRequirementsReceiveRequest) SetRequirementOrderId(requirementOrderId string) error {
    r.requirementOrderId = requirementOrderId
    r.Set("requirement_order_id", requirementOrderId)
    return nil
}

// RequirementOrderId Getter
func (r AlibabaJymRequirementsReceiveRequest) GetRequirementOrderId() string {
    return r.requirementOrderId
}
