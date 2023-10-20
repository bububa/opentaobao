package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempoolstairqueryitem 换购商品查询
// alibaba.hm.marketing.itempool.stair.queryitem
//
// 换购商品查询
func Alibabahmmarketingitempoolstairqueryitem(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempoolstairqueryitemAPIRequest, session string) (*wdk.AlibabahmmarketingitempoolstairqueryitemAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempoolstairqueryitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
