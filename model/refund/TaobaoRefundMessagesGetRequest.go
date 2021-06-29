package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询退款留言/凭证列表 API请求
taobao.refund.messages.get

查询退款留言/凭证列表
*/
type TaobaoRefundMessagesGetRequest struct {
    model.Params
    // 需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。
    _fields   []string
    // 退款单号
    _refundId   int64
    // 页码
    _pageNo   int64
    // 每页条数
    _pageSize   int64
    // 退款阶段，可选值：onsale（售中），aftersale（售后），天猫退款为必传。
    _refundPhase   string
}

// 初始化TaobaoRefundMessagesGetRequest对象
func NewTaobaoRefundMessagesGetRequest() *TaobaoRefundMessagesGetRequest{
    return &TaobaoRefundMessagesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRefundMessagesGetRequest) GetApiMethodName() string {
    return "taobao.refund.messages.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRefundMessagesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。
func (r *TaobaoRefundMessagesGetRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoRefundMessagesGetRequest) GetFields() []string {
    return r._fields
}
// RefundId Setter
// 退款单号
func (r *TaobaoRefundMessagesGetRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRefundMessagesGetRequest) GetRefundId() int64 {
    return r._refundId
}
// PageNo Setter
// 页码
func (r *TaobaoRefundMessagesGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoRefundMessagesGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页条数
func (r *TaobaoRefundMessagesGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoRefundMessagesGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// RefundPhase Setter
// 退款阶段，可选值：onsale（售中），aftersale（售后），天猫退款为必传。
func (r *TaobaoRefundMessagesGetRequest) SetRefundPhase(_refundPhase string) error {
    r._refundPhase = _refundPhase
    r.Set("refund_phase", _refundPhase)
    return nil
}

// RefundPhase Getter
func (r TaobaoRefundMessagesGetRequest) GetRefundPhase() string {
    return r._refundPhase
}
