package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseOrdernewCreateAPIRequest
创建订单 API请求
alibaba.ele.enterprise.ordernew.create

创建订单 */
type AlibabaEleEnterpriseOrdernewCreateAPIRequest struct {
	model.Params
	// 订单来源地址经度
	_longitude string
	// 餐厅Id
	_erestaurantId string
	// 使用的券号
	_couponSn string
	// 订单备注信息
	_description string
	// 电话号码，主要号码必须是手机号；多个手机号以逗号分隔
	_phones string
	// 订单来源IP地址
	_ip string
	// 订单来源地址纬度
	_latitude string
	// 购物车Id（创建购物车返回的购物车id）
	_cartId string
	// 第三方订单Id（需保证唯一）
	_tpOrderId string
	// 送餐地址
	_address string
	// 收餐人姓名
	_consignee string
	// 暂时不用传（忽略此字段）
	_deliverTime string
	// 纳税人识别号
	_invoiceNumber string
	// 发票抬头（个人发票请填写个人），不传表示不要发票
	_invoice string
	// 发票类型（发票类型, 1: 个人, 2: 企业; 空为兼容数据, 由商户判断发票类型）
	_invoiceType int64
}

// New
