package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBanamadpcItemRenderAPIRequest
新发商品发布页 API请求
taobao.banamadpc.item.render

巴拿马供应商通过此接口新发商品发布页 */
type TaobaoBanamadpcItemRenderAPIRequest struct {
	model.Params
	// 类目ID
	_catId int64
}

// New
