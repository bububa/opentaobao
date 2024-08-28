package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolActivityCreate 创建活动新接口
// alibaba.wdk.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
func AlibabaWdkMarketingItempoolActivityCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolActivityCreateAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolActivityCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
