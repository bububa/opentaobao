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
    _leaveMessagePics   *model.File
    // 留言说明
    _leaveMessage   string
    // 换货单号ID
    _disputeId   int64
    // 拒绝原因ID
    _sellerRefuseReasonId   int64
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
func (r *TmallExchangeReturngoodsRefuseRequest) SetLeaveMessagePics(_leaveMessagePics *model.File) error {
    r._leaveMessagePics = _leaveMessagePics
    r.Set("leave_message_pics", _leaveMessagePics)
    return nil
}

// LeaveMessagePics Getter
func (r TmallExchangeReturngoodsRefuseRequest) GetLeaveMessagePics() *model.File {
    return r._leaveMessagePics
}
// LeaveMessage Setter
// 留言说明
func (r *TmallExchangeReturngoodsRefuseRequest) SetLeaveMessage(_leaveMessage string) error {
    r._leaveMessage = _leaveMessage
    r.Set("leave_message", _leaveMessage)
    return nil
}

// LeaveMessage Getter
func (r TmallExchangeReturngoodsRefuseRequest) GetLeaveMessage() string {
    return r._leaveMessage
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeReturngoodsRefuseRequest) SetDisputeId(_disputeId int64) error {
    r._disputeId = _disputeId
    r.Set("dispute_id", _disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeReturngoodsRefuseRequest) GetDisputeId() int64 {
    return r._disputeId
}
// SellerRefuseReasonId Setter
// 拒绝原因ID
func (r *TmallExchangeReturngoodsRefuseRequest) SetSellerRefuseReasonId(_sellerRefuseReasonId int64) error {
    r._sellerRefuseReasonId = _sellerRefuseReasonId
    r.Set("seller_refuse_reason_id", _sellerRefuseReasonId)
    return nil
}

// SellerRefuseReasonId Getter
func (r TmallExchangeReturngoodsRefuseRequest) GetSellerRefuseReasonId() int64 {
    return r._sellerRefuseReasonId
}
