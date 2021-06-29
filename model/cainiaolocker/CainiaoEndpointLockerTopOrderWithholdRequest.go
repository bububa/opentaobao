package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代扣支付 API请求
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

// 初始化CainiaoEndpointLockerTopOrderWithholdRequest对象
func NewCainiaoEndpointLockerTopOrderWithholdRequest() *CainiaoEndpointLockerTopOrderWithholdRequest{
    return &CainiaoEndpointLockerTopOrderWithholdRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.order.withhold"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyCode Setter
// 柜子公司编码
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetCompanyCode(companyCode string) error {
    r.companyCode = companyCode
    r.Set("company_code", companyCode)
    return nil
}

// CompanyCode Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetCompanyCode() string {
    return r.companyCode
}
// GuiId Setter
// 柜子id
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetGuiId(guiId string) error {
    r.guiId = guiId
    r.Set("gui_id", guiId)
    return nil
}

// GuiId Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetGuiId() string {
    return r.guiId
}
// OrderType Setter
// 订单类型(0-取件业务，1-寄件业务，2-派样业务)
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

// OrderType Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetOrderType() int64 {
    return r.orderType
}
// OpenUserId Setter
// 开放用户id
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

// OpenUserId Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetOpenUserId() string {
    return r.openUserId
}
// TotalFee Setter
// 代扣金额（全额），单位：分
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetTotalFee(totalFee int64) error {
    r.totalFee = totalFee
    r.Set("total_fee", totalFee)
    return nil
}

// TotalFee Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetTotalFee() int64 {
    return r.totalFee
}
// Extra Setter
// 扩展字段
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetExtra() string {
    return r.extra
}
// OrderCode Setter
// 柜子订单编码
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetOrderCode() string {
    return r.orderCode
}
// MailNo Setter
// 运单号
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

// MailNo Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetMailNo() string {
    return r.mailNo
}
