package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupRecommendProduct 推品
// alibaba.scbp.ad.group.recommend.product
//
// 推品
func AlibabaScbpAdGroupRecommendProduct(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupRecommendProductAPIRequest, resp *scbp.AlibabaScbpAdGroupRecommendProductAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
