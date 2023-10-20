package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempoolqueryactivity 查找商品池活动
// alibaba.wdk.marketing.itempool.queryactivity
//
// 查找商品池活动
func Alibabawdkmarketingitempoolqueryactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempoolqueryactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingitempoolqueryactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempoolqueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
