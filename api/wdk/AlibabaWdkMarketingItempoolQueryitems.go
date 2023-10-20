package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempoolqueryitems 查询商品池活动下的商品
// alibaba.wdk.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
func Alibabawdkmarketingitempoolqueryitems(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempoolqueryitemsAPIRequest, session string) (*wdk.AlibabawdkmarketingitempoolqueryitemsAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempoolqueryitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
