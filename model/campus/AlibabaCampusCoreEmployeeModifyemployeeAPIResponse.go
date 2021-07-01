package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改员工基本信息 API返回值 
alibaba.campus.core.employee.modifyemployee

根据用户ID和公司ID更新员工基本信息（头像、性别、昵称）
*/
type AlibabaCampusCoreEmployeeModifyemployeeAPIResponse struct {
    model.CommonResponse
    AlibabaCampusCoreEmployeeModifyemployeeAPIResponseModel
}

// 修改员工基本信息 成功返回结果
type AlibabaCampusCoreEmployeeModifyemployeeAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_campus_core_employee_modifyemployee_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求响应
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
