package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmerchantitemquery 商家商品查询
// alibaba.wdk.merchant.item.query
//
// 商家商品查询
func Alibabawdkmerchantitemquery(clt *core.SDKClient, req *wdk.AlibabawdkmerchantitemqueryAPIRequest, session string) (*wdk.AlibabawdkmerchantitemqueryAPIResponse, error) {
	var resp wdk.AlibabawdkmerchantitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
