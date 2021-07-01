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
type AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest struct {
    model.Params
    // 员工工号
    _employeeNos   []string
}

// 初始化AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest对象
func NewAlibabaEleEnterpriseEmployeeBatchdeleteRequest() *AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest{
    return &AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.employee.batchdelete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EmployeeNos Setter
// 员工工号
func (r *AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest) SetEmployeeNos(_employeeNos []string) error {
    r._employeeNos = _employeeNos
    r.Set("employee_nos", _employeeNos)
    return nil
}

// EmployeeNos Getter
func (r AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest) GetEmployeeNos() []string {
    return r._employeeNos
}
