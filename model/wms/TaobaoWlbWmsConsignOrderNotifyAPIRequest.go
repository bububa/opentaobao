package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsConsignOrderNotifyAPIRequest
发货订单通知 API请求
taobao.wlb.wms.consign.order.notify

发货订单通知 */
type TaobaoWlbWmsConsignOrderNotifyAPIRequest struct {
	model.Params
	// ERP订单号
	_orderCode string
	// 单据类型 201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单
	_orderType int64
	// 订单标识 (1: cod –货到付款，4:invoiceinfo-需要发票)
	_orderFlag string
	// 订单来源（213 天猫，201 淘宝，214 京东，202 1688 阿里中文站 ，203 苏宁在线，204 亚马逊中国，205 当当，208 1号店，207 唯品会，209 国美在线，210 拍拍，206 易贝ebay，211 聚美优品，212 乐蜂网，215 邮乐，216 凡客，217 优购，218 银泰，219 易讯，221 聚尚网，222 蘑菇街，223 POS门店，301 其他）
	_orderSource int64
	// 仓库编码，此字段为空时，由菜鸟路由仓库发货
	_storeCode string
	// 快递公司编码，此字段为空时，由菜鸟选择快递配送
	_tmsServiceCode string
	// 快递公司名称
	_tmsServiceName string
	// 前物流订单号，订单类型为502 换货出库单 503 补发出库单时，需求传入此内容
	_prevOrderCode string
	// 下单时间，订单在交易平台创建时间
	_orderShopCreateTime string
	// 订单支付时间
	_orderPayTime string
	// 订单创建时间,ERP创建订单时间
	_orderCreateTime string
	// 订单审核时间,ERP创建支付时间
	_orderExaminationTime string
	// 订单总金额,=总商品金额-订单优惠金额+快递费用，单位分
	_orderAmount int64
	// 订单优惠金额，整单优惠金额，单位分
	_discountAmount int64
	// 订单应收金额，消费者还需要付的金额，单位分
	_arAmount int64
	// 订单已付金额，消费者已经支付的金额，单位分
	_gotAmount int64
	// 快递费用，单位分
	_postfee int64
	// COD服务费，单位分
	_serviceFee int64
	// 配送要求
	_deliverRequirements *Deliverrequirementswlbwmsconsignordernotify
	// 收件人信息
	_receiverInfo *Receiverwlbwmsconsignordernotify
	// 发货方信息
	_senderInfo *Senderwlbwmsconsignordernotify
	// 订单商品信息列表
	_orderItemList []Orderitemlistwlbwmsconsignordernotify
	// 发票信息列表
	_invoiceInfoList []Invoicelistwlbwmsconsignordernotify
	// 拓展属性
	_extendFields string
	// 备注
	_remark string
}

// New
