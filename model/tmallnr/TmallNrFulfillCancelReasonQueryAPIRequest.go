package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillCancelReasonQueryAPIRequest 查询取消履约的原因列表 API请求
// tmall.nr.fulfill.cancel.reason.query
//
// 新零售门店业务查询取消上门揽件的原因列表
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrFulfillCancelReasonQueryAPIRequest) Reset() {
	r._sellerId = 0
	r._mainOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrFulfillCancelReasonQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nr.fulfill.cancel.reason.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrFulfillCancelReasonQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrFulfillCancelReasonQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerId is SellerId Setter
// 商家的sellerID
func (r *TmallNrFulfillCancelReasonQueryAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TmallNrFulfillCancelReasonQueryAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// SetMainOrderId is MainOrderId Setter
// 淘宝交易的主订单号
func (r *TmallNrFulfillCancelReasonQueryAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TmallNrFulfillCancelReasonQueryAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

var poolTmallNrFulfillCancelReasonQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrFulfillCancelReasonQueryRequest()
	},
}

// GetTmallNrFulfillCancelReasonQueryRequest 从 sync.Pool 获取 TmallNrFulfillCancelReasonQueryAPIRequest
func GetTmallNrFulfillCancelReasonQueryAPIRequest() *TmallNrFulfillCancelReasonQueryAPIRequest {
	return poolTmallNrFulfillCancelReasonQueryAPIRequest.Get().(*TmallNrFulfillCancelReasonQueryAPIRequest)
}

// ReleaseTmallNrFulfillCancelReasonQueryAPIRequest 将 TmallNrFulfillCancelReasonQueryAPIRequest 放入 sync.Pool
func ReleaseTmallNrFulfillCancelReasonQueryAPIRequest(v *TmallNrFulfillCancelReasonQueryAPIRequest) {
	v.Reset()
	poolTmallNrFulfillCancelReasonQueryAPIRequest.Put(v)
}
