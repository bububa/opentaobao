package wms

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoBimTradeorderConsignAPIRequest) GetApiMethodName() string {
	return "cainiao.bim.tradeorder.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoBimTradeorderConsignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
