package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopsecretbilldetailAPIRequest 服务商的商家解密账单详情查询 API请求
// taobao.top.secret.bill.detail
//
// 服务商的商家解密账单详情查询，仅对90天内的账单提供SLA保障。
type TaobaotopsecretbilldetailAPIRequest struct {
	model.Params
	// 卖家账单查询
	_sellerBillQueryRequest *SellerBillQueryRequest
}

// NewTaobaotopsecretbilldetailRequest 初始化TaobaotopsecretbilldetailAPIRequest对象
func NewTaobaotopsecretbilldetailRequest() *TaobaotopsecretbilldetailAPIRequest {
	return &TaobaotopsecretbilldetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopsecretbilldetailAPIRequest) GetApiMethodName() string {
	return "taobao.top.secret.bill.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopsecretbilldetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopsecretbilldetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerBillQueryRequest is SellerBillQueryRequest Setter
// 卖家账单查询
func (r *TaobaotopsecretbilldetailAPIRequest) SetSellerBillQueryRequest(_sellerBillQueryRequest *SellerBillQueryRequest) error {
	r._sellerBillQueryRequest = _sellerBillQueryRequest
	r.Set("seller_bill_query_request", _sellerBillQueryRequest)
	return nil
}

// GetSellerBillQueryRequest SellerBillQueryRequest Getter
func (r TaobaotopsecretbilldetailAPIRequest) GetSellerBillQueryRequest() *SellerBillQueryRequest {
	return r._sellerBillQueryRequest
}
