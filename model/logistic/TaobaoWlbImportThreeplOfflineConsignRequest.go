package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
3PL直邮线下发货 API请求
taobao.wlb.import.threepl.offline.consign

菜鸟认证直邮线下发货
*/
type TaobaoWlbImportThreeplOfflineConsignAPIRequest struct {
    model.Params
    // 交易单号
    _tradeId   int64
    // 资源id
    _resId   int64
    // 资源code
    _resCode   string
    // 运单号
    _waybillNo   string
    // 发件人地址库id
    _fromId   int64
}

// 初始化TaobaoWlbImportThreeplOfflineConsignAPIRequest对象
func NewTaobaoWlbImportThreeplOfflineConsignRequest() *TaobaoWlbImportThreeplOfflineConsignAPIRequest{
    return &TaobaoWlbImportThreeplOfflineConsignAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.import.threepl.offline.consign"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeId Setter
// 交易单号
func (r *TaobaoWlbImportThreeplOfflineConsignAPIRequest) SetTradeId(_tradeId int64) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetTradeId() int64 {
    return r._tradeId
}
// ResId Setter
// 资源id
func (r *TaobaoWlbImportThreeplOfflineConsignAPIRequest) SetResId(_resId int64) error {
    r._resId = _resId
    r.Set("res_id", _resId)
    return nil
}

// ResId Getter
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetResId() int64 {
    return r._resId
}
// ResCode Setter
// 资源code
func (r *TaobaoWlbImportThreeplOfflineConsignAPIRequest) SetResCode(_resCode string) error {
    r._resCode = _resCode
    r.Set("res_code", _resCode)
    return nil
}

// ResCode Getter
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetResCode() string {
    return r._resCode
}
// WaybillNo Setter
// 运单号
func (r *TaobaoWlbImportThreeplOfflineConsignAPIRequest) SetWaybillNo(_waybillNo string) error {
    r._waybillNo = _waybillNo
    r.Set("waybill_no", _waybillNo)
    return nil
}

// WaybillNo Getter
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetWaybillNo() string {
    return r._waybillNo
}
// FromId Setter
// 发件人地址库id
func (r *TaobaoWlbImportThreeplOfflineConsignAPIRequest) SetFromId(_fromId int64) error {
    r._fromId = _fromId
    r.Set("from_id", _fromId)
    return nil
}

// FromId Getter
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetFromId() int64 {
    return r._fromId
}
