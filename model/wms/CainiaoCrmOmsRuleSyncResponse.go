package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家ERP订单处理规则同步 APIResponse
cainiao.crm.oms.rule.sync

将商家ERP订单处理规则同步到菜鸟CRM系统
*/
type CainiaoCrmOmsRuleSyncAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_crm_oms_rule_sync_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // success
    
    WlSuccess   bool `json:"wl_success,omitempty" xml:"