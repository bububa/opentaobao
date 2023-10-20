package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadgroupcountadgroup 统计adgroup数量
// alibaba.scbp.ad.group.count.ad.group
//
// 统计adgroup数量
func Alibabascbpadgroupcountadgroup(clt *core.SDKClient, req *scbp.AlibabascbpadgroupcountadgroupAPIRequest, session string) (*scbp.AlibabascbpadgroupcountadgroupAPIResponse, error) {
	var resp scbp.AlibabascbpadgroupcountadgroupAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
