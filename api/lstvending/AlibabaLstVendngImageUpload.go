package lstvending

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

// AlibabaLstVendngImageUpload 售货机商品图片上传
// alibaba.lst.vendng.image.upload
//
// 零售通自动售货机商品图片上传接口，主要为ISV厂商提供图片同步的通道，从而建立统一的商品图片库。
func AlibabaLstVendngImageUpload(ctx context.Context, clt *core.SDKClient, req *lstvending.AlibabaLstVendngImageUploadAPIRequest, resp *lstvending.AlibabaLstVendngImageUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
