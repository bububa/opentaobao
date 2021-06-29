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
type AlibabaAlscMerchantExtTicketRefundRequest struct {
    model.Params
    // 券核销流水号，针对该次核销发起售后退操作
    transId   string
    // 外部请求号，支持英文字母和数字，由开发者自行定义（不允许重复）
    ticketRequestId   string
    // 凭证码，包括内部凭证码和外部凭证码，内部凭证码为12位，纯数字，且唯一不重复
    ticketCode   string
}

// 初始化AlibabaAlscMerchantExtTicketRefundRequest对象
func NewAlibabaAlscMerchantExtTicketRefundRequest() *AlibabaAlscMerchantExtTicketRefundRequest{
    return &AlibabaAlscMerchantExtTicketRefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscMerchantExtTicketRefundRequest) GetApiMethodName() string {
    return "alibaba.alsc.merchant.ext.ticket.refund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscMerchantExtTicketRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TransId Setter
// 券核销流水号，针对该次核销发起售后退操作
func (r *AlibabaAlscMerchantExtTicketRefundRequest) SetTransId(transId string) error {
    r.transId = transId
    r.Set("trans_id", transId)
    return nil
}

// TransId Getter
func (r AlibabaAlscMerchantExtTicketRefundRequest) GetTransId() string {
    return r.transId
}
// TicketRequestId Setter
// 外部请求号，支持英文字母和数字，由开发者自行定义（不允许重复）
func (r *AlibabaAlscMerchantExtTicketRefundRequest) SetTicketRequestId(ticketRequestId string) error {
    r.ticketRequestId = ticketRequestId
    r.Set("ticket_request_id", ticketRequestId)
    return nil
}

// TicketRequestId Getter
func (r AlibabaAlscMerchantExtTicketRefundRequest) GetTicketRequestId() string {
    return r.ticketRequestId
}
// TicketCode Setter
// 凭证码，包括内部凭证码和外部凭证码，内部凭证码为12位，纯数字，且唯一不重复
func (r *AlibabaAlscMerchantExtTicketRefundRequest) SetTicketCode(ticketCode string) error {
    r.ticketCode = ticketCode
    r.Set("ticket_code", ticketCode)
    return nil
}

// TicketCode Getter
func (r AlibabaAlscMerchantExtTicketRefundRequest) GetTicketCode() string {
    return r.ticketCode
}
