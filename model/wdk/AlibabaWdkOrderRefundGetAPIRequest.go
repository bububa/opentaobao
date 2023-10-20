package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkorderrefundgetAPIRequest 五道口订单退款按ID查询 API请求
// alibaba.wdk.order.refund.get
//
// 按照退款ID或者五道口中台订单ID查询退款信息详情
type AlibabawdkorderrefundgetAPIRequest struct {
	model.Params
	// 五道口订单列表（子订单）
	_bizOrderIds []int64
	// 退款订单列表
	_refundIds []int64
	// 渠道店id
	_shopId string
	// 经营店id
	_storeId string
	// 渠道来源 3：饿了么 4：盒马
	_orderFrom int64
}

// NewAlibabawdkorderrefundgetRequest 初始化AlibabawdkorderrefundgetAPIRequest对象
func NewAlibabawdkorderrefundgetRequest() *AlibabawdkorderrefundgetAPIRequest {
	return &AlibabawdkorderrefundgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkorderrefundgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkorderrefundgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkorderrefundgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderIds is BizOrderIds Setter
// 五道口订单列表（子订单）
func (r *AlibabawdkorderrefundgetAPIRequest) SetBizOrderIds(_bizOrderIds []int64) error {
	r._bizOrderIds = _bizOrderIds
	r.Set("biz_order_ids", _bizOrderIds)
	return nil
}

// GetBizOrderIds BizOrderIds Getter
func (r AlibabawdkorderrefundgetAPIRequest) GetBizOrderIds() []int64 {
	return r._bizOrderIds
}

// SetRefundIds is RefundIds Setter
// 退款订单列表
func (r *AlibabawdkorderrefundgetAPIRequest) SetRefundIds(_refundIds []int64) error {
	r._refundIds = _refundIds
	r.Set("refund_ids", _refundIds)
	return nil
}

// GetRefundIds RefundIds Getter
func (r AlibabawdkorderrefundgetAPIRequest) GetRefundIds() []int64 {
	return r._refundIds
}

// SetShopId is ShopId Setter
// 渠道店id
func (r *AlibabawdkorderrefundgetAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabawdkorderrefundgetAPIRequest) GetShopId() string {
	return r._shopId
}

// SetStoreId is StoreId Setter
// 经营店id
func (r *AlibabawdkorderrefundgetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabawdkorderrefundgetAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetOrderFrom is OrderFrom Setter
// 渠道来源 3：饿了么 4：盒马
func (r *AlibabawdkorderrefundgetAPIRequest) SetOrderFrom(_orderFrom int64) error {
	r._orderFrom = _orderFrom
	r.Set("order_from", _orderFrom)
	return nil
}

// GetOrderFrom OrderFrom Getter
func (r AlibabawdkorderrefundgetAPIRequest) GetOrderFrom() int64 {
	return r._orderFrom
}
