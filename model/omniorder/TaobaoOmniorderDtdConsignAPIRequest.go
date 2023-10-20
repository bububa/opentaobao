package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderdtdconsignAPIRequest 门店自送发货 API请求
// taobao.omniorder.dtd.consign
//
// 该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单
type TaobaoomniorderdtdconsignAPIRequest struct {
	model.Params
	// 淘宝订单主订单号
	_mainOrderId int64
	// 发货对应的商户中心门店ID
	_storeId int64
}

// NewTaobaoomniorderdtdconsignRequest 初始化TaobaoomniorderdtdconsignAPIRequest对象
func NewTaobaoomniorderdtdconsignRequest() *TaobaoomniorderdtdconsignAPIRequest {
	return &TaobaoomniorderdtdconsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderdtdconsignAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.dtd.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderdtdconsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderdtdconsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 淘宝订单主订单号
func (r *TaobaoomniorderdtdconsignAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoomniorderdtdconsignAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetStoreId is StoreId Setter
// 发货对应的商户中心门店ID
func (r *TaobaoomniorderdtdconsignAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoomniorderdtdconsignAPIRequest) GetStoreId() int64 {
	return r._storeId
}
