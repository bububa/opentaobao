package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrFulfillCancelReasonQueryAPIRequest
查询取消履约的原因列表 API请求
tmall.nr.fulfill.cancel.reason.query

新零售门店业务查询取消上门揽件的原因列表 */
type TmallNrFulfillCancelReasonQueryAPIRequest struct {
	model.Params
	// 商家的sellerID
	_sellerId int64
	// 淘宝交易的主订单号
	_mainOrderId int64
}

// New
