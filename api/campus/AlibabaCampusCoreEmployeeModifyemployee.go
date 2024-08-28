package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusCoreEmployeeModifyemployee 修改员工基本信息
// alibaba.campus.core.employee.modifyemployee
//
// 根据用户ID和公司ID更新员工基本信息（头像、性别、昵称）
func AlibabaCampusCoreEmployeeModifyemployee(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusCoreEmployeeModifyemployeeAPIRequest, resp *campus.AlibabaCampusCoreEmployeeModifyemployeeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
