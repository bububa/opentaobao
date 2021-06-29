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
    _merchantId   string
    // 请求ID
    _requestId   string
    // 业务类型。支付方式：payment_method
    _bizType   string
    // 经营设置json串
    _operationConfig   string
    // 操作员ID
    _outerOperatorId   string
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
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetMerchantId(_merchantId string) error {
    r._merchantId = _merchantId
    r.Set("merchant_id", _merchantId)
    return nil
}

// MerchantId Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetMerchantId() string {
    return r._merchantId
}
// RequestId Setter
// 请求ID
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetRequestId(_requestId string) error {
    r._requestId = _requestId
    r.Set("request_id", _requestId)
    return nil
}

// RequestId Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetRequestId() string {
    return r._requestId
}
// BizType Setter
// 业务类型。支付方式：payment_method
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetBizType() string {
    return r._bizType
}
// OperationConfig Setter
// 经营设置json串
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetOperationConfig(_operationConfig string) error {
    r._operationConfig = _operationConfig
    r.Set("operation_config", _operationConfig)
    return nil
}

// OperationConfig Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetOperationConfig() string {
    return r._operationConfig
}
// OuterOperatorId Setter
// 操作员ID
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetOuterOperatorId(_outerOperatorId string) error {
    r._outerOperatorId = _outerOperatorId
    r.Set("outer_operator_id", _outerOperatorId)
    return nil
}

// OuterOperatorId Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetOuterOperatorId() string {
    return r._outerOperatorId
}
