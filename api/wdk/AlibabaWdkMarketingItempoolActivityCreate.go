package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempoolactivitycreate 创建活动新接口
// alibaba.wdk.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
func Alibabawdkmarketingitempoolactivitycreate(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempoolactivitycreateAPIRequest, session string) (*wdk.AlibabawdkmarketingitempoolactivitycreateAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempoolactivitycreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
