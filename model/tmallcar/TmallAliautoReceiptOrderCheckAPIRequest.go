package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoReceiptOrderCheckAPIRequest 查看工单查询订单是否已付款 API请求
// tmall.aliauto.receipt.order.check
//
// 查看工单查询订单是否已付款
type TmallAliautoReceiptOrderCheckAPIRequest struct {
	model.Params
	// 服务商自定义门店编码
	_outerShopId string
	// 工单号
	_receiptId int64
}

// NewTmallAliautoReceiptOrderCheckRequest 初始化TmallAliautoReceiptOrderCheckAPIRequest对象
func NewTmallAliautoReceiptOrderCheckRequest() *TmallAliautoReceiptOrderCheckAPIRequest {
	return &TmallAliautoReceiptOrderCheckAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoReceiptOrderCheckAPIRequest) Reset() {
	r._outerShopId = ""
	r._receiptId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoReceiptOrderCheckAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.receipt.order.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoReceiptOrderCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoReceiptOrderCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterShopId is OuterShopId Setter
// 服务商自定义门店编码
func (r *TmallAliautoReceiptOrderCheckAPIRequest) SetOuterShopId(_outerShopId string) error {
	r._outerShopId = _outerShopId
	r.Set("outer_shop_id", _outerShopId)
	return nil
}

// GetOuterShopId OuterShopId Getter
func (r TmallAliautoReceiptOrderCheckAPIRequest) GetOuterShopId() string {
	return r._outerShopId
}

// SetReceiptId is ReceiptId Setter
// 工单号
func (r *TmallAliautoReceiptOrderCheckAPIRequest) SetReceiptId(_receiptId int64) error {
	r._receiptId = _receiptId
	r.Set("receipt_id", _receiptId)
	return nil
}

// GetReceiptId ReceiptId Getter
func (r TmallAliautoReceiptOrderCheckAPIRequest) GetReceiptId() int64 {
	return r._receiptId
}

var poolTmallAliautoReceiptOrderCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoReceiptOrderCheckRequest()
	},
}

// GetTmallAliautoReceiptOrderCheckRequest 从 sync.Pool 获取 TmallAliautoReceiptOrderCheckAPIRequest
func GetTmallAliautoReceiptOrderCheckAPIRequest() *TmallAliautoReceiptOrderCheckAPIRequest {
	return poolTmallAliautoReceiptOrderCheckAPIRequest.Get().(*TmallAliautoReceiptOrderCheckAPIRequest)
}

// ReleaseTmallAliautoReceiptOrderCheckAPIRequest 将 TmallAliautoReceiptOrderCheckAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoReceiptOrderCheckAPIRequest(v *TmallAliautoReceiptOrderCheckAPIRequest) {
	v.Reset()
	poolTmallAliautoReceiptOrderCheckAPIRequest.Put(v)
}
