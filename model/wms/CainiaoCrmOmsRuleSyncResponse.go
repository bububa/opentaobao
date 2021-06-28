package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家ERP订单处理规则同步 APIResponse
cainiao.crm.oms.rule.sync

将商家ERP订单处理规则同步到菜鸟CRM系统
*/
type CainiaoCrmOmsRuleSyncAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoCrmOmsRuleSyncResponse `json:"cainiao_crm_oms_rule_sync_response,omitempty"` 
    CainiaoCrmOmsRuleSyncResponse
}

/* model for simplify = false
type CainiaoCrmOmsRuleSyncResponse struct {

    // success
    
    WlSuccess   bool `json:"wl_success,omitempty"`
    

    // errorCode
    
    WlErrorCode   string `json:"wl_error_code,omitempty"`
    

    // errorMsg
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`
    

}
*/

type CainiaoCrmOmsRuleSyncResponse struct {

    // success
    WlSuccess   bool `json:"wl_success,omitempty"`

    // errorCode
    WlErrorCode   string `json:"wl_error_code,omitempty"`

    // errorMsg
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

}
