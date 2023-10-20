package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempoolcreateactivity 添加商品池活动
// alibaba.wdk.marketing.itempool.createactivity
//
// 添加商品池活动
func Alibabawdkmarketingitempoolcreateactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempoolcreateactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingitempoolcreateactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempoolcreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
