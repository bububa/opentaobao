package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkitemmerchantcategoryquery 查询商品的商家叶子类目
// alibaba.wdk.item.merchant.category.query
//
// 查询商品的商家叶子类目
func Alibabawdkitemmerchantcategoryquery(clt *core.SDKClient, req *wdk.AlibabawdkitemmerchantcategoryqueryAPIRequest, session string) (*wdk.AlibabawdkitemmerchantcategoryqueryAPIResponse, error) {
	var resp wdk.AlibabawdkitemmerchantcategoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
