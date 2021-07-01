package lstfundbill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeOrderFundbillQueryAPIRequest
结算明细数据查询（品牌商视角） API请求
alibaba.lst.trade.order.fundbill.query

按照指定日期提供交易账单维度的结算明细数据，比供应商工作台上的结算账单还多一些数据项。 */
type AlibabaLstTradeOrderFundbillQueryAPIRequest struct {
	model.Params
	// 每页最大记录数
	_size int64
	// 账单日期，格式：yyyy-MM-dd
	_billDate string
	// 页码
	_page int64
	// 为true时,返回相应的商品详细信息，item_id和unit
	_needItemDetail bool
}

// New
