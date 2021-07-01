package lstfundbill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest
结算明细数据查询（卖家视角） API请求
alibaba.lst.trade.seller.order.fundbill.list.query

提供For供应商的结算接口，以交易账单维度提供开放数据，满足供应商自动化结算的诉求 */
type AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest struct {
	model.Params
	// 每页最大主订单数，注意：返回的content_list数据按照子订单维度展开
	_size int64
	// 账单日期，格式：yyyy-MM-dd
	_billDate string
	// 页码
	_page int64
}

// New
