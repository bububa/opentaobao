package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadgroupfindadgroup 查询推广组
// alibaba.scbp.ad.group.find.ad.group
//
// 查询推广组
func Alibabascbpadgroupfindadgroup(clt *core.SDKClient, req *scbp.AlibabascbpadgroupfindadgroupAPIRequest, session string) (*scbp.AlibabascbpadgroupfindadgroupAPIResponse, error) {
	var resp scbp.AlibabascbpadgroupfindadgroupAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
