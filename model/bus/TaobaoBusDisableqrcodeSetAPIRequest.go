package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusDisableqrcodeSetAPIRequest
自助机失效二维码 API请求
taobao.bus.disableqrcode.set

使创建的二维码失效 */
type TaobaoBusDisableqrcodeSetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
}

// New
