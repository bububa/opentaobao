package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallProductSpecPicUpload 上传产品规格认证图片
// tmall.product.spec.pic.upload
//
// 上传指定类型的产品规格认证文件，并返回存有上传成功图片url的产品规格对象
func TmallProductSpecPicUpload(ctx context.Context, clt *core.SDKClient, req *product.TmallProductSpecPicUploadAPIRequest, resp *product.TmallProductSpecPicUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
