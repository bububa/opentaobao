package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmCustomerAppoint ISV直播间预约
// alibaba.lsy.crm.customer.appoint
//
// ISV直播间预约
func AlibabaLsyCrmCustomerAppoint(ctx context.Context, clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmCustomerAppointAPIRequest, resp *tmallnr.AlibabaLsyCrmCustomerAppointAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
