package icbudropshipping

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaDropshippingStoreSave 阿里巴巴dropshipping店铺数据保存接口
// alibaba.dropshipping.store.save
//
// 阿里巴巴dropshipping店铺数据保存
func AlibabaDropshippingStoreSave(ctx context.Context, clt *core.SDKClient, req *icbudropshipping.AlibabaDropshippingStoreSaveAPIRequest, resp *icbudropshipping.AlibabaDropshippingStoreSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
