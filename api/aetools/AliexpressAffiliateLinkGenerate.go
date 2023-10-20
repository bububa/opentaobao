package aetools

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aetools"
)

// Aliexpressaffiliatelinkgenerate 联盟推广链接生成
// aliexpress.affiliate.link.generate
//
// AE联盟推广链接生成接口
func Aliexpressaffiliatelinkgenerate(clt *core.SDKClient, req *aetools.AliexpressaffiliatelinkgenerateAPIRequest, session string) (*aetools.AliexpressaffiliatelinkgenerateAPIResponse, error) {
	var resp aetools.AliexpressaffiliatelinkgenerateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
