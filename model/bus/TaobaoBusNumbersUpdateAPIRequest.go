package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusNumbersUpdateAPIRequest
汽车票车次更新服务 API请求
taobao.bus.numbers.update

用于汽车票车次信息的新增、更新和逻辑删除 */
type TaobaoBusNumbersUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramTopBusNumberUpdateRQ *TopBusNumberUpdateRq
}

// New
