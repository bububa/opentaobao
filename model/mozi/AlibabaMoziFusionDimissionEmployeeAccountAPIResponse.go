package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
人员离职 API返回值 
alibaba.mozi.fusion.dimission.employee.account

人员离职并且回收账号
*/
type AlibabaMoziFusionDimissionEmployeeAccountAPIResponse struct {
    model.CommonResponse
    AlibabaMoziFusionDimissionEmployeeAccountAPIResponseModel
}

// 人员离职 成功返回结果
type AlibabaMoziFusionDimissionEmployeeAccountAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mozi_fusion_dimission_employee_account_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *RemoveTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}
