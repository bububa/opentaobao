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
    _requirementId   string
    // 接单者手机号
    _receiverMobile   string
    // 需求订单id
    _requirementOrderId   string
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
func (r *AlibabaJymRequirementsReceiveRequest) SetRequirementId(_requirementId string) error {
    r._requirementId = _requirementId
    r.Set("requirement_id", _requirementId)
    return nil
}

// RequirementId Getter
func (r AlibabaJymRequirementsReceiveRequest) GetRequirementId() string {
    return r._requirementId
}
// ReceiverMobile Setter
// 接单者手机号
func (r *AlibabaJymRequirementsReceiveRequest) SetReceiverMobile(_receiverMobile string) error {
    r._receiverMobile = _receiverMobile
    r.Set("receiver_mobile", _receiverMobile)
    return nil
}

// ReceiverMobile Getter
func (r AlibabaJymRequirementsReceiveRequest) GetReceiverMobile() string {
    return r._receiverMobile
}
// RequirementOrderId Setter
// 需求订单id
func (r *AlibabaJymRequirementsReceiveRequest) SetRequirementOrderId(_requirementOrderId string) error {
    r._requirementOrderId = _requirementOrderId
    r.Set("requirement_order_id", _requirementOrderId)
    return nil
}

// RequirementOrderId Getter
func (r AlibabaJymRequirementsReceiveRequest) GetRequirementOrderId() string {
    return r._requirementOrderId
}
