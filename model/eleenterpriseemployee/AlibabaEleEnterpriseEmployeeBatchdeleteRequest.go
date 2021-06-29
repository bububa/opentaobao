package eleenterpriseemployee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除员工 APIRequest
alibaba.ele.enterprise.employee.batchdelete

批量删除员工
*/
type AlibabaEleEnterpriseEmployeeBatchdeleteRequest struct {
    model.Params

    // 员工工号
    employeeNos   []string 

}

func NewAlibabaEleEnterpriseEmployeeBatchdeleteRequest() *AlibabaEleEnterpriseEmployeeBatchdeleteRequest{
    return &AlibabaEleEnterpriseEmployeeBatchdeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseEmployeeBatchdeleteRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.employee.batchdelete"
}

func (r AlibabaEleEnterpriseEmployeeBatchdeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseEmployeeBatchdeleteRequest) SetEmployeeNos(employeeNos []string) error {
    r.employeeNos = employeeNos
    r.Set("employee_nos", employeeNos)
    return nil
}

func (r AlibabaEleEnterpriseEmployeeBatchdeleteRequest) GetEmployeeNos() []string {
    return r.employeeNos
}

