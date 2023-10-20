package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyShopUpdateinfo 更新渠道店基础信息
// alibaba.aelophy.shop.updateinfo
//
// 更新渠道店基础信息
func AlibabaAelophyShopUpdateinfo(clt *core.SDKClient, req *wdk.AlibabaAelophyShopUpdateinfoAPIRequest, session string) (*wdk.AlibabaAelophyShopUpdateinfoAPIResponse, error) {
	var resp wdk.AlibabaAelophyShopUpdateinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
