package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家同意换货申请 API请求
tmall.exchange.agree

卖家同意换货申请
*/
type TmallExchangeAgreeRequest struct {
    model.Params
    // 邮政编码
    post   string
    // 上传图片举证
    leaveMessagePics   []*model.File
    // 卖家留言
    leaveMessage   string
    // 收货地址id，如需获取请调用该top接口：taobao.logistics.address.search，对应属性为contact_id
    addressId   int64
    // 详细收货地址
    completeAddress   string
    // 换货单号ID
    disputeId   int64
    // 返回字段。当前支持的有 dispute_id, bizorder_id, modified, status
    fields   []string
    // 收货人手机号
    mobile   string
}

// 初始化TmallExchangeAgreeRequest对象
func NewTmallExchangeAgreeRequest() *TmallExchangeAgreeRequest{
    return &TmallExchangeAgreeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallExchangeAgreeRequest) GetApiMethodName() string {
    return "tmall.exchange.agree"
}

// IRequest interface 方法, 获取API参数
func (r TmallExchangeAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Post Setter
// 邮政编码
func (r *TmallExchangeAgreeRequest) SetPost(post string) error {
    r.post = post
    r.Set("post", post)
    return nil
}

// Post Getter
func (r TmallExchangeAgreeRequest) GetPost() string {
    return r.post
}
// LeaveMessagePics Setter
// 上传图片举证
func (r *TmallExchangeAgreeRequest) SetLeaveMessagePics(leaveMessagePics []*model.File) error {
    r.leaveMessagePics = leaveMessagePics
    r.Set("leave_message_pics", leaveMessagePics)
    return nil
}

// LeaveMessagePics Getter
func (r TmallExchangeAgreeRequest) GetLeaveMessagePics() []*model.File {
    return r.leaveMessagePics
}
// LeaveMessage Setter
// 卖家留言
func (r *TmallExchangeAgreeRequest) SetLeaveMessage(leaveMessage string) error {
    r.leaveMessage = leaveMessage
    r.Set("leave_message", leaveMessage)
    return nil
}

// LeaveMessage Getter
func (r TmallExchangeAgreeRequest) GetLeaveMessage() string {
    return r.leaveMessage
}
// AddressId Setter
// 收货地址id，如需获取请调用该top接口：taobao.logistics.address.search，对应属性为contact_id
func (r *TmallExchangeAgreeRequest) SetAddressId(addressId int64) error {
    r.addressId = addressId
    r.Set("address_id", addressId)
    return nil
}

// AddressId Getter
func (r TmallExchangeAgreeRequest) GetAddressId() int64 {
    return r.addressId
}
// CompleteAddress Setter
// 详细收货地址
func (r *TmallExchangeAgreeRequest) SetCompleteAddress(completeAddress string) error {
    r.completeAddress = completeAddress
    r.Set("complete_address", completeAddress)
    return nil
}

// CompleteAddress Getter
func (r TmallExchangeAgreeRequest) GetCompleteAddress() string {
    return r.completeAddress
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeAgreeRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeAgreeRequest) GetDisputeId() int64 {
    return r.disputeId
}
// Fields Setter
// 返回字段。当前支持的有 dispute_id, bizorder_id, modified, status
func (r *TmallExchangeAgreeRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TmallExchangeAgreeRequest) GetFields() []string {
    return r.fields
}
// Mobile Setter
// 收货人手机号
func (r *TmallExchangeAgreeRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r TmallExchangeAgreeRequest) GetMobile() string {
    return r.mobile
}
