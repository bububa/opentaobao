package aetools

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aetools"
)

// AliexpressAffiliateLinkGenerate 联盟推广链接生成
// aliexpress.affiliate.link.generate
//
// AE联盟推广链接生成接口
func AliexpressAffiliateLinkGenerate(clt *core.SDKClient, req *aetools.AliexpressAffiliateLinkGenerateAPIRequest, session string) (*aetools.AliexpressAffiliateLinkGenerateAPIResponse, error) {
	var resp aetools.AliexpressAffiliateLinkGenerateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
