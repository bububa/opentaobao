package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Alibabalsycrmactivitygetbaseinfo ISV查询活动
// alibaba.lsy.crm.activity.getbaseinfo
//
// ISV查询活动
func Alibabalsycrmactivitygetbaseinfo(clt *core.SDKClient, req *tmallnr.AlibabalsycrmactivitygetbaseinfoAPIRequest, session string) (*tmallnr.AlibabalsycrmactivitygetbaseinfoAPIResponse, error) {
	var resp tmallnr.AlibabalsycrmactivitygetbaseinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
