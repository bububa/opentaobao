package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempooldeleteactivity 删除商品池活动
// alibaba.wdk.marketing.itempool.deleteactivity
//
// 删除商品池活动
func Alibabawdkmarketingitempooldeleteactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempooldeleteactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingitempooldeleteactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempooldeleteactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
