package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Alibabalsycrmactivityupdate ISV活动修改
// alibaba.lsy.crm.activity.update
//
// ISV活动修改
func Alibabalsycrmactivityupdate(clt *core.SDKClient, req *tmallnr.AlibabalsycrmactivityupdateAPIRequest, session string) (*tmallnr.AlibabalsycrmactivityupdateAPIResponse, error) {
	var resp tmallnr.AlibabalsycrmactivityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
