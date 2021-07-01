package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusLastplaceGetAPIRequest
获取目的地数据 API请求
taobao.bus.lastplace.get

传入城市 获取对应的目的地 */
type TaobaoBusLastplaceGetAPIRequest struct {
	model.Params
	// 目的地查询参数
	_paramLastPlaceSearchRQ *ParamLastPlaceSearchRq
}

// New
