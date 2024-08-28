package eleenterpriseemployee

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseemployee"
)

// AlibabaEleEnterpriseEmployeeBatchdelete 批量删除员工
// alibaba.ele.enterprise.employee.batchdelete
//
// 批量删除员工
func AlibabaEleEnterpriseEmployeeBatchdelete(ctx context.Context, clt *core.SDKClient, req *eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchdeleteAPIRequest, resp *eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
