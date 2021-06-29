package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建MOZI自建人员和账号 APIResponse
alibaba.mozi.fusion.create.employee.account

创建MOZI自建人员和账号
*/
type AlibabaMoziFusionCreateEmployeeAccountAPIResponse struct {
    model.CommonResponse
    AlibabaMoziFusionCreateEmployeeAccountResponse
}

type AlibabaMoziFusionCreateEmployeeAccountResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_fusion_create_employee_account_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CreateTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
