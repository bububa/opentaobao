package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改员工基本信息 APIResponse
alibaba.campus.core.employee.modifyemployee

根据用户ID和公司ID更新员工基本信息（头像、性别、昵称）
*/
type AlibabaCampusCoreEmployeeModifyemployeeAPIResponse struct {
    model.CommonResponse
    AlibabaCampusCoreEmployeeModifyemployeeResponse
}

type AlibabaCampusCoreEmployeeModifyemployeeResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_core_employee_modifyemployee_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求响应
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
