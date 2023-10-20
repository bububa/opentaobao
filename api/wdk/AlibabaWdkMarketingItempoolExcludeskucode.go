package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempoolexcludeskucode 商品池排除商品【品类优惠使用】
// alibaba.wdk.marketing.itempool.excludeskucode
//
// 品类优惠新增排除池
func Alibabawdkmarketingitempoolexcludeskucode(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempoolexcludeskucodeAPIRequest, session string) (*wdk.AlibabawdkmarketingitempoolexcludeskucodeAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempoolexcludeskucodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
