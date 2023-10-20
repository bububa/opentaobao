package aetools

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aetools"
)

// AliexpressAffiliateLinkGenerate 联盟推广链接生成
// aliexpress.affiliate.link.generate
//
// AE联盟推广链接生成接口
func AliexpressAffiliateLinkGenerate(clt *core.SDKClient, req *aetools.AliexpressAffiliateLinkGenerateAPIRequest, resp *aetools.AliexpressAffiliateLinkGenerateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
