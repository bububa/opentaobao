package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMallitemcenterSupplierPriceUpload 天猫服务商服务报价上传
// tmall.mallitemcenter.supplier.price.upload
//
// 天猫服务商上传服务价格
func TmallMallitemcenterSupplierPriceUpload(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallMallitemcenterSupplierPriceUploadAPIRequest, resp *tmallservice.TmallMallitemcenterSupplierPriceUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
