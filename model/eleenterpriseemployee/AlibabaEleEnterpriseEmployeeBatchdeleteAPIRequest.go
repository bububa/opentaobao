package eleenterpriseemployee

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest 批量删除员工 API请求
// alibaba.ele.enterprise.employee.batchdelete
//
// 批量删除员工
type AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest struct {
	model.Params
	// 员工工号
	_employeeNos []string
}

// NewAlibabaEleEnterpriseEmployeeBatchdeleteRequest 初始化AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest对象
func NewAlibabaEleEnterpriseEmployeeBatchdeleteRequest() *AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest {
	return &AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest) Reset() {
	r._employeeNos = r._employeeNos[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.employee.batchdelete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEmployeeNos is EmployeeNos Setter
// 员工工号
func (r *AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest) SetEmployeeNos(_employeeNos []string) error {
	r._employeeNos = _employeeNos
	r.Set("employee_nos", _employeeNos)
	return nil
}

// GetEmployeeNos EmployeeNos Getter
func (r AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest) GetEmployeeNos() []string {
	return r._employeeNos
}

var poolAlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseEmployeeBatchdeleteRequest()
	},
}

// GetAlibabaEleEnterpriseEmployeeBatchdeleteRequest 从 sync.Pool 获取 AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest
func GetAlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest() *AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest {
	return poolAlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest.Get().(*AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest)
}

// ReleaseAlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest 将 AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest(v *AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest.Put(v)
}
