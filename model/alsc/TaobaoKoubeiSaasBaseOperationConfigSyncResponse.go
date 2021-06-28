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
	RequestId     string         `json:"request_id,omitempty" xml:"koubei_saas_base_operation_config_sync_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异常信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"