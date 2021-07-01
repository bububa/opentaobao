package larkiot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLarkIotOrderConfirmorderAPIRequest
iot渠道卖品落单 API请求
taobao.lark.iot.order.confirmorder

云智对接无人超市，接收无人超市订单信息 */
type TaobaoLarkIotOrderConfirmorderAPIRequest struct {
	model.Params
	// 渠道编码
	_channelCode string
	// 影院内码
	_cinemaLinkId string
	// 外部订单号
	_outGoodsOrderId string
	// 工作站id
	_workstationId string
	// 工作站名称
	_workstationName string
	// 支付方式
	_paymentList string
	// 优惠列表
	_promotionList string
	// 卖品列表
	_goodsList string
	// 手机号
	_mobile string
	// 管理员
	_operatorUserId string
}

// New
