package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingVersionGenerate 生成发布使用的版本号
// alibaba.hm.marketing.version.generate
//
// 生成发布使用的版本号
func AlibabaHmMarketingVersionGenerate(clt *core.SDKClient, req *wdk.AlibabaHmMarketingVersionGenerateAPIRequest, session string) (*wdk.AlibabaHmMarketingVersionGenerateAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingVersionGenerateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
