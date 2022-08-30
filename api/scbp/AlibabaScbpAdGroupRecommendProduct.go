package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupRecommendProduct 推品
// alibaba.scbp.ad.group.recommend.product
//
// 推品
func AlibabaScbpAdGroupRecommendProduct(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupRecommendProductAPIRequest, session string) (*scbp.AlibabaScbpAdGroupRecommendProductAPIResponse, error) {
	var resp scbp.AlibabaScbpAdGroupRecommendProductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
