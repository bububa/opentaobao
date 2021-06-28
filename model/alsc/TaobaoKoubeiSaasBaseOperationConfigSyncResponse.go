package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家基础经营设置信息同步 APIResponse
taobao.koubei.saas.base.operation.config.sync

ISV接入口碑SAAS后, 经营设置数据同步到口碑SAAS
*/
type TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoKoubeiSaasBaseOperationConfigSyncResponse `json:"koubei_saas_base_operation_config_sync_response,omitempty"` 
    TaobaoKoubeiSaasBaseOperationConfigSyncResponse
}

/* model for simplify = false
type TaobaoKoubeiSaasBaseOperationConfigSyncResponse struct {

    // 异常信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 异常码
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // 是否成功
    
    BizSuccess   bool `json:"biz_success,omitempty"`
    

}
*/

type TaobaoKoubeiSaasBaseOperationConfigSyncResponse struct {

    // 异常信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 异常码
    MsgCode   string `json:"msg_code,omitempty"`

    // 是否成功
    BizSuccess   bool `json:"biz_success,omitempty"`

}
