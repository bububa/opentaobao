package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolQueryactivity 查找商品池活动
// alibaba.hm.marketing.itempool.queryactivity
//
// 查找商品池活动
func AlibabaHmMarketingItempoolQueryactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolQueryactivityAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolQueryactivityAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolQueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
