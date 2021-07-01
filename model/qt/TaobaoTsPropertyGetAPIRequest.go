package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTsPropertyGetAPIRequest
淘宝服务属性查询 API请求
taobao.ts.property.get

淘宝服务属性查询 */
type TaobaoTsPropertyGetAPIRequest struct {
	model.Params
	// 服务收费项code
	_serviceItemCode string
}

// New
