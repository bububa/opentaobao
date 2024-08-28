package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoProductImgUpload 上传单张产品非主图，如果需要传多张，可调多次
// taobao.product.img.upload
//
// 1.传入产品ID &lt;br/&gt;2.传入图片内容 &lt;br/&gt;注意：图片最大为500K,只支持JPG,GIF格式,如果需要传多张，可调多次
func TaobaoProductImgUpload(ctx context.Context, clt *core.SDKClient, req *product.TaobaoProductImgUploadAPIRequest, resp *product.TaobaoProductImgUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
