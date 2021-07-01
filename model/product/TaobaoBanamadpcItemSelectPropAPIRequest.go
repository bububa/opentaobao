package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBanamadpcItemSelectPropAPIRequest
获取子属性 API请求
taobao.banamadpc.item.select.prop

巴拿马供应商通过此接口获取子属性 */
type TaobaoBanamadpcItemSelectPropAPIRequest struct {
	model.Params
	// 子属性的schema xml
	_xml string
	// 属性id
	_propId int64
	// 类目id
	_catId int64
}

// New
