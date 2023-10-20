package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradequeuequeryAPIRequest 尖货交易排队信息查询 API请求
// taobao.opentrade.queue.query
//
// 尖货交易排队信息查询
type TaobaoopentradequeuequeryAPIRequest struct {
	model.Params
	// 排队用户状态，新用户为NEW
	_status string
	// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
	_activityId string
	// 排队商品SKU ID，不存在传0
	_skuId int64
	// 排队商品ID
	_itemId int64
	// 分页参数，每页大小
	_pageSize int64
	// 分页参数，当前页，以0开始
	_pageIndex int64
}

// NewTaobaoopentradequeuequeryRequest 初始化TaobaoopentradequeuequeryAPIRequest对象
func NewTaobaoopentradequeuequeryRequest() *TaobaoopentradequeuequeryAPIRequest {
	return &TaobaoopentradequeuequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradequeuequeryAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.queue.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradequeuequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradequeuequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 排队用户状态，新用户为NEW
func (r *TaobaoopentradequeuequeryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoopentradequeuequeryAPIRequest) GetStatus() string {
	return r._status
}

// SetActivityId is ActivityId Setter
// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
func (r *TaobaoopentradequeuequeryAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoopentradequeuequeryAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetSkuId is SkuId Setter
// 排队商品SKU ID，不存在传0
func (r *TaobaoopentradequeuequeryAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoopentradequeuequeryAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetItemId is ItemId Setter
// 排队商品ID
func (r *TaobaoopentradequeuequeryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoopentradequeuequeryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetPageSize is PageSize Setter
// 分页参数，每页大小
func (r *TaobaoopentradequeuequeryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoopentradequeuequeryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 分页参数，当前页，以0开始
func (r *TaobaoopentradequeuequeryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoopentradequeuequeryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
