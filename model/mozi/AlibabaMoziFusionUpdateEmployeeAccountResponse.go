package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新人员和账号属性 APIResponse
alibaba.mozi.fusion.update.employee.account

更新人员和账号基本属性
*/
type AlibabaMoziFusionUpdateEmployeeAccountAPIResponse struct {
    model.CommonResponse
    AlibabaMoziFusionUpdateEmployeeAccountResponse
}

type AlibabaMoziFusionUpdateEmployeeAccountResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_fusion_update_employee_account_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *UpdateTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
