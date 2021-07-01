package eleenterpriseemployee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest
批量新增更新员工 API请求
alibaba.ele.enterprise.employee.batchupdate

批量新增更新员工 */
type AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest struct {
	model.Params
	// 批量员工信息
	_enterpriseDatas []EmployeeInfoDto
}

// New
