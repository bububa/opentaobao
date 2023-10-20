package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabanewretailpurchasepricesave 共享库存 采购价上传接口
// alibaba.newretail.purchase.price.save
//
// 共享库存业务 供应商上传商品采购价
func Alibabanewretailpurchasepricesave(clt *core.SDKClient, req *wdk.AlibabanewretailpurchasepricesaveAPIRequest, session string) (*wdk.AlibabanewretailpurchasepricesaveAPIResponse, error) {
	var resp wdk.AlibabanewretailpurchasepricesaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
