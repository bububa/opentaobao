package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家发货 APIRequest
tmall.exchange.consigngoods

卖家发货
*/
type TmallExchangeConsigngoodsRequest struct {
    model.Params

    // 换货单号ID
    disputeId   int64 

    // 卖家发货的物流单号
    logisticsNo   string 

    // 卖家发货的物流类型，100表示平邮，200表示快递
    logisticsType   int64 

    // 卖家发货的快递公司
    logisticsCompanyName   string 

    // 返回字段
    fields   []string 

}

func NewTmallExchangeConsigngoodsRequest() *TmallExchangeConsigngoodsRequest{
    return &TmallExchangeConsigngoodsRequest{
        Params: model.NewParams(),
    }
}

func (r TmallExchangeConsigngoodsRequest) GetApiMethodName() string {
    return "tmall.exchange.consigngoods"
}

func (r TmallExchangeConsigngoodsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallExchangeConsigngoodsRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

func (r TmallExchangeConsigngoodsRequest) GetDisputeId() int64 {
    return r.disputeId
}

func (r *TmallExchangeConsigngoodsRequest) SetLogisticsNo(logisticsNo string) error {
    r.logisticsNo = logisticsNo
    r.Set("logistics_no", logisticsNo)
    return nil
}

func (r TmallExchangeConsigngoodsRequest) GetLogisticsNo() string {
    return r.logisticsNo
}

func (r *TmallExchangeConsigngoodsRequest) SetLogisticsType(logisticsType int64) error {
    r.logisticsType = logisticsType
    r.Set("logistics_type", logisticsType)
    return nil
}

func (r TmallExchangeConsigngoodsRequest) GetLogisticsType() int64 {
    return r.logisticsType
}

func (r *TmallExchangeConsigngoodsRequest) SetLogisticsCompanyName(logisticsCompanyName string) error {
    r.logisticsCompanyName = logisticsCompanyName
    r.Set("logistics_company_name", logisticsCompanyName)
    return nil
}

func (r TmallExchangeConsigngoodsRequest) GetLogisticsCompanyName() string {
    return r.logisticsCompanyName
}

func (r *TmallExchangeConsigngoodsRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TmallExchangeConsigngoodsRequest) GetFields() []string {
    return r.fields
}

