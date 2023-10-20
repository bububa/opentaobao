package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadgrouprecommendproduct 推品
// alibaba.scbp.ad.group.recommend.product
//
// 推品
func Alibabascbpadgrouprecommendproduct(clt *core.SDKClient, req *scbp.AlibabascbpadgrouprecommendproductAPIRequest, session string) (*scbp.AlibabascbpadgrouprecommendproductAPIResponse, error) {
	var resp scbp.AlibabascbpadgrouprecommendproductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
