package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBanamadpcItemAddAPIRequest
新发商品 API请求
taobao.banamadpc.item.add

巴拿马供应商通过此接口新发商品 */
type TaobaoBanamadpcItemAddAPIRequest struct {
	model.Params
	// 类目id
	_catId int64
	// 商品的schema xml
	_xml string
}

// New
