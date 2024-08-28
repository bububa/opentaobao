package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingVersionGenerate 生成发布使用的版本号
// alibaba.wdk.marketing.version.generate
//
// 生成发布使用的版本号
func AlibabaWdkMarketingVersionGenerate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingVersionGenerateAPIRequest, resp *wdk.AlibabaWdkMarketingVersionGenerateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
