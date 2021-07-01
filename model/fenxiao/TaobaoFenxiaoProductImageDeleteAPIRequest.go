package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductImageDeleteAPIRequest
产品图片删除 API请求
taobao.fenxiao.product.image.delete

产品图片删除，只删除图片信息，不真正删除图片 */
type TaobaoFenxiaoProductImageDeleteAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 图片位置
	_position int64
	// properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
	_properties string
}

// New
