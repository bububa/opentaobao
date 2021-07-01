package alscmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑凭证售后退 API请求
alibaba.alsc.merchant.ext.ticket.refund

口碑凭证售后退
*/
type AlibabaAlscMerchantExtTicketRefundAPIRequest struct {
    model.Params
    // 券核销流水号，针对该次核销发起售后退操作
    _transId   string
    // 外部请求号，支持英文字母和数字，由开发者自行定义（不允许重复）
    _ticketRequestId   string
    // 凭证码，包括内部凭证码和外部凭证码，内部凭证码为12位，纯数字，且唯一不重复
    _ticketCode   string
}

// 初始化AlibabaAlscMerchantExtTicketRefundAPIRequest对象
func NewAlibabaAlscMerchantExtTicketRefundRequest() *AlibabaAlscMerchantExtTicketRefundAPIRequest{
    return &AlibabaAlscMerchantExtTicketRefundAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscMerchantExtTicketRefundAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.merchant.ext.ticket.refund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscMerchantExtTicketRefundAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TransId Setter
// 券核销流水号，针对该次核销发起售后退操作
func (r *AlibabaAlscMerchantExtTicketRefundAPIRequest) SetTransId(_transId string) error {
    r._transId = _transId
    r.Set("trans_id", _transId)
    return nil
}

// TransId Getter
func (r AlibabaAlscMerchantExtTicketRefundAPIRequest) GetTransId() string {
    return r._transId
}
// TicketRequestId Setter
// 外部请求号，支持英文字母和数字，由开发者自行定义（不允许重复）
func (r *AlibabaAlscMerchantExtTicketRefundAPIRequest) SetTicketRequestId(_ticketRequestId string) error {
    r._ticketRequestId = _ticketRequestId
    r.Set("ticket_request_id", _ticketRequestId)
    return nil
}

// TicketRequestId Getter
func (r AlibabaAlscMerchantExtTicketRefundAPIRequest) GetTicketRequestId() string {
    return r._ticketRequestId
}
// TicketCode Setter
// 凭证码，包括内部凭证码和外部凭证码，内部凭证码为12位，纯数字，且唯一不重复
func (r *AlibabaAlscMerchantExtTicketRefundAPIRequest) SetTicketCode(_ticketCode string) error {
    r._ticketCode = _ticketCode
    r.Set("ticket_code", _ticketCode)
    return nil
}

// TicketCode Getter
func (r AlibabaAlscMerchantExtTicketRefundAPIRequest) GetTicketCode() string {
    return r._ticketCode
}
