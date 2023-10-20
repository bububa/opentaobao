package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadgroupdeleteforbiddenproduct 删除屏蔽品
// alibaba.scbp.ad.group.delete.forbidden.product
//
// 删除屏蔽品
func Alibabascbpadgroupdeleteforbiddenproduct(clt *core.SDKClient, req *scbp.AlibabascbpadgroupdeleteforbiddenproductAPIRequest, session string) (*scbp.AlibabascbpadgroupdeleteforbiddenproductAPIResponse, error) {
	var resp scbp.AlibabascbpadgroupdeleteforbiddenproductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
