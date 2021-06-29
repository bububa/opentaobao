package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
驱动保税交易订单发货 API请求
cainiao.bim.tradeorder.consign

驱动保税交易订单发货
*/
type CainiaoBimTradeorderConsignRequest struct {
    model.Params
    // 交易单号
    _tradeId   string
    // 仓储编码，ERP指定仓库发货时需要传值，编码由菜鸟提供
    _storeCode   string
    // 选择的线路ID非必填字段
    _resId   string
}

// 初始化CainiaoBimTradeorderConsignRequest对象
func NewCainiaoBimTradeorderConsignRequest() *CainiaoBimTradeorderConsignRequest{
    return &CainiaoBimTradeorderConsignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoBimTradeorderConsignRequest) GetApiMethodName() string {
    return "cainiao.bim.tradeorder.consign"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoBimTradeorderConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeId Setter
// 交易单号
func (r *CainiaoBimTradeorderConsignRequest) SetTradeId(_tradeId string) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r CainiaoBimTradeorderConsignRequest) GetTradeId() string {
    return r._tradeId
}
// StoreCode Setter
// 仓储编码，ERP指定仓库发货时需要传值，编码由菜鸟提供
func (r *CainiaoBimTradeorderConsignRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r CainiaoBimTradeorderConsignRequest) GetStoreCode() string {
    return r._storeCode
}
// ResId Setter
// 选择的线路ID非必填字段
func (r *CainiaoBimTradeorderConsignRequest) SetResId(_resId string) error {
    r._resId = _resId
    r.Set("res_id", _resId)
    return nil
}

// ResId Getter
func (r CainiaoBimTradeorderConsignRequest) GetResId() string {
    return r._resId
}
