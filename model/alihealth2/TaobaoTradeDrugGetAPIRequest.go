package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeDrugGetAPIRequest
查询商家未确认订单列表 API请求
taobao.trade.drug.get

可以按商家或是店铺维度的进行查询买家付款卖家未确认订单，一次返回不大于20条订单 */
type TaobaoTradeDrugGetAPIRequest struct {
	model.Params
	// 店铺id
	_storeId int64
	// true-商家下所有店铺的待确认订单, false—指定店铺的订单
	_isAll bool
	// 返回记录数，超过20按20条返回数据
	_maxSize int64
}

// New
