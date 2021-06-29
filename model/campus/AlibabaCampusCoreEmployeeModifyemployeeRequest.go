package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改员工基本信息 APIRequest
alibaba.campus.core.employee.modifyemployee

根据用户ID和公司ID更新员工基本信息（头像、性别、昵称）
*/
type AlibabaCampusCoreEmployeeModifyemployeeRequest struct {
    model.Params

    // WorkBenchContext
    workBenchContext   *WorkBenchContext 

    // EmployeeDto
    employeeDto   *EmployeeDto 

    // 用户ID
    accountId   string 

}

func NewAlibabaCampusCoreEmployeeModifyemployeeRequest() *AlibabaCampusCoreEmployeeModifyemployeeRequest{
    return &AlibabaCampusCoreEmployeeModifyemployeeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetApiMethodName() string {
    return "alibaba.campus.core.employee.modifyemployee"
}

func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusCoreEmployeeModifyemployeeRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}

func (r *AlibabaCampusCoreEmployeeModifyemployeeRequest) SetEmployeeDto(employeeDto *EmployeeDto) error {
    r.employeeDto = employeeDto
    r.Set("employee_dto", employeeDto)
    return nil
}

func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetEmployeeDto() *EmployeeDto {
    return r.employeeDto
}

func (r *AlibabaCampusCoreEmployeeModifyemployeeRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetAccountId() string {
    return r.accountId
}

