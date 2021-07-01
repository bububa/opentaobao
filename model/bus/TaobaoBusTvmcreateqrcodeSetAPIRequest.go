package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusTvmcreateqrcodeSetAPIRequest
自助机生成支付宝支付二维码 API请求
taobao.bus.tvmcreateqrcode.set

用于汽车票线下自助机调用获取支付宝的二维码 */
type TaobaoBusTvmcreateqrcodeSetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
	// 超时时间（分钟）
	_timeoutExpress int64
}

// New
