package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductImageDelete 产品图片删除
// taobao.fenxiao.product.image.delete
//
// 产品图片删除，只删除图片信息，不真正删除图片
func TaobaoFenxiaoProductImageDelete(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductImageDeleteAPIRequest, resp *fenxiao.TaobaoFenxiaoProductImageDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
