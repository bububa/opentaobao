package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolDeleteactivity 删除商品池活动
// alibaba.hm.marketing.itempool.deleteactivity
//
// 删除商品池活动
func AlibabaHmMarketingItempoolDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolDeleteactivityAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolDeleteactivityAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolDeleteactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
