package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempoolqueryitems 查询商品池活动下面的商品
// alibaba.hm.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
func Alibabahmmarketingitempoolqueryitems(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempoolqueryitemsAPIRequest, session string) (*wdk.AlibabahmmarketingitempoolqueryitemsAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempoolqueryitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
