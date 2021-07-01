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
type AlibabaJymRequirementsReceiveAPIRequest struct {
    model.Params
    // 需求id
    _requirementId   string
    // 接单者手机号
    _receiverMobile   string
    // 需求订单id
    _requirementOrderId   string
}

// 初始化AlibabaJymRequirementsReceiveAPIRequest对象
func NewAlibabaJymRequirementsReceiveRequest() *AlibabaJymRequirementsReceiveAPIRequest{
    return &AlibabaJymRequirementsReceiveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaJymRequirementsReceiveAPIRequest) GetApiMethodName() string {
    return "alibaba.jym.requirements.receive"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaJymRequirementsReceiveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequirementId Setter
// 需求id
func (r *AlibabaJymRequirementsReceiveAPIRequest) SetRequirementId(_requirementId string) error {
    r._requirementId = _requirementId
    r.Set("requirement_id", _requirementId)
    return nil
}

// RequirementId Getter
func (r AlibabaJymRequirementsReceiveAPIRequest) GetRequirementId() string {
    return r._requirementId
}
// ReceiverMobile Setter
// 接单者手机号
func (r *AlibabaJymRequirementsReceiveAPIRequest) SetReceiverMobile(_receiverMobile string) error {
    r._receiverMobile = _receiverMobile
    r.Set("receiver_mobile", _receiverMobile)
    return nil
}

// ReceiverMobile Getter
func (r AlibabaJymRequirementsReceiveAPIRequest) GetReceiverMobile() string {
    return r._receiverMobile
}
// RequirementOrderId Setter
// 需求订单id
func (r *AlibabaJymRequirementsReceiveAPIRequest) SetRequirementOrderId(_requirementOrderId string) error {
    r._requirementOrderId = _requirementOrderId
    r.Set("requirement_order_id", _requirementOrderId)
    return nil
}

// RequirementOrderId Getter
func (r AlibabaJymRequirementsReceiveAPIRequest) GetRequirementOrderId() string {
    return r._requirementOrderId
}
