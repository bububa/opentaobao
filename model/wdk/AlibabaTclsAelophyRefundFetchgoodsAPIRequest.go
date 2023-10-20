package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophyrefundfetchgoodsAPIRequest saas 售后逆向 商户发起逆向取货 API请求
// alibaba.tcls.aelophy.refund.fetchgoods
//
// saas 售后逆向 商户发起逆向取货
type AlibabatclsaelophyrefundfetchgoodsAPIRequest struct {
	model.Params
	// 经营店ID
	_storeId string
	// 外部订单ID
	_outOrderId string
	// 渠道退款单ID
	_refundId string
	// 取货开始时间
	_fetchStartTime string
	// 取货结束时间
	_fetchEndTime string
	// 备注
	_remark string
	// 外部子订单列表
	_subRefundList *Subrefundlist
	// 渠道来源
	_orderFrom int64
}

// NewAlibabatclsaelophyrefundfetchgoodsRequest 初始化AlibabatclsaelophyrefundfetchgoodsAPIRequest对象
func NewAlibabatclsaelophyrefundfetchgoodsRequest() *AlibabatclsaelophyrefundfetchgoodsAPIRequest {
	return &AlibabatclsaelophyrefundfetchgoodsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophyrefundfetchgoodsAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.fetchgoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophyrefundfetchgoodsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophyrefundfetchgoodsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 经营店ID
func (r *AlibabatclsaelophyrefundfetchgoodsAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabatclsaelophyrefundfetchgoodsAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetOutOrderId is OutOrderId Setter
// 外部订单ID
func (r *AlibabatclsaelophyrefundfetchgoodsAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabatclsaelophyrefundfetchgoodsAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetRefundId is RefundId Setter
// 渠道退款单ID
func (r *AlibabatclsaelophyrefundfetchgoodsAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabatclsaelophyrefundfetchgoodsAPIRequest) GetRefundId() string {
	return r._refundId
}

// SetFetchStartTime is FetchStartTime Setter
// 取货开始时间
func (r *AlibabatclsaelophyrefundfetchgoodsAPIRequest) SetFetchStartTime(_fetchStartTime string) error {
	r._fetchStartTime = _fetchStartTime
	r.Set("fetch_start_time", _fetchStartTime)
	return nil
}

// GetFetchStartTime FetchStartTime Getter
func (r AlibabatclsaelophyrefundfetchgoodsAPIRequest) GetFetchStartTime() string {
	return r._fetchStartTime
}

// SetFetchEndTime is FetchEndTime Setter
// 取货结束时间
func (r *AlibabatclsaelophyrefundfetchgoodsAPIRequest) SetFetchEndTime(_fetchEndTime string) error {
	r._fetchEndTime = _fetchEndTime
	r.Set("fetch_end_time", _fetchEndTime)
	return nil
}

// GetFetchEndTime FetchEndTime Getter
func (r AlibabatclsaelophyrefundfetchgoodsAPIRequest) GetFetchEndTime() string {
	return r._fetchEndTime
}

// SetRemark is Remark Setter
// 备注
func (r *AlibabatclsaelophyrefundfetchgoodsAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r AlibabatclsaelophyrefundfetchgoodsAPIRequest) GetRemark() string {
	return r._remark
}

// SetSubRefundList is SubRefundList Setter
// 外部子订单列表
func (r *AlibabatclsaelophyrefundfetchgoodsAPIRequest) SetSubRefundList(_subRefundList *Subrefundlist) error {
	r._subRefundList = _subRefundList
	r.Set("sub_refund_list", _subRefundList)
	return nil
}

// GetSubRefundList SubRefundList Getter
func (r AlibabatclsaelophyrefundfetchgoodsAPIRequest) GetSubRefundList() *Subrefundlist {
	return r._subRefundList
}

// SetOrderFrom is OrderFrom Setter
// 渠道来源
func (r *AlibabatclsaelophyrefundfetchgoodsAPIRequest) SetOrderFrom(_orderFrom int64) error {
	r._orderFrom = _orderFrom
	r.Set("order_from", _orderFrom)
	return nil
}

// GetOrderFrom OrderFrom Getter
func (r AlibabatclsaelophyrefundfetchgoodsAPIRequest) GetOrderFrom() int64 {
	return r._orderFrom
}
