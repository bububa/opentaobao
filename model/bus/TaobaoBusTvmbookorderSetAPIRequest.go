package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusTvmbookorderSetAPIRequest
线下自助机通知出票接口 API请求
taobao.bus.tvmbookorder.set

出票，当成功的时候告知出票；当失败的时候告知出票失败，飞猪退款给用户。 */
type TaobaoBusTvmbookorderSetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
	// 出票时间 2017-03-03 11:22:33
	_bookTime string
	// true代表出票成功；false代表出票失败
	_success bool
	// 取值范围  ALIPAY （飞猪渠道）; WECHAT（微信）; BANKCARD（银行卡）;CASH（现金）; OWN_ALIPAY（自身支付宝渠道，非飞猪渠道）
	_payMode string
	// 取票人手机号
	_fetchPhone string
	// 乘客票面信息
	_passengers []TvmPassengerVo
	// 是否支持电子票
	_supportEticket bool
	// 检票口
	_ticketGate string
}

// New
