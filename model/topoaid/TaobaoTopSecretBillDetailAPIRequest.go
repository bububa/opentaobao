package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopSecretBillDetailAPIRequest 服务商的商家解密账单详情查询 API请求
// taobao.top.secret.bill.detail
//
// 服务商的商家解密账单详情查询，仅对90天内的账单提供SLA保障。
type TaobaoTopSecretBillDetailAPIRequest struct {
	model.Params
	// 卖家账单查询
	_sellerBillQueryRequest *SellerBillQueryRequest
}

// NewTaobaoTopSecretBillDetailRequest 初始化TaobaoTopSecretBillDetailAPIRequest对象
func NewTaobaoTopSecretBillDetailRequest() *TaobaoTopSecretBillDetailAPIRequest {
	return &TaobaoTopSecretBillDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopSecretBillDetailAPIRequest) GetApiMethodName() string {
	return "taobao.top.secret.bill.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopSecretBillDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSellerBillQueryRequest is SellerBillQueryRequest Setter
// 卖家账单查询
func (r *TaobaoTopSecretBillDetailAPIRequest) SetSellerBillQueryRequest(_sellerBillQueryRequest *SellerBillQueryRequest) error {
	r._sellerBillQueryRequest = _sellerBillQueryRequest
	r.Set("seller_bill_query_request", _sellerBillQueryRequest)
	return nil
}

// GetSellerBillQueryRequest SellerBillQueryRequest Getter
func (r TaobaoTopSecretBillDetailAPIRequest) GetSellerBillQueryRequest() *SellerBillQueryRequest {
	return r._sellerBillQueryRequest
}
