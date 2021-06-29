package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
人员离职 APIResponse
alibaba.mozi.fusion.dimission.employee.account

人员离职并且回收账号
*/
type AlibabaMoziFusionDimissionEmployeeAccountAPIResponse struct {
    model.CommonResponse
    AlibabaMoziFusionDimissionEmployeeAccountResponse
}

type AlibabaMoziFusionDimissionEmployeeAccountResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_fusion_dimission_employee_account_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *RemoveTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
