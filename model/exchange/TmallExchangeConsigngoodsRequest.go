package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家发货 API请求
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

// 初始化TmallExchangeConsigngoodsRequest对象
func NewTmallExchangeConsigngoodsRequest() *TmallExchangeConsigngoodsRequest{
    return &TmallExchangeConsigngoodsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallExchangeConsigngoodsRequest) GetApiMethodName() string {
    return "tmall.exchange.consigngoods"
}

// IRequest interface 方法, 获取API参数
func (r TmallExchangeConsigngoodsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeConsigngoodsRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeConsigngoodsRequest) GetDisputeId() int64 {
    return r.disputeId
}
// LogisticsNo Setter
// 卖家发货的物流单号
func (r *TmallExchangeConsigngoodsRequest) SetLogisticsNo(logisticsNo string) error {
    r.logisticsNo = logisticsNo
    r.Set("logistics_no", logisticsNo)
    return nil
}

// LogisticsNo Getter
func (r TmallExchangeConsigngoodsRequest) GetLogisticsNo() string {
    return r.logisticsNo
}
// LogisticsType Setter
// 卖家发货的物流类型，100表示平邮，200表示快递
func (r *TmallExchangeConsigngoodsRequest) SetLogisticsType(logisticsType int64) error {
    r.logisticsType = logisticsType
    r.Set("logistics_type", logisticsType)
    return nil
}

// LogisticsType Getter
func (r TmallExchangeConsigngoodsRequest) GetLogisticsType() int64 {
    return r.logisticsType
}
// LogisticsCompanyName Setter
// 卖家发货的快递公司
func (r *TmallExchangeConsigngoodsRequest) SetLogisticsCompanyName(logisticsCompanyName string) error {
    r.logisticsCompanyName = logisticsCompanyName
    r.Set("logistics_company_name", logisticsCompanyName)
    return nil
}

// LogisticsCompanyName Getter
func (r TmallExchangeConsigngoodsRequest) GetLogisticsCompanyName() string {
    return r.logisticsCompanyName
}
// Fields Setter
// 返回字段
func (r *TmallExchangeConsigngoodsRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TmallExchangeConsigngoodsRequest) GetFields() []string {
    return r.fields
}
