package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolActivitySave 【同城零售】保存商品池活动
// alibaba.retail.marketing.itempool.activity.save
//
// 同城零售商品池活动保存
func AlibabaRetailMarketingItempoolActivitySave(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivitySaveAPIRequest, session string) (*wdk.AlibabaRetailMarketingItempoolActivitySaveAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingItempoolActivitySaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
