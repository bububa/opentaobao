package eleenterpriseemployee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseemployee"
)

// AlibabaEleEnterpriseEmployeeBatchdelete 批量删除员工
// alibaba.ele.enterprise.employee.batchdelete
//
// 批量删除员工
func AlibabaEleEnterpriseEmployeeBatchdelete(clt *core.SDKClient, req *eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest, session string) (*eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse, error) {
	var resp eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
