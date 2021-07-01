package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeDrugOrdersGetAPIRequest
阿里健康获取某一药店全部订单 API请求
taobao.trade.drug.orders.get

阿里健康获取某一药店全部订单 */
type TaobaoTradeDrugOrdersGetAPIRequest struct {
	model.Params
	// 外卖分店ID
	_shopId int64
	// 关键字
	_keyword string
	// true-查询仅按商家维度  false-查询按商家下所属店铺维度
	_isAllShop bool
	// true 仅有支付宝订单,false 包括所有类型订单(货到付款,支付券等)
	_isAllOrder bool
	// （必填字段）订单状态 待确认订单2 , 退款中订单4 , 已发货12 关闭20 交易成功21
	_orderStatus int64
	// 返回记录数，超过20按20条返回数据
	_pageSize int64
	// 页码
	_pageNo int64
}

// New
