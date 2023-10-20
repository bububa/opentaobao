package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingVersionGenerate 生成发布使用的版本号
// alibaba.hm.marketing.version.generate
//
// 生成发布使用的版本号
func AlibabaHmMarketingVersionGenerate(clt *core.SDKClient, req *wdk.AlibabaHmMarketingVersionGenerateAPIRequest, resp *wdk.AlibabaHmMarketingVersionGenerateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
