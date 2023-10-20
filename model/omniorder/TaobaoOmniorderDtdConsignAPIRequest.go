package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderDtdConsignAPIRequest 门店自送发货 API请求
// taobao.omniorder.dtd.consign
//
// 该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单
type TaobaoOmniorderDtdConsignAPIRequest struct {
	model.Params
	// 淘宝订单主订单号
	_mainOrderId int64
	// 发货对应的商户中心门店ID
	_storeId int64
}

// NewTaobaoOmniorderDtdConsignRequest 初始化TaobaoOmniorderDtdConsignAPIRequest对象
func NewTaobaoOmniorderDtdConsignRequest() *TaobaoOmniorderDtdConsignAPIRequest {
	return &TaobaoOmniorderDtdConsignAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderDtdConsignAPIRequest) Reset() {
	r._mainOrderId = 0
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderDtdConsignAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.dtd.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderDtdConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderDtdConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝订单主订单号
func (r *TaobaoOmniorderDtdConsignAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoOmniorderDtdConsignAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetStoreId is StoreId Setter
// 发货对应的商户中心门店ID
func (r *TaobaoOmniorderDtdConsignAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoOmniorderDtdConsignAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolTaobaoOmniorderDtdConsignAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderDtdConsignRequest()
	},
}

// GetTaobaoOmniorderDtdConsignRequest 从 sync.Pool 获取 TaobaoOmniorderDtdConsignAPIRequest
func GetTaobaoOmniorderDtdConsignAPIRequest() *TaobaoOmniorderDtdConsignAPIRequest {
	return poolTaobaoOmniorderDtdConsignAPIRequest.Get().(*TaobaoOmniorderDtdConsignAPIRequest)
}

// ReleaseTaobaoOmniorderDtdConsignAPIRequest 将 TaobaoOmniorderDtdConsignAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderDtdConsignAPIRequest(v *TaobaoOmniorderDtdConsignAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderDtdConsignAPIRequest.Put(v)
}
