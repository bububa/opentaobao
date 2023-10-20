package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempoolexcludeskucode 商品池排除商品【品类优惠使用】
// alibaba.hm.marketing.itempool.excludeskucode
//
// 品类优惠新增排除池
func Alibabahmmarketingitempoolexcludeskucode(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempoolexcludeskucodeAPIRequest, session string) (*wdk.AlibabahmmarketingitempoolexcludeskucodeAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempoolexcludeskucodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
