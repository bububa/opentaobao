package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmCustomerAppoint ISV直播间预约
// alibaba.lsy.crm.customer.appoint
//
// ISV直播间预约
func AlibabaLsyCrmCustomerAppoint(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmCustomerAppointAPIRequest, session string) (*tmallnr.AlibabaLsyCrmCustomerAppointAPIResponse, error) {
	var resp tmallnr.AlibabaLsyCrmCustomerAppointAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
