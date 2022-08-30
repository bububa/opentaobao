package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaDropshippingStoreSave 阿里巴巴dropshipping店铺数据保存接口
// alibaba.dropshipping.store.save
//
// 阿里巴巴dropshipping店铺数据保存
func AlibabaDropshippingStoreSave(clt *core.SDKClient, req *icbudropshipping.AlibabaDropshippingStoreSaveAPIRequest, session string) (*icbudropshipping.AlibabaDropshippingStoreSaveAPIResponse, error) {
	var resp icbudropshipping.AlibabaDropshippingStoreSaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
