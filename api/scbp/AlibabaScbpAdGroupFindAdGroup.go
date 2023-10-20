package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupFindAdGroup 查询推广组
// alibaba.scbp.ad.group.find.ad.group
//
// 查询推广组
func AlibabaScbpAdGroupFindAdGroup(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupFindAdGroupAPIRequest, resp *scbp.AlibabaScbpAdGroupFindAdGroupAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
