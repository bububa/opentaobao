package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// Aliexpressaffiliateproductquery 联盟推广商品获取接口
// aliexpress.affiliate.product.query
//
// 联盟推广商品搜索接口，用于搜索联盟推广商品数据
func Aliexpressaffiliateproductquery(clt *core.SDKClient, req *aecreatives.AliexpressaffiliateproductqueryAPIRequest, session string) (*aecreatives.AliexpressaffiliateproductqueryAPIResponse, error) {
	var resp aecreatives.AliexpressaffiliateproductqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
