package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加人员和账号复合接口 APIResponse
alibaba.mozi.fusion.addorupdate.employee.account

添加人员和账号复合接口
*/
type AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse struct {
    model.CommonResponse
    AlibabaMoziFusionAddorupdateEmployeeAccountResponse
}

type AlibabaMoziFusionAddorupdateEmployeeAccountResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_fusion_addorupdate_employee_account_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AddOrUpdateTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
