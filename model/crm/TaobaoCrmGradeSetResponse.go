package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家设置等级规则 APIResponse
taobao.crm.grade.set

设置等级信息，可以设置层级等级，也可以单独设置一个等级。出于安全原因，折扣现最低只能设置到700即7折。
*/
type TaobaoCrmGradeSetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCrmGradeSetResponse `json:"crm_grade_set_response,omitempty"` 
    TaobaoCrmGradeSetResponse
}

/* model for simplify = false
type TaobaoCrmGradeSetResponse struct {

    // true：成功 false：失败
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoCrmGradeSetResponse struct {

    // true：成功 false：失败
    IsSuccess   bool `json:"is_success,omitempty"`

}
