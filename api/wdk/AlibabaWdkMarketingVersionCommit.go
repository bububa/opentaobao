package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingVersionCommit 提交版本号
// alibaba.wdk.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
func AlibabaWdkMarketingVersionCommit(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingVersionCommitAPIRequest, resp *wdk.AlibabaWdkMarketingVersionCommitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
