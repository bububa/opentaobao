package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改员工基本信息 API请求
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

// 初始化AlibabaCampusCoreEmployeeModifyemployeeRequest对象
func NewAlibabaCampusCoreEmployeeModifyemployeeRequest() *AlibabaCampusCoreEmployeeModifyemployeeRequest{
    return &AlibabaCampusCoreEmployeeModifyemployeeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetApiMethodName() string {
    return "alibaba.campus.core.employee.modifyemployee"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkBenchContext Setter
// WorkBenchContext
func (r *AlibabaCampusCoreEmployeeModifyemployeeRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}
// EmployeeDto Setter
// EmployeeDto
func (r *AlibabaCampusCoreEmployeeModifyemployeeRequest) SetEmployeeDto(employeeDto *EmployeeDto) error {
    r.employeeDto = employeeDto
    r.Set("employee_dto", employeeDto)
    return nil
}

// EmployeeDto Getter
func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetEmployeeDto() *EmployeeDto {
    return r.employeeDto
}
// AccountId Setter
// 用户ID
func (r *AlibabaCampusCoreEmployeeModifyemployeeRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetAccountId() string {
    return r.accountId
}
