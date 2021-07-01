package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsStockOutOrderNotifyAPIRequest
出库单通知 API请求
taobao.wlb.wms.stock.out.order.notify

出库单通知 */
type TaobaoWlbWmsStockOutOrderNotifyAPIRequest struct {
	model.Params
	// 仓储编码
	_storeCode string
	// ERP单据ID
	_orderCode string
	// 单据类型 301 调拨出库单、901普通出库单、903 其他出库单 305 B2B出库
	_orderType int64
	// ERP可选择性文本透传至WMS
	_outboundTypeDesc string
	// 订单创建时间
	_orderCreateTime string
	// 要求出库日期
	_sendTime string
	// 收件人信息
	_receiverInfo *Receiverwlbwmsstockoutordernotify
	// 发货方信息
	_senderInfo *Senderwlbwmsstockoutordernotify
	// 出库方式
	_transportMode string
	// 承运商名称
	_carriersName string
	// 取件人姓名
	_pickName string
	// 取件人电话
	_pickCall string
	// 取件人身份证ID
	_pickId string
	// 车牌号
	_carNo string
	// 订单商品信息列表
	_orderItemList []Orderitemlistwlbwmsstockoutordernotify
	// 备注
	_remark string
	// 前物流订单号，如退货入库单可能会用到
	_prevOrderCode string
	// 拓展属性
	_extendFields string
}

// New
