package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeStoreDeliversGetAPIRequest 获取门店采购单下的发货单列表 API请求
// alibaba.nlife.store.delivers.get
//
// 获取门店采购单下的发货单列表
type AlibabaNlifeStoreDeliversGetAPIRequest struct {
	model.Params
	// 门店采购订单号
	_tradeNo string
	// 零售商的门店id
	_storeId int64
	// 每页的数量
	_pageSize int64
	// 查询的页码
	_pageNo int64
}

// NewAlibabaNlifeStoreDeliversGetRequest 初始化AlibabaNlifeStoreDeliversGetAPIRequest对象
func NewAlibabaNlifeStoreDeliversGetRequest() *AlibabaNlifeStoreDeliversGetAPIRequest {
	return &AlibabaNlifeStoreDeliversGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreDeliversGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.store.delivers.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreDeliversGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TradeNo Setter
// 门店采购订单号
func (r *AlibabaNlifeStoreDeliversGetAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// Get TradeNo Getter
func (r AlibabaNlifeStoreDeliversGetAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// Set is StoreId Setter
// 零售商的门店id
func (r *AlibabaNlifeStoreDeliversGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaNlifeStoreDeliversGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is PageSize Setter
// 每页的数量
func (r *AlibabaNlifeStoreDeliversGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaNlifeStoreDeliversGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is PageNo Setter
// 查询的页码
func (r *AlibabaNlifeStoreDeliversGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r AlibabaNlifeStoreDeliversGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
