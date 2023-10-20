package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupFindAdGroup 查询推广组
// alibaba.scbp.ad.group.find.ad.group
//
// 查询推广组
func AlibabaScbpAdGroupFindAdGroup(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupFindAdGroupAPIRequest, session string) (*scbp.AlibabaScbpAdGroupFindAdGroupAPIResponse, error) {
	var resp scbp.AlibabaScbpAdGroupFindAdGroupAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
