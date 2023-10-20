package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// Aliexpressaffiliateproductdetailget 联盟商品详情获取接口
// aliexpress.affiliate.productdetail.get
//
// 联盟推广商品搜索接口，用于搜索联盟推广商品数据
func Aliexpressaffiliateproductdetailget(clt *core.SDKClient, req *aecreatives.AliexpressaffiliateproductdetailgetAPIRequest, session string) (*aecreatives.AliexpressaffiliateproductdetailgetAPIResponse, error) {
	var resp aecreatives.AliexpressaffiliateproductdetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
