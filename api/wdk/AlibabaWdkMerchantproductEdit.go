package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmerchantproductedit 商家产品服务-编辑产品
// alibaba.wdk.merchantproduct.edit
//
// 商家产品服务-编辑产品
func Alibabawdkmerchantproductedit(clt *core.SDKClient, req *wdk.AlibabawdkmerchantproducteditAPIRequest, session string) (*wdk.AlibabawdkmerchantproducteditAPIResponse, error) {
	var resp wdk.AlibabawdkmerchantproducteditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
