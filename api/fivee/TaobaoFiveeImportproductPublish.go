package fivee

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// TaobaoFiveeImportproductPublish 进口商品发布
// taobao.fivee.importproduct.publish
//
// 直营业务商家入住发布商品时，上传商品及商家证照信息
func TaobaoFiveeImportproductPublish(ctx context.Context, clt *core.SDKClient, req *fivee.TaobaoFiveeImportproductPublishAPIRequest, resp *fivee.TaobaoFiveeImportproductPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
