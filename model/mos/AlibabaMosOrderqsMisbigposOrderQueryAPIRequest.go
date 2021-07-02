package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOrderqsMisbigposOrderQueryAPIRequest 大pos新选单退 API请求
// alibaba.mos.orderqs.misbigpos.order.query
//
// 大pos新选单退
type AlibabaMosOrderqsMisbigposOrderQueryAPIRequest struct {
	model.Params
	// 外部门店号
	_storeNo string
	// 基本信息获取参数
	_queryBaseData bool
	// 小票号
	_receiptNo string
	// 券扩展数据获取
	_queryCouponExtern bool
}

// NewAlibabaMosOrderqsMisbigposOrderQueryRequest 初始化AlibabaMosOrderqsMisbigposOrderQueryAPIRequest对象
func NewAlibabaMosOrderqsMisbigposOrderQueryRequest() *AlibabaMosOrderqsMisbigposOrderQueryAPIRequest {
	return &AlibabaMosOrderqsMisbigposOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosOrderqsMisbigposOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.orderqs.misbigpos.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosOrderqsMisbigposOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreNo Setter
// 外部门店号
func (r *AlibabaMosOrderqsMisbigposOrderQueryAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// Get StoreNo Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// Set is QueryBaseData Setter
// 基本信息获取参数
func (r *AlibabaMosOrderqsMisbigposOrderQueryAPIRequest) SetQueryBaseData(_queryBaseData bool) error {
	r._queryBaseData = _queryBaseData
	r.Set("query_base_data", _queryBaseData)
	return nil
}

// Get QueryBaseData Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryAPIRequest) GetQueryBaseData() bool {
	return r._queryBaseData
}

// Set is ReceiptNo Setter
// 小票号
func (r *AlibabaMosOrderqsMisbigposOrderQueryAPIRequest) SetReceiptNo(_receiptNo string) error {
	r._receiptNo = _receiptNo
	r.Set("receipt_no", _receiptNo)
	return nil
}

// Get ReceiptNo Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryAPIRequest) GetReceiptNo() string {
	return r._receiptNo
}

// Set is QueryCouponExtern Setter
// 券扩展数据获取
func (r *AlibabaMosOrderqsMisbigposOrderQueryAPIRequest) SetQueryCouponExtern(_queryCouponExtern bool) error {
	r._queryCouponExtern = _queryCouponExtern
	r.Set("query_coupon_extern", _queryCouponExtern)
	return nil
}

// Get QueryCouponExtern Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryAPIRequest) GetQueryCouponExtern() bool {
	return r._queryCouponExtern
}
