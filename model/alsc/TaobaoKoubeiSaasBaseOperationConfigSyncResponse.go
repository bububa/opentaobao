package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家基础经营设置信息同步 APIResponse
taobao.koubei.saas.base.operation.config.sync

ISV接入口碑SAAS后, 经营设置数据同步到口碑SAAS
*/
type TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiSaasBaseOperationConfigSyncResponse
}

type TaobaoKoubeiSaasBaseOperationConfigSyncResponse struct {
    XMLName xml.Name `xml:"koubei_saas_base_operation_config_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异常信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 异常码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 是否成功
    
    BizSuccess   bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`

    
}
