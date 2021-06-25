package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
3PL直邮线下发货 APIRequest
taobao.wlb.import.threepl.offline.consign

菜鸟认证直邮线下发货
*/
type TaobaoWlbImportThreeplOfflineConsignRequest struct {
    model.Params

    // 交易单号
    tradeId   int64 

    // 资源id
    resId   int64 

    // 资源code
    resCode   string 

    // 运单号
    waybillNo   string 

    // 发件人地址库id
    fromId   int64 

}

func NewTaobaoWlbImportThreeplOfflineConsignRequest() *TaobaoWlbImportThreeplOfflineConsignRequest{
    return &TaobaoWlbImportThreeplOfflineConsignRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbImportThreeplOfflineConsignRequest) GetApiMethodName() string {
    return "taobao.wlb.import.threepl.offline.consign"
}

func (r TaobaoWlbImportThreeplOfflineConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbImportThreeplOfflineConsignRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r TaobaoWlbImportThreeplOfflineConsignRequest) GetTradeId() int64 {
    return r.tradeId
}

func (r *TaobaoWlbImportThreeplOfflineConsignRequest) SetResId(resId int64) error {
    r.resId = resId
    r.Set("res_id", resId)
    return nil
}

func (r TaobaoWlbImportThreeplOfflineConsignRequest) GetResId() int64 {
    return r.resId
}

func (r *TaobaoWlbImportThreeplOfflineConsignRequest) SetResCode(resCode string) error {
    r.resCode = resCode
    r.Set("res_code", resCode)
    return nil
}

func (r TaobaoWlbImportThreeplOfflineConsignRequest) GetResCode() string {
    return r.resCode
}

func (r *TaobaoWlbImportThreeplOfflineConsignRequest) SetWaybillNo(waybillNo string) error {
    r.waybillNo = waybillNo
    r.Set("waybill_no", waybillNo)
    return nil
}

func (r TaobaoWlbImportThreeplOfflineConsignRequest) GetWaybillNo() string {
    return r.waybillNo
}

func (r *TaobaoWlbImportThreeplOfflineConsignRequest) SetFromId(fromId int64) error {
    r.fromId = fromId
    r.Set("from_id", fromId)
    return nil
}

func (r TaobaoWlbImportThreeplOfflineConsignRequest) GetFromId() int64 {
    return r.fromId
}

