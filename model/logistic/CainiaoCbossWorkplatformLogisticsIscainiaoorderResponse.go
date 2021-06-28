package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据交易单号判断是否为菜鸟发货订单 APIResponse
cainiao.cboss.workplatform.logistics.iscainiaoorder

根据交易单号判断是否为菜鸟发货订单
*/
type CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoCbossWorkplatformLogisticsIscainiaoorderResponse `json:"cainiao_cboss_workplatform_logistics_iscainiaoorder_response,omitempty"` 
    CainiaoCbossWorkplatformLogisticsIscainiaoorderResponse
}

/* model for simplify = false
type CainiaoCbossWorkplatformLogisticsIscainiaoorderResponse struct {

    // isCainiaoOrder
    
    IsCainiaoOrder   bool `json:"is_cainiao_order,omitempty"`
    

    // success
    
    ResSuccess   bool `json:"res_success,omitempty"`
    

    // errorCode
    
    ResErrorCode   string `json:"res_error_code,omitempty"`
    

    // errorMsg
    
    ResErrorMsg   string `json:"res_error_msg,omitempty"`
    

}
*/

type CainiaoCbossWorkplatformLogisticsIscainiaoorderResponse struct {

    // isCainiaoOrder
    IsCainiaoOrder   bool `json:"is_cainiao_order,omitempty"`

    // success
    ResSuccess   bool `json:"res_success,omitempty"`

    // errorCode
    ResErrorCode   string `json:"res_error_code,omitempty"`

    // errorMsg
    ResErrorMsg   string `json:"res_error_msg,omitempty"`

}
