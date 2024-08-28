package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingVersionCommit 提交版本号
// alibaba.hm.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
func AlibabaHmMarketingVersionCommit(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingVersionCommitAPIRequest, resp *wdk.AlibabaHmMarketingVersionCommitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
