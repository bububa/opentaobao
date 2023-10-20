package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// Aliexpressaffiliateproductsmartmatch 联盟物料智能推荐api
// aliexpress.affiliate.product.smartmatch
//
// 联盟物料算法智能推荐
func Aliexpressaffiliateproductsmartmatch(clt *core.SDKClient, req *aecreatives.AliexpressaffiliateproductsmartmatchAPIRequest, session string) (*aecreatives.AliexpressaffiliateproductsmartmatchAPIResponse, error) {
	var resp aecreatives.AliexpressaffiliateproductsmartmatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
