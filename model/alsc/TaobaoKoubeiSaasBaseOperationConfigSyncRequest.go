package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家基础经营设置信息同步 API请求
taobao.koubei.saas.base.operation.config.sync

ISV接入口碑SAAS后, 经营设置数据同步到口碑SAAS
*/
type TaobaoKoubeiSaasBaseOperationConfigSyncRequest struct {
    model.Params
    // 商户ID
    merchantId   string
    // 请求ID
    requestId   string
    // 业务类型。支付方式：payment_method
    bizType   string
    // 经营设置json串
    operationConfig   string
    // 操作员ID
    outerOperatorId   string
}

// 初始化TaobaoKoubeiSaasBaseOperationConfigSyncRequest对象
func NewTaobaoKoubeiSaasBaseOperationConfigSyncRequest() *TaobaoKoubeiSaasBaseOperationConfigSyncRequest{
    return &TaobaoKoubeiSaasBaseOperationConfigSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetApiMethodName() string {
    return "taobao.koubei.saas.base.operation.config.sync"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MerchantId Setter
// 商户ID
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetMerchantId(merchantId string) error {
    r.merchantId = merchantId
    r.Set("merchant_id", merchantId)
    return nil
}

// MerchantId Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetMerchantId() string {
    return r.merchantId
}
// RequestId Setter
// 请求ID
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

// RequestId Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetRequestId() string {
    return r.requestId
}
// BizType Setter
// 业务类型。支付方式：payment_method
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetBizType() string {
    return r.bizType
}
// OperationConfig Setter
// 经营设置json串
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetOperationConfig(operationConfig string) error {
    r.operationConfig = operationConfig
    r.Set("operation_config", operationConfig)
    return nil
}

// OperationConfig Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetOperationConfig() string {
    return r.operationConfig
}
// OuterOperatorId Setter
// 操作员ID
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetOuterOperatorId(outerOperatorId string) error {
    r.outerOperatorId = outerOperatorId
    r.Set("outer_operator_id", outerOperatorId)
    return nil
}

// OuterOperatorId Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetOuterOperatorId() string {
    return r.outerOperatorId
}
