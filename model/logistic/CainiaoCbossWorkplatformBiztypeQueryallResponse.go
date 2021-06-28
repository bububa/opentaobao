package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型 APIResponse
cainiao.cboss.workplatform.biztype.queryall

菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型。 目前调用者ISV
*/
type CainiaoCbossWorkplatformBiztypeQueryallAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoCbossWorkplatformBiztypeQueryallResponse `json:"cainiao_cboss_workplatform_biztype_queryall_response,omitempty"` 
    CainiaoCbossWorkplatformBiztypeQueryallResponse
}

/* model for simplify = false
type CainiaoCbossWorkplatformBiztypeQueryallResponse struct {

    // bizTypeJson
    
    BizTypeJson   string `json:"biz_type_json,omitempty"`
    

    // success
    
    ResSuccess   bool `json:"res_success,omitempty"`
    

    // errorCode
    
    ResErrorCode   string `json:"res_error_code,omitempty"`
    

    // errorMsg
    
    ResErrorMsg   string `json:"res_error_msg,omitempty"`
    

}
*/

type CainiaoCbossWorkplatformBiztypeQueryallResponse struct {

    // bizTypeJson
    BizTypeJson   string `json:"biz_type_json,omitempty"`

    // success
    ResSuccess   bool `json:"res_success,omitempty"`

    // errorCode
    ResErrorCode   string `json:"res_error_code,omitempty"`

    // errorMsg
    ResErrorMsg   string `json:"res_error_msg,omitempty"`

}
