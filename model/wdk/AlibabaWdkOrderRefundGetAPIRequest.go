package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderRefundGetAPIRequest 五道口订单退款按ID查询 API请求
// alibaba.wdk.order.refund.get
//
// 按照退款ID或者五道口中台订单ID查询退款信息详情
type AlibabaWdkOrderRefundGetAPIRequest struct {
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

// NewAlibabaWdkOrderRefundGetRequest 初始化AlibabaWdkOrderRefundGetAPIRequest对象
func NewAlibabaWdkOrderRefundGetRequest() *AlibabaWdkOrderRefundGetAPIRequest {
	return &AlibabaWdkOrderRefundGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkOrderRefundGetAPIRequest) Reset() {
	r._bizOrderIds = r._bizOrderIds[:0]
	r._refundIds = r._refundIds[:0]
	r._shopId = ""
	r._storeId = ""
	r._orderFrom = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderRefundGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderRefundGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkOrderRefundGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderIds is BizOrderIds Setter
// 五道口订单列表（子订单）
func (r *AlibabaWdkOrderRefundGetAPIRequest) SetBizOrderIds(_bizOrderIds []int64) error {
	r._bizOrderIds = _bizOrderIds
	r.Set("biz_order_ids", _bizOrderIds)
	return nil
}

// GetBizOrderIds BizOrderIds Getter
func (r AlibabaWdkOrderRefundGetAPIRequest) GetBizOrderIds() []int64 {
	return r._bizOrderIds
}

// SetRefundIds is RefundIds Setter
// 退款订单列表
func (r *AlibabaWdkOrderRefundGetAPIRequest) SetRefundIds(_refundIds []int64) error {
	r._refundIds = _refundIds
	r.Set("refund_ids", _refundIds)
	return nil
}

// GetRefundIds RefundIds Getter
func (r AlibabaWdkOrderRefundGetAPIRequest) GetRefundIds() []int64 {
	return r._refundIds
}

// SetShopId is ShopId Setter
// 渠道店id
func (r *AlibabaWdkOrderRefundGetAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabaWdkOrderRefundGetAPIRequest) GetShopId() string {
	return r._shopId
}

// SetStoreId is StoreId Setter
// 经营店id
func (r *AlibabaWdkOrderRefundGetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaWdkOrderRefundGetAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetOrderFrom is OrderFrom Setter
// 渠道来源 3：饿了么 4：盒马
func (r *AlibabaWdkOrderRefundGetAPIRequest) SetOrderFrom(_orderFrom int64) error {
	r._orderFrom = _orderFrom
	r.Set("order_from", _orderFrom)
	return nil
}

// GetOrderFrom OrderFrom Getter
func (r AlibabaWdkOrderRefundGetAPIRequest) GetOrderFrom() int64 {
	return r._orderFrom
}

var poolAlibabaWdkOrderRefundGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkOrderRefundGetRequest()
	},
}

// GetAlibabaWdkOrderRefundGetRequest 从 sync.Pool 获取 AlibabaWdkOrderRefundGetAPIRequest
func GetAlibabaWdkOrderRefundGetAPIRequest() *AlibabaWdkOrderRefundGetAPIRequest {
	return poolAlibabaWdkOrderRefundGetAPIRequest.Get().(*AlibabaWdkOrderRefundGetAPIRequest)
}

// ReleaseAlibabaWdkOrderRefundGetAPIRequest 将 AlibabaWdkOrderRefundGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkOrderRefundGetAPIRequest(v *AlibabaWdkOrderRefundGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkOrderRefundGetAPIRequest.Put(v)
}
