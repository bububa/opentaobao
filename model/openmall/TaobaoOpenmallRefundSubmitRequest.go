package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交OpenMall退款单物流 API请求
taobao.openmall.refund.submit

提交OpenMall退款单物流
*/
type TaobaoOpenmallRefundSubmitRequest struct {
    model.Params
    // 渠道
    distributor   string
    // 物流公司编码
    logisticsCompanyCode   string
    // 物流公司名称
    logisticsCompanyName   string
    // 快递单号
    logisticsNo   string
    // 退款单ID
    refundId   int64
}

// 初始化TaobaoOpenmallRefundSubmitRequest对象
func NewTaobaoOpenmallRefundSubmitRequest() *TaobaoOpenmallRefundSubmitRequest{
    return &TaobaoOpenmallRefundSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundSubmitRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.submit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 渠道
func (r *TaobaoOpenmallRefundSubmitRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundSubmitRequest) GetDistributor() string {
    return r.distributor
}
// LogisticsCompanyCode Setter
// 物流公司编码
func (r *TaobaoOpenmallRefundSubmitRequest) SetLogisticsCompanyCode(logisticsCompanyCode string) error {
    r.logisticsCompanyCode = logisticsCompanyCode
    r.Set("logistics_company_code", logisticsCompanyCode)
    return nil
}

// LogisticsCompanyCode Getter
func (r TaobaoOpenmallRefundSubmitRequest) GetLogisticsCompanyCode() string {
    return r.logisticsCompanyCode
}
// LogisticsCompanyName Setter
// 物流公司名称
func (r *TaobaoOpenmallRefundSubmitRequest) SetLogisticsCompanyName(logisticsCompanyName string) error {
    r.logisticsCompanyName = logisticsCompanyName
    r.Set("logistics_company_name", logisticsCompanyName)
    return nil
}

// LogisticsCompanyName Getter
func (r TaobaoOpenmallRefundSubmitRequest) GetLogisticsCompanyName() string {
    return r.logisticsCompanyName
}
// LogisticsNo Setter
// 快递单号
func (r *TaobaoOpenmallRefundSubmitRequest) SetLogisticsNo(logisticsNo string) error {
    r.logisticsNo = logisticsNo
    r.Set("logistics_no", logisticsNo)
    return nil
}

// LogisticsNo Getter
func (r TaobaoOpenmallRefundSubmitRequest) GetLogisticsNo() string {
    return r.logisticsNo
}
// RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundSubmitRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoOpenmallRefundSubmitRequest) GetRefundId() int64 {
    return r.refundId
}
