package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaNewretailPurchasePriceSave
共享库存 采购价上传接口
alibaba.newretail.purchase.price.save

共享库存业务 供应商上传商品采购价 */
func AlibabaNewretailPurchasePriceSave(clt *core.SDKClient, req *wdk.AlibabaNewretailPurchasePriceSaveAPIRequest, session string) (*wdk.AlibabaNewretailPurchasePriceSaveAPIResponse, error) {
	var resp wdk.AlibabaNewretailPurchasePriceSaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
