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
    _workBenchContext   *WorkBenchContext
    // EmployeeDto
    _employeeDto   *EmployeeDTO
    // 用户ID
    _accountId   string
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
func (r *AlibabaCampusCoreEmployeeModifyemployeeRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
    r._workBenchContext = _workBenchContext
    r.Set("work_bench_context", _workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetWorkBenchContext() *WorkBenchContext {
    return r._workBenchContext
}
// EmployeeDto Setter
// EmployeeDto
func (r *AlibabaCampusCoreEmployeeModifyemployeeRequest) SetEmployeeDto(_employeeDto *EmployeeDTO) error {
    r._employeeDto = _employeeDto
    r.Set("employee_dto", _employeeDto)
    return nil
}

// EmployeeDto Getter
func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetEmployeeDto() *EmployeeDTO {
    return r._employeeDto
}
// AccountId Setter
// 用户ID
func (r *AlibabaCampusCoreEmployeeModifyemployeeRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusCoreEmployeeModifyemployeeRequest) GetAccountId() string {
    return r._accountId
}
