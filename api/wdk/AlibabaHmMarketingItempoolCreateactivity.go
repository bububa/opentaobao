package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolCreateactivity 添加商品池活动
// alibaba.hm.marketing.itempool.createactivity
//
// 添加商品池活动
func AlibabaHmMarketingItempoolCreateactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolCreateactivityAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolCreateactivityAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolCreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
