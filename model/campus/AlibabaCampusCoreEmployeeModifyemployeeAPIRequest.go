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
type AlibabaCampusCoreEmployeeModifyemployeeAPIRequest struct {
    model.Params
    // WorkBenchContext
    _workBenchContext   *WorkBenchContext
    // EmployeeDto
    _employeeDto   *EmployeeDto
    // 用户ID
    _accountId   string
}

// 初始化AlibabaCampusCoreEmployeeModifyemployeeAPIRequest对象
func NewAlibabaCampusCoreEmployeeModifyemployeeRequest() *AlibabaCampusCoreEmployeeModifyemployeeAPIRequest{
    return &AlibabaCampusCoreEmployeeModifyemployeeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusCoreEmployeeModifyemployeeAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.core.employee.modifyemployee"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusCoreEmployeeModifyemployeeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkBenchContext Setter
// WorkBenchContext
func (r *AlibabaCampusCoreEmployeeModifyemployeeAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
    r._workBenchContext = _workBenchContext
    r.Set("work_bench_context", _workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusCoreEmployeeModifyemployeeAPIRequest) GetWorkBenchContext() *WorkBenchContext {
    return r._workBenchContext
}
// EmployeeDto Setter
// EmployeeDto
func (r *AlibabaCampusCoreEmployeeModifyemployeeAPIRequest) SetEmployeeDto(_employeeDto *EmployeeDto) error {
    r._employeeDto = _employeeDto
    r.Set("employee_dto", _employeeDto)
    return nil
}

// EmployeeDto Getter
func (r AlibabaCampusCoreEmployeeModifyemployeeAPIRequest) GetEmployeeDto() *EmployeeDto {
    return r._employeeDto
}
// AccountId Setter
// 用户ID
func (r *AlibabaCampusCoreEmployeeModifyemployeeAPIRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusCoreEmployeeModifyemployeeAPIRequest) GetAccountId() string {
    return r._accountId
}
