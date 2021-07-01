package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusTvmcreateorderSetAPIRequest
线下自助机创建订单 API请求
taobao.bus.tvmcreateorder.set

提供给汽车票线下自助机的创建订单使用 */
type TaobaoBusTvmcreateorderSetAPIRequest struct {
	model.Params
	// 创建订单对象
	_paramTVMCreateOrderRQ *TvmCreateOrderRq
}

// New
