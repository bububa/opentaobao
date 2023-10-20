package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadgroupcreateforbiddenproduct 创建屏蔽品
// alibaba.scbp.ad.group.create.forbidden.product
//
// 创建屏蔽品
func Alibabascbpadgroupcreateforbiddenproduct(clt *core.SDKClient, req *scbp.AlibabascbpadgroupcreateforbiddenproductAPIRequest, session string) (*scbp.AlibabascbpadgroupcreateforbiddenproductAPIResponse, error) {
	var resp scbp.AlibabascbpadgroupcreateforbiddenproductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
