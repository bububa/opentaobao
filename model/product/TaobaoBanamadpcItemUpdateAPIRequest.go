package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBanamadpcItemUpdateAPIRequest
编辑商品 API请求
taobao.banamadpc.item.update

巴拿马供应商通过此接口编辑商品 */
type TaobaoBanamadpcItemUpdateAPIRequest struct {
	model.Params
	// 商品的schema xml
	_xml string
	// 商品id
	_itemId int64
}

// New
