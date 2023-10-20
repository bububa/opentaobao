package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmCustomerAppoint ISV直播间预约
// alibaba.lsy.crm.customer.appoint
//
// ISV直播间预约
func AlibabaLsyCrmCustomerAppoint(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmCustomerAppointAPIRequest, resp *tmallnr.AlibabaLsyCrmCustomerAppointAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
