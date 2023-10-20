package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadgroupfindforbiddenproduct 查询屏蔽品
// alibaba.scbp.ad.group.find.forbidden.product
//
// 查询屏蔽品
func Alibabascbpadgroupfindforbiddenproduct(clt *core.SDKClient, req *scbp.AlibabascbpadgroupfindforbiddenproductAPIRequest, session string) (*scbp.AlibabascbpadgroupfindforbiddenproductAPIResponse, error) {
	var resp scbp.AlibabascbpadgroupfindforbiddenproductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
