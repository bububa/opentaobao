package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

/* AlibabaAilabUserTokenGet
三方账号获取 token
alibaba.ailab.user.token.get

inside 设备的三方 app，通过 extId、schema 生成 token */
func AlibabaAilabUserTokenGet(clt *core.SDKClient, req *alilabs.AlibabaAilabUserTokenGetAPIRequest, session string) (*alilabs.AlibabaAilabUserTokenGetAPIResponse, error) {
	var resp alilabs.AlibabaAilabUserTokenGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
