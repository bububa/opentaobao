package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsReturnOrderNotifyAPIRequest
销售退货通知 API请求
taobao.wlb.wms.return.order.notify

销售退货通知 */
type TaobaoWlbWmsReturnOrderNotifyAPIRequest struct {
	model.Params
	// ERP单据编码
	_orderCode string
	// 仓库编码
	_storeCode string
	// 用字符串格式来表示订单标记列表：比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL 等，中间用“^”来隔开 ----------------------------------------  订单标记list（所有字母全部大写）： 9:VISIT-上门12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
	_orderFlag string
	// 来源单据号，销售退货时填充原发货的LBX号
	_prevOrderCode string
	// 快递公司编码
	_tmsServiceCode string
	// 运单号
	_tmsOrderCode string
	// 销退时请提供退货的原因
	_returnReason string
	// 买家昵称
	_buyerNick string
	// 发件人信息
	_senderInfo *Senderinfowlbwmsreturnordernotify
	// 收件人信息
	_receiverInfo *Receiverinfowlbwmsreturnordernotify
	// 商品信息列表
	_orderItemList []Orderitemlistwlbwmsreturnordernotify
	// 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
	_extendFields string
	// 备注
	_remark string
	// 订单来源 201 淘宝，202 1688，203 苏宁，204 亚马逊中国，205 当当，206 ebay，207 唯品会，208 1号店，209 国美，210 拍拍，211 聚美优品，212 乐峰，214 京东，301 其他
	_orderSource string
	// 订单类型 501 销退入库
	_orderType string
	// 货主ID
	_ownerUserId string
	// ERP订单创建时间
	_orderCreateTime string
	// 快递公司名称
	_tmsServiceName string
}

// New
