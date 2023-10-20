package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolExcludeskucode 商品池排除商品【品类优惠使用】
// alibaba.hm.marketing.itempool.excludeskucode
//
// 品类优惠新增排除池
func AlibabaHmMarketingItempoolExcludeskucode(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolExcludeskucodeAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolExcludeskucodeAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolExcludeskucodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
