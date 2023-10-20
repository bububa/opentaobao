package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Alibabalsycrmactivitypageupdate ISV活动页面创建修改
// alibaba.lsy.crm.activity.page.update
//
// ISV活动页面创建修改
func Alibabalsycrmactivitypageupdate(clt *core.SDKClient, req *tmallnr.AlibabalsycrmactivitypageupdateAPIRequest, session string) (*tmallnr.AlibabalsycrmactivitypageupdateAPIResponse, error) {
	var resp tmallnr.AlibabalsycrmactivitypageupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
