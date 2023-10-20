package eleenterpriseemployee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriseemployeebatchdeleteAPIRequest 批量删除员工 API请求
// alibaba.ele.enterprise.employee.batchdelete
//
// 批量删除员工
type AlibabaeleenterpriseemployeebatchdeleteAPIRequest struct {
	model.Params
	// 员工工号
	_employeeNos []string
}

// NewAlibabaeleenterpriseemployeebatchdeleteRequest 初始化AlibabaeleenterpriseemployeebatchdeleteAPIRequest对象
func NewAlibabaeleenterpriseemployeebatchdeleteRequest() *AlibabaeleenterpriseemployeebatchdeleteAPIRequest {
	return &AlibabaeleenterpriseemployeebatchdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterpriseemployeebatchdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.employee.batchdelete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterpriseemployeebatchdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterpriseemployeebatchdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEmployeeNos is EmployeeNos Setter
// 员工工号
func (r *AlibabaeleenterpriseemployeebatchdeleteAPIRequest) SetEmployeeNos(_employeeNos []string) error {
	r._employeeNos = _employeeNos
	r.Set("employee_nos", _employeeNos)
	return nil
}

// GetEmployeeNos EmployeeNos Getter
func (r AlibabaeleenterpriseemployeebatchdeleteAPIRequest) GetEmployeeNos() []string {
	return r._employeeNos
}
