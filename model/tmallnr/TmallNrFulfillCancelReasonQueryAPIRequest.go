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

// NewTmallNrFulfillCancelReasonQueryRequest 初始化TmallNrFulfillCancelReasonQueryAPIRequest对象
func NewTmallNrFulfillCancelReasonQueryRequest() *TmallNrFulfillCancelReasonQueryAPIRequest {
	return &TmallNrFulfillCancelReasonQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrFulfillCancelReasonQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nr.fulfill.cancel.reason.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrFulfillCancelReasonQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SellerId Setter
// 商家的sellerID
func (r *TmallNrFulfillCancelReasonQueryAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// Get SellerId Getter
func (r TmallNrFulfillCancelReasonQueryAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// Set is MainOrderId Setter
// 淘宝交易的主订单号
func (r *TmallNrFulfillCancelReasonQueryAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// Get MainOrderId Getter
func (r TmallNrFulfillCancelReasonQueryAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}
