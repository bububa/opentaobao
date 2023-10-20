package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabanewretailpurchasepricedelete 共享库存 商户删除采购价
// alibaba.newretail.purchase.price.delete
//
// 共享库存 商户删除采购价
func Alibabanewretailpurchasepricedelete(clt *core.SDKClient, req *wdk.AlibabanewretailpurchasepricedeleteAPIRequest, session string) (*wdk.AlibabanewretailpurchasepricedeleteAPIResponse, error) {
	var resp wdk.AlibabanewretailpurchasepricedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
