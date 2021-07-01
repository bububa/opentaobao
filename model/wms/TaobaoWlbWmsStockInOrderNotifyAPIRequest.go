package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsStockInOrderNotifyAPIRequest
入库通知单 API请求
taobao.wlb.wms.stock.in.order.notify

入库通知单 */
type TaobaoWlbWmsStockInOrderNotifyAPIRequest struct {
	model.Params
	// 仓库编码
	_storeCode string
	// 入库单据编码
	_orderCode string
	// 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、306 B2B入库
	_orderType int64
	// 可选择性文本透传至WMS，比如加工归还、委外归还、借出归还、内部归还等
	_inboundTypeDesc string
	// 订单标记以逗号分隔：  9:上门退货入库 13: 退货时是否收取发票，默认不收取（即没13为多选项，如1,2,8,9）
	_orderFlag string
	// 单据创建时间
	_orderCreateTime string
	// 供应商编码，往来单位编码
	_supplierCode string
	// 供应商名称 ，往来单位名称
	_supplierName string
	// 配送公司编码，拒收（非妥投）的销退订单，由ERP填充原单配送公司编码
	_tmsServiceCode string
	// 快递公司名称
	_tmsServiceName string
	// 运单号&托运单号 1)	入库单支持多运单号录入 2)	销退场景下如果是拒收（非妥投运单）由ERP填充原运单号
	_tmsOrderCode string
	// 来源单据号，如销售退货时填充原销售订单号
	_prevOrderCode string
	// 销退时请提供退货的原因
	_returnReason string
	// 预期送达开始时间
	_expectStartTime string
	// 预期送达结束时间
	_expectEndTime string
	// 系统自动生成
	_receiverInfo *Receiverinfowlbwmsstockinordernotifywl
	// 系统自动生成
	_senderInfo *Senderinfowlbwmsstockinordernotifywl
	// 系统自动生成
	_orderItemList []Orderitemlistwlbwmsstockinordernotifywl
	// 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
	_extendFields string
	// 备注，销退入库订单需要留言备注填充到此字段
	_remark string
}

// New
