package dmp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dmp"
)

// TaobaoDmpCrowdBasicFind DMP_BP版人群列表查询
// taobao.dmp.crowd.basic.find
//
// DMP_BP版人群列表查询
func TaobaoDmpCrowdBasicFind(clt *core.SDKClient, req *dmp.TaobaoDmpCrowdBasicFindAPIRequest, resp *dmp.TaobaoDmpCrowdBasicFindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
