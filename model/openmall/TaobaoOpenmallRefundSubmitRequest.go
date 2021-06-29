package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交OpenMall退款单物流 APIRequest
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

func NewTaobaoOpenmallRefundSubmitRequest() *TaobaoOpenmallRefundSubmitRequest{
    return &TaobaoOpenmallRefundSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallRefundSubmitRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.submit"
}

func (r TaobaoOpenmallRefundSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallRefundSubmitRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallRefundSubmitRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallRefundSubmitRequest) SetLogisticsCompanyCode(logisticsCompanyCode string) error {
    r.logisticsCompanyCode = logisticsCompanyCode
    r.Set("logistics_company_code", logisticsCompanyCode)
    return nil
}

func (r TaobaoOpenmallRefundSubmitRequest) GetLogisticsCompanyCode() string {
    return r.logisticsCompanyCode
}

func (r *TaobaoOpenmallRefundSubmitRequest) SetLogisticsCompanyName(logisticsCompanyName string) error {
    r.logisticsCompanyName = logisticsCompanyName
    r.Set("logistics_company_name", logisticsCompanyName)
    return nil
}

func (r TaobaoOpenmallRefundSubmitRequest) GetLogisticsCompanyName() string {
    return r.logisticsCompanyName
}

func (r *TaobaoOpenmallRefundSubmitRequest) SetLogisticsNo(logisticsNo string) error {
    r.logisticsNo = logisticsNo
    r.Set("logistics_no", logisticsNo)
    return nil
}

func (r TaobaoOpenmallRefundSubmitRequest) GetLogisticsNo() string {
    return r.logisticsNo
}

func (r *TaobaoOpenmallRefundSubmitRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoOpenmallRefundSubmitRequest) GetRefundId() int64 {
    return r.refundId
}

