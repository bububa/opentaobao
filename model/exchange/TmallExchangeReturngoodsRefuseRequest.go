package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝确认收货 API请求
tmall.exchange.returngoods.refuse

卖家拒绝买家换货申请
*/
type TmallExchangeReturngoodsRefuseRequest struct {
    model.Params
    // 凭证图片
    leaveMessagePics   []*model.File
    // 留言说明
    leaveMessage   string
    // 换货单号ID
    disputeId   int64
    // 拒绝原因ID
    sellerRefuseReasonId   int64
}

// 初始化TmallExchangeReturngoodsRefuseRequest对象
func NewTmallExchangeReturngoodsRefuseRequest() *TmallExchangeReturngoodsRefuseRequest{
    return &TmallExchangeReturngoodsRefuseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallExchangeReturngoodsRefuseRequest) GetApiMethodName() string {
    return "tmall.exchange.returngoods.refuse"
}

// IRequest interface 方法, 获取API参数
func (r TmallExchangeReturngoodsRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LeaveMessagePics Setter
// 凭证图片
func (r *TmallExchangeReturngoodsRefuseRequest) SetLeaveMessagePics(leaveMessagePics []*model.File) error {
    r.leaveMessagePics = leaveMessagePics
    r.Set("leave_message_pics", leaveMessagePics)
    return nil
}

// LeaveMessagePics Getter
func (r TmallExchangeReturngoodsRefuseRequest) GetLeaveMessagePics() []*model.File {
    return r.leaveMessagePics
}
// LeaveMessage Setter
// 留言说明
func (r *TmallExchangeReturngoodsRefuseRequest) SetLeaveMessage(leaveMessage string) error {
    r.leaveMessage = leaveMessage
    r.Set("leave_message", leaveMessage)
    return nil
}

// LeaveMessage Getter
func (r TmallExchangeReturngoodsRefuseRequest) GetLeaveMessage() string {
    return r.leaveMessage
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeReturngoodsRefuseRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeReturngoodsRefuseRequest) GetDisputeId() int64 {
    return r.disputeId
}
// SellerRefuseReasonId Setter
// 拒绝原因ID
func (r *TmallExchangeReturngoodsRefuseRequest) SetSellerRefuseReasonId(sellerRefuseReasonId int64) error {
    r.sellerRefuseReasonId = sellerRefuseReasonId
    r.Set("seller_refuse_reason_id", sellerRefuseReasonId)
    return nil
}

// SellerRefuseReasonId Getter
func (r TmallExchangeReturngoodsRefuseRequest) GetSellerRefuseReasonId() int64 {
    return r.sellerRefuseReasonId
}
