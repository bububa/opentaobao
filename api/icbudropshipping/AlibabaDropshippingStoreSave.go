package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaDropshippingStoreSave 阿里巴巴dropshipping店铺数据保存接口
// alibaba.dropshipping.store.save
//
// 阿里巴巴dropshipping店铺数据保存
func AlibabaDropshippingStoreSave(clt *core.SDKClient, req *icbudropshipping.AlibabaDropshippingStoreSaveAPIRequest, resp *icbudropshipping.AlibabaDropshippingStoreSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
