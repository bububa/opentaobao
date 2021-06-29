package eleenterpriseemployee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除员工 API请求
alibaba.ele.enterprise.employee.batchdelete

批量删除员工
*/
type AlibabaEleEnterpriseEmployeeBatchdeleteRequest struct {
    model.Params
    // 员工工号
    _employeeNos   []string
}

// 初始化AlibabaEleEnterpriseEmployeeBatchdeleteRequest对象
func NewAlibabaEleEnterpriseEmployeeBatchdeleteRequest() *AlibabaEleEnterpriseEmployeeBatchdeleteRequest{
    return &AlibabaEleEnterpriseEmployeeBatchdeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseEmployeeBatchdeleteRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.employee.batchdelete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseEmployeeBatchdeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EmployeeNos Setter
// 员工工号
func (r *AlibabaEleEnterpriseEmployeeBatchdeleteRequest) SetEmployeeNos(_employeeNos []string) error {
    r._employeeNos = _employeeNos
    r.Set("employee_nos", _employeeNos)
    return nil
}

// EmployeeNos Getter
func (r AlibabaEleEnterpriseEmployeeBatchdeleteRequest) GetEmployeeNos() []string {
    return r._employeeNos
}
