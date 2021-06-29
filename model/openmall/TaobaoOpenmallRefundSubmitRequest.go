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
    _distributor   string
    // 物流公司编码
    _logisticsCompanyCode   string
    // 物流公司名称
    _logisticsCompanyName   string
    // 快递单号
    _logisticsNo   string
    // 退款单ID
    _refundId   int64
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
func (r *TaobaoOpenmallRefundSubmitRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundSubmitRequest) GetDistributor() string {
    return r._distributor
}
// LogisticsCompanyCode Setter
// 物流公司编码
func (r *TaobaoOpenmallRefundSubmitRequest) SetLogisticsCompanyCode(_logisticsCompanyCode string) error {
    r._logisticsCompanyCode = _logisticsCompanyCode
    r.Set("logistics_company_code", _logisticsCompanyCode)
    return nil
}

// LogisticsCompanyCode Getter
func (r TaobaoOpenmallRefundSubmitRequest) GetLogisticsCompanyCode() string {
    return r._logisticsCompanyCode
}
// LogisticsCompanyName Setter
// 物流公司名称
func (r *TaobaoOpenmallRefundSubmitRequest) SetLogisticsCompanyName(_logisticsCompanyName string) error {
    r._logisticsCompanyName = _logisticsCompanyName
    r.Set("logistics_company_name", _logisticsCompanyName)
    return nil
}

// LogisticsCompanyName Getter
func (r TaobaoOpenmallRefundSubmitRequest) GetLogisticsCompanyName() string {
    return r._logisticsCompanyName
}
// LogisticsNo Setter
// 快递单号
func (r *TaobaoOpenmallRefundSubmitRequest) SetLogisticsNo(_logisticsNo string) error {
    r._logisticsNo = _logisticsNo
    r.Set("logistics_no", _logisticsNo)
    return nil
}

// LogisticsNo Getter
func (r TaobaoOpenmallRefundSubmitRequest) GetLogisticsNo() string {
    return r._logisticsNo
}
// RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundSubmitRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoOpenmallRefundSubmitRequest) GetRefundId() int64 {
    return r._refundId
}
