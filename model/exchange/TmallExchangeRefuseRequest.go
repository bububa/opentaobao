package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝换货申请 API请求
tmall.exchange.refuse

卖家拒绝换货申请
*/
type TmallExchangeRefuseRequest struct {
    model.Params
    // 凭证图片
    _leaveMessagePics   []*model.File
    // 拒绝换货申请时的留言
    _leaveMessage   string
    // 换货单号ID
    _disputeId   int64
    // 换货原因对应ID
    _sellerRefuseReasonId   int64
    // 返回字段。目前支持dispute_id, bizorder_id, modified, status
    _fields   []string
}

// 初始化TmallExchangeRefuseRequest对象
func NewTmallExchangeRefuseRequest() *TmallExchangeRefuseRequest{
    return &TmallExchangeRefuseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallExchangeRefuseRequest) GetApiMethodName() string {
    return "tmall.exchange.refuse"
}

// IRequest interface 方法, 获取API参数
func (r TmallExchangeRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LeaveMessagePics Setter
// 凭证图片
func (r *TmallExchangeRefuseRequest) SetLeaveMessagePics(_leaveMessagePics []*model.File) error {
    r._leaveMessagePics = _leaveMessagePics
    r.Set("leave_message_pics", _leaveMessagePics)
    return nil
}

// LeaveMessagePics Getter
func (r TmallExchangeRefuseRequest) GetLeaveMessagePics() []*model.File {
    return r._leaveMessagePics
}
// LeaveMessage Setter
// 拒绝换货申请时的留言
func (r *TmallExchangeRefuseRequest) SetLeaveMessage(_leaveMessage string) error {
    r._leaveMessage = _leaveMessage
    r.Set("leave_message", _leaveMessage)
    return nil
}

// LeaveMessage Getter
func (r TmallExchangeRefuseRequest) GetLeaveMessage() string {
    return r._leaveMessage
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeRefuseRequest) SetDisputeId(_disputeId int64) error {
    r._disputeId = _disputeId
    r.Set("dispute_id", _disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeRefuseRequest) GetDisputeId() int64 {
    return r._disputeId
}
// SellerRefuseReasonId Setter
// 换货原因对应ID
func (r *TmallExchangeRefuseRequest) SetSellerRefuseReasonId(_sellerRefuseReasonId int64) error {
    r._sellerRefuseReasonId = _sellerRefuseReasonId
    r.Set("seller_refuse_reason_id", _sellerRefuseReasonId)
    return nil
}

// SellerRefuseReasonId Getter
func (r TmallExchangeRefuseRequest) GetSellerRefuseReasonId() int64 {
    return r._sellerRefuseReasonId
}
// Fields Setter
// 返回字段。目前支持dispute_id, bizorder_id, modified, status
func (r *TmallExchangeRefuseRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TmallExchangeRefuseRequest) GetFields() []string {
    return r._fields
}