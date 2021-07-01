package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradePostageUpdateAPIRequest
修改交易邮费价格 API请求
taobao.trade.postage.update

修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
<br/> <span style="color:red"> API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750</span> */
type TaobaoTradePostageUpdateAPIRequest struct {
	model.Params
	// 主订单编号
	_tid int64
	// 邮费价格(邮费单位是元）
	_postFee string
}

// New
