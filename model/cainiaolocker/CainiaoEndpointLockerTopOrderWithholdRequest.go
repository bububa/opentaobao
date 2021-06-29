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
    _companyCode   string
    // 柜子id
    _guiId   string
    // 订单类型(0-取件业务，1-寄件业务，2-派样业务)
    _orderType   int64
    // 开放用户id
    _openUserId   string
    // 代扣金额（全额），单位：分
    _totalFee   int64
    // 扩展字段
    _extra   string
    // 柜子订单编码
    _orderCode   string
    // 运单号
    _mailNo   string
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
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetCompanyCode(_companyCode string) error {
    r._companyCode = _companyCode
    r.Set("company_code", _companyCode)
    return nil
}

// CompanyCode Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetCompanyCode() string {
    return r._companyCode
}
// GuiId Setter
// 柜子id
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetGuiId(_guiId string) error {
    r._guiId = _guiId
    r.Set("gui_id", _guiId)
    return nil
}

// GuiId Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetGuiId() string {
    return r._guiId
}
// OrderType Setter
// 订单类型(0-取件业务，1-寄件业务，2-派样业务)
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetOrderType(_orderType int64) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetOrderType() int64 {
    return r._orderType
}
// OpenUserId Setter
// 开放用户id
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetOpenUserId(_openUserId string) error {
    r._openUserId = _openUserId
    r.Set("open_user_id", _openUserId)
    return nil
}

// OpenUserId Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetOpenUserId() string {
    return r._openUserId
}
// TotalFee Setter
// 代扣金额（全额），单位：分
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetTotalFee(_totalFee int64) error {
    r._totalFee = _totalFee
    r.Set("total_fee", _totalFee)
    return nil
}

// TotalFee Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetTotalFee() int64 {
    return r._totalFee
}
// Extra Setter
// 扩展字段
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetExtra(_extra string) error {
    r._extra = _extra
    r.Set("extra", _extra)
    return nil
}

// Extra Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetExtra() string {
    return r._extra
}
// OrderCode Setter
// 柜子订单编码
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetOrderCode() string {
    return r._orderCode
}
// MailNo Setter
// 运单号
func (r *CainiaoEndpointLockerTopOrderWithholdRequest) SetMailNo(_mailNo string) error {
    r._mailNo = _mailNo
    r.Set("mail_no", _mailNo)
    return nil
}

// MailNo Getter
func (r CainiaoEndpointLockerTopOrderWithholdRequest) GetMailNo() string {
    return r._mailNo
}
