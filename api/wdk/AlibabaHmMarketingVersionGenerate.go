package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingVersionGenerate 生成发布使用的版本号
// alibaba.hm.marketing.version.generate
//
// 生成发布使用的版本号
func AlibabaHmMarketingVersionGenerate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingVersionGenerateAPIRequest, resp *wdk.AlibabaHmMarketingVersionGenerateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
