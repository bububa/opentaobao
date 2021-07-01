package cityretail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCityretailWmflConvertWarehouseAPIRequest
同城零售完美履约转仓 API请求
taobao.cityretail.wmfl.convert.warehouse

同城零售完美履约转仓 */
type TaobaoCityretailWmflConvertWarehouseAPIRequest struct {
	model.Params
	// 淘宝交易单id
	_tbOrderId string
}

// New
