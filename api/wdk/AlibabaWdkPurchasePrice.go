package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkpurchaseprice rt回传采购价
// alibaba.wdk.purchase.price
//
// 猫超共享库存项目-rt回传采购价
func Alibabawdkpurchaseprice(clt *core.SDKClient, req *wdk.AlibabawdkpurchasepriceAPIRequest, session string) (*wdk.AlibabawdkpurchasepriceAPIResponse, error) {
	var resp wdk.AlibabawdkpurchasepriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
