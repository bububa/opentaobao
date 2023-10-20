package wms

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoBimTradeorderConsignAPIRequest 驱动保税交易订单发货 API请求
// cainiao.bim.tradeorder.consign
//
// 驱动保税交易订单发货
type CainiaoBimTradeorderConsignAPIRequest struct {
	model.Params
	// 交易单号
	_tradeId string
	// 仓储编码，ERP指定仓库发货时需要传值，编码由菜鸟提供
	_storeCode string
	// 选择的线路ID非必填字段
	_resId string
}

// NewCainiaoBimTradeorderConsignRequest 初始化CainiaoBimTradeorderConsignAPIRequest对象
func NewCainiaoBimTradeorderConsignRequest() *CainiaoBimTradeorderConsignAPIRequest {
	return &CainiaoBimTradeorderConsignAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoBimTradeorderConsignAPIRequest) Reset() {
	r._tradeId = ""
	r._storeCode = ""
	r._resId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoBimTradeorderConsignAPIRequest) GetApiMethodName() string {
	return "cainiao.bim.tradeorder.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoBimTradeorderConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoBimTradeorderConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeId is TradeId Setter
// 交易单号
func (r *CainiaoBimTradeorderConsignAPIRequest) SetTradeId(_tradeId string) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r CainiaoBimTradeorderConsignAPIRequest) GetTradeId() string {
	return r._tradeId
}

// SetStoreCode is StoreCode Setter
// 仓储编码，ERP指定仓库发货时需要传值，编码由菜鸟提供
func (r *CainiaoBimTradeorderConsignAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r CainiaoBimTradeorderConsignAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetResId is ResId Setter
// 选择的线路ID非必填字段
func (r *CainiaoBimTradeorderConsignAPIRequest) SetResId(_resId string) error {
	r._resId = _resId
	r.Set("res_id", _resId)
	return nil
}

// GetResId ResId Getter
func (r CainiaoBimTradeorderConsignAPIRequest) GetResId() string {
	return r._resId
}

var poolCainiaoBimTradeorderConsignAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoBimTradeorderConsignRequest()
	},
}

// GetCainiaoBimTradeorderConsignRequest 从 sync.Pool 获取 CainiaoBimTradeorderConsignAPIRequest
func GetCainiaoBimTradeorderConsignAPIRequest() *CainiaoBimTradeorderConsignAPIRequest {
	return poolCainiaoBimTradeorderConsignAPIRequest.Get().(*CainiaoBimTradeorderConsignAPIRequest)
}

// ReleaseCainiaoBimTradeorderConsignAPIRequest 将 CainiaoBimTradeorderConsignAPIRequest 放入 sync.Pool
func ReleaseCainiaoBimTradeorderConsignAPIRequest(v *CainiaoBimTradeorderConsignAPIRequest) {
	v.Reset()
	poolCainiaoBimTradeorderConsignAPIRequest.Put(v)
}
