package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolCreateactivity 添加商品池活动
// alibaba.wdk.marketing.itempool.createactivity
//
// 添加商品池活动
func AlibabaWdkMarketingItempoolCreateactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolCreateactivityAPIRequest, session string) (*wdk.AlibabaWdkMarketingItempoolCreateactivityAPIResponse, error) {
	var resp wdk.AlibabaWdkMarketingItempoolCreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
