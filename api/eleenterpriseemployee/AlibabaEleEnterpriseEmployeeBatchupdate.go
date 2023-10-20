package eleenterpriseemployee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseemployee"
)

// AlibabaEleEnterpriseEmployeeBatchupdate 批量新增更新员工
// alibaba.ele.enterprise.employee.batchupdate
//
// 批量新增更新员工
func AlibabaEleEnterpriseEmployeeBatchupdate(clt *core.SDKClient, req *eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest, resp *eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
