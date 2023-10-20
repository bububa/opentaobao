package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Alibabalsycrmcustomerappoint ISV直播间预约
// alibaba.lsy.crm.customer.appoint
//
// ISV直播间预约
func Alibabalsycrmcustomerappoint(clt *core.SDKClient, req *tmallnr.AlibabalsycrmcustomerappointAPIRequest, session string) (*tmallnr.AlibabalsycrmcustomerappointAPIResponse, error) {
	var resp tmallnr.AlibabalsycrmcustomerappointAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
