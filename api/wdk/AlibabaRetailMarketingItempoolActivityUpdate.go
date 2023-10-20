package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolActivityUpdate 更新商品池活动【同城零售】
// alibaba.retail.marketing.itempool.activity.update
//
// 同城零售商品池活动更新
func AlibabaRetailMarketingItempoolActivityUpdate(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivityUpdateAPIRequest, session string) (*wdk.AlibabaRetailMarketingItempoolActivityUpdateAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingItempoolActivityUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
