package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempoolstairqueryitem 换购商品查询
// alibaba.wdk.marketing.itempool.stair.queryitem
//
// 换购商品查询
func Alibabawdkmarketingitempoolstairqueryitem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempoolstairqueryitemAPIRequest, session string) (*wdk.AlibabawdkmarketingitempoolstairqueryitemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempoolstairqueryitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
