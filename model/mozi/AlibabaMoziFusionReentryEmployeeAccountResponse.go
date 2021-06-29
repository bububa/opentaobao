package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
重新入职并且重新启用账号 APIResponse
alibaba.mozi.fusion.reentry.employee.account

重新入职并且重新启用账号
*/
type AlibabaMoziFusionReentryEmployeeAccountAPIResponse struct {
    model.CommonResponse
    AlibabaMoziFusionReentryEmployeeAccountResponse
}

type AlibabaMoziFusionReentryEmployeeAccountResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_fusion_reentry_employee_account_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ReEntryTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
