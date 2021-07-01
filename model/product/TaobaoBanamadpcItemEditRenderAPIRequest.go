package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBanamadpcItemEditRenderAPIRequest
编辑商品发布页 API请求
taobao.banamadpc.item.edit.render

巴拿马供应商通过此接口获取编辑商品发布页 */
type TaobaoBanamadpcItemEditRenderAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// New
