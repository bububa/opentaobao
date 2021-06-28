package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单创建接口 APIResponse
cainiao.cboss.workplatform.workorder.create

菜鸟工单创建接口，目前调用者ISV
*/
type CainiaoCbossWorkplatformWorkorderCreateAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoCbossWorkplatformWorkorderCreateResponse `json:"cainiao_cboss_workplatform_workorder_create_response,omitempty"` 
    CainiaoCbossWorkplatformWorkorderCreateResponse
}

/* model for simplify = false
type CainiaoCbossWorkplatformWorkorderCreateResponse struct {

    // 接口返回model
    
    Result  *struct {
        CainiaoCbossWorkplatformWorkorderCreateResult  *CainiaoCbossWorkplatformWorkorderCreateResult `json:"cainiao_cboss_workplatform_workorder_create_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoCbossWorkplatformWorkorderCreateResponse struct {

    // 接口返回model
    Result   *CainiaoCbossWorkplatformWorkorderCreateResult `json:"result,omitempty"`

}
