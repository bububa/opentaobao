package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolActivityCreate 创建活动新接口
// alibaba.wdk.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
func AlibabaWdkMarketingItempoolActivityCreate(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolActivityCreateAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolActivityCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
