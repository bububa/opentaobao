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
    leaveMessagePics   []*model.File
    // 拒绝换货申请时的留言
    leaveMessage   string
    // 换货单号ID
    disputeId   int64
    // 换货原因对应ID
    sellerRefuseReasonId   int64
    // 返回字段。目前支持dispute_id, bizorder_id, modified, status
    fields   []string
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
func (r *TmallExchangeRefuseRequest) SetLeaveMessagePics(leaveMessagePics []*model.File) error {
    r.leaveMessagePics = leaveMessagePics
    r.Set("leave_message_pics", leaveMessagePics)
    return nil
}

// LeaveMessagePics Getter
func (r TmallExchangeRefuseRequest) GetLeaveMessagePics() []*model.File {
    return r.leaveMessagePics
}
// LeaveMessage Setter
// 拒绝换货申请时的留言
func (r *TmallExchangeRefuseRequest) SetLeaveMessage(leaveMessage string) error {
    r.leaveMessage = leaveMessage
    r.Set("leave_message", leaveMessage)
    return nil
}

// LeaveMessage Getter
func (r TmallExchangeRefuseRequest) GetLeaveMessage() string {
    return r.leaveMessage
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeRefuseRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeRefuseRequest) GetDisputeId() int64 {
    return r.disputeId
}
// SellerRefuseReasonId Setter
// 换货原因对应ID
func (r *TmallExchangeRefuseRequest) SetSellerRefuseReasonId(sellerRefuseReasonId int64) error {
    r.sellerRefuseReasonId = sellerRefuseReasonId
    r.Set("seller_refuse_reason_id", sellerRefuseReasonId)
    return nil
}

// SellerRefuseReasonId Getter
func (r TmallExchangeRefuseRequest) GetSellerRefuseReasonId() int64 {
    return r.sellerRefuseReasonId
}
// Fields Setter
// 返回字段。目前支持dispute_id, bizorder_id, modified, status
func (r *TmallExchangeRefuseRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TmallExchangeRefuseRequest) GetFields() []string {
    return r.fields
}
