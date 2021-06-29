package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代扣支付 APIRequest
cainiao.endpoint.locker.top.order.withhold

提供代扣，允许有一笔欠款。
*/
type CainiaoEndpointLockerTopOrderWithholdRequest struct {
    model.Params

    // 柜子公司编码
    companyCode   string 

    // 柜子id
    guiId   string 

    // 订单类型(0-取件业务，1-寄件业务，2-派样业务)
    orderType   int64 

    // 开放用户id
    openUserId   string 

    // 代扣金额（全额），单位：分
    totalFee   int64 

    // 扩展字段
    extra   string 

    // 柜子订单编码
    orderCode   string 

    // 运单号
    mailNo   string 

}

func NewCainiaoEndpointLockerTopOrderWithholdRequest() *CainiaoEndpointLockerTopOrderWithholdRequest{
    return &CainiaoEndpointLockerTopOrderWithholdRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.order.withhold"
}

func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetCompanyCode(companyCode string) error {
    r.companyCode = companyCode
    r.Set("company_code", companyCode)
    return nil
}

func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetCompanyCode() string {
    return r.companyCode
}

func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetGuiId(guiId string) error {
    r.guiId = guiId
    r.Set("gui_id", guiId)
    return nil
}

func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetGuiId() string {
    return r.guiId
}

func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetOrderType() int64 {
    return r.orderType
}

func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetOpenUserId() string {
    return r.openUserId
}

func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetTotalFee(totalFee int64) error {
    r.totalFee = totalFee
    r.Set("total_fee", totalFee)
    return nil
}

func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetTotalFee() int64 {
    return r.totalFee
}

func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetExtra() string {
    return r.extra
}

func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetMailNo() string {
    return r.mailNo
}

