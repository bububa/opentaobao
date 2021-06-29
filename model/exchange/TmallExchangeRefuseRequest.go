package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝换货申请 APIRequest
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

func NewTmallExchangeRefuseRequest() *TmallExchangeRefuseRequest{
    return &TmallExchangeRefuseRequest{
        Params: model.NewParams(),
    }
}

func (r TmallExchangeRefuseRequest) GetApiMethodName() string {
    return "tmall.exchange.refuse"
}

func (r TmallExchangeRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallExchangeRefuseRequest) SetLeaveMessagePics(leaveMessagePics []*model.File) error {
    r.leaveMessagePics = leaveMessagePics
    r.Set("leave_message_pics", leaveMessagePics)
    return nil
}

func (r TmallExchangeRefuseRequest) GetLeaveMessagePics() []*model.File {
    return r.leaveMessagePics
}

func (r *TmallExchangeRefuseRequest) SetLeaveMessage(leaveMessage string) error {
    r.leaveMessage = leaveMessage
    r.Set("leave_message", leaveMessage)
    return nil
}

func (r TmallExchangeRefuseRequest) GetLeaveMessage() string {
    return r.leaveMessage
}

func (r *TmallExchangeRefuseRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

func (r TmallExchangeRefuseRequest) GetDisputeId() int64 {
    return r.disputeId
}

func (r *TmallExchangeRefuseRequest) SetSellerRefuseReasonId(sellerRefuseReasonId int64) error {
    r.sellerRefuseReasonId = sellerRefuseReasonId
    r.Set("seller_refuse_reason_id", sellerRefuseReasonId)
    return nil
}

func (r TmallExchangeRefuseRequest) GetSellerRefuseReasonId() int64 {
    return r.sellerRefuseReasonId
}

func (r *TmallExchangeRefuseRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TmallExchangeRefuseRequest) GetFields() []string {
    return r.fields
}

