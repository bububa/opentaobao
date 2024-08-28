package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolActivityCreate 创建活动新接口
// alibaba.hm.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
func AlibabaHmMarketingItempoolActivityCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolActivityCreateAPIRequest, resp *wdk.AlibabaHmMarketingItempoolActivityCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
