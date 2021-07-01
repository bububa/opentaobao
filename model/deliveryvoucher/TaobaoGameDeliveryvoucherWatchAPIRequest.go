package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoGameDeliveryvoucherWatchAPIRequest
监控预约数据 API请求
taobao.game.deliveryvoucher.watch

监控预约数据 */
type TaobaoGameDeliveryvoucherWatchAPIRequest struct {
	model.Params
	// 入参
	_param0 *WatchAppointmentRequest
}

// New
