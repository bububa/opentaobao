package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripEmployeeQuery 企业员工查询
// alitrip.btrip.employee.query
//
// 企业员工查询
func AlitripBtripEmployeeQuery(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripEmployeeQueryAPIRequest, resp *btrip.AlitripBtripEmployeeQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
