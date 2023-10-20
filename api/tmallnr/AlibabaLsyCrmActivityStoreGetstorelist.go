package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Alibabalsycrmactivitystoregetstorelist ISV查询门店
// alibaba.lsy.crm.activity.store.getstorelist
//
// ISV查询门店
func Alibabalsycrmactivitystoregetstorelist(clt *core.SDKClient, req *tmallnr.AlibabalsycrmactivitystoregetstorelistAPIRequest, session string) (*tmallnr.AlibabalsycrmactivitystoregetstorelistAPIResponse, error) {
	var resp tmallnr.AlibabalsycrmactivitystoregetstorelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
