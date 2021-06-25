package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家基础经营设置信息同步 APIRequest
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

func NewTaobaoKoubeiSaasBaseOperationConfigSyncRequest() *TaobaoKoubeiSaasBaseOperationConfigSyncRequest{
    return &TaobaoKoubeiSaasBaseOperationConfigSyncRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetApiMethodName() string {
    return "taobao.koubei.saas.base.operation.config.sync"
}

func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetMerchantId(merchantId string) error {
    r.merchantId = merchantId
    r.Set("merchant_id", merchantId)
    return nil
}

func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetMerchantId() string {
    return r.merchantId
}

func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetRequestId() string {
    return r.requestId
}

func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetBizType() string {
    return r.bizType
}

func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetOperationConfig(operationConfig string) error {
    r.operationConfig = operationConfig
    r.Set("operation_config", operationConfig)
    return nil
}

func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetOperationConfig() string {
    return r.operationConfig
}

func (r *TaobaoKoubeiSaasBaseOperationConfigSyncRequest) SetOuterOperatorId(outerOperatorId string) error {
    r.outerOperatorId = outerOperatorId
    r.Set("outer_operator_id", outerOperatorId)
    return nil
}

func (r TaobaoKoubeiSaasBaseOperationConfigSyncRequest) GetOuterOperatorId() string {
    return r.outerOperatorId
}

