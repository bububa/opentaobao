package eleenterpriseemployee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest
批量删除员工 API请求
alibaba.ele.enterprise.employee.batchdelete

批量删除员工 */
type AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest struct {
	model.Params
	// 员工工号
	_employeeNos []string
}

// New
