package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductImageDelete 产品图片删除
// taobao.fenxiao.product.image.delete
//
// 产品图片删除，只删除图片信息，不真正删除图片
func TaobaoFenxiaoProductImageDelete(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductImageDeleteAPIRequest, resp *fenxiao.TaobaoFenxiaoProductImageDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
