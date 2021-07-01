package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsOrderCreateAPIRequest
创建物流订单 API请求
taobao.logistics.order.create

用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。 */
type TaobaoLogisticsOrderCreateAPIRequest struct {
	model.Params
	// 发货方式,默认为自己联系发货。可选值:ONLINE(在线下单)、OFFLINE(自己联系)。
	_logisType string
	// 发货的物流公司代码，如申通=STO，圆通=YTO。is_consign=true时，此项必填。
	_logisCompanyCode string
	// 发货的物流公司运单号。在logis_type=OFFLINE且is_consign=true时，此项必填。
	_mailNo string
	// 运送的货物名称列表，用|号隔开
	_goodsNames string
	// 运送货物的数量列表，用|号隔开
	_goodsQuantities string
	// 运送货物的单价列表(注意：单位为分），用|号隔开
	_itemValues string
	// 卖家旺旺号
	_sellerWangwangId string
	// 订单的交易号码
	_tradeId int64
	// 创建订单同时是否进行发货，默认发货。
	_isConsign bool
	// 运费承担方式。1为买家承担运费，2为卖家承担运费，其他值为错误参数。
	_shipping int64
}

// New
