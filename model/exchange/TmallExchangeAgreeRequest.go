package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家同意换货申请 APIRequest
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

func NewTmallExchangeAgreeRequest() *TmallExchangeAgreeRequest{
    return &TmallExchangeAgreeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallExchangeAgreeRequest) GetApiMethodName() string {
    return "tmall.exchange.agree"
}

func (r TmallExchangeAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallExchangeAgreeRequest) SetPost(post string) error {
    r.post = post
    r.Set("post", post)
    return nil
}

func (r TmallExchangeAgreeRequest) GetPost() string {
    return r.post
}

func (r *TmallExchangeAgreeRequest) SetLeaveMessagePics(leaveMessagePics []*model.File) error {
    r.leaveMessagePics = leaveMessagePics
    r.Set("leave_message_pics", leaveMessagePics)
    return nil
}

func (r TmallExchangeAgreeRequest) GetLeaveMessagePics() []*model.File {
    return r.leaveMessagePics
}

func (r *TmallExchangeAgreeRequest) SetLeaveMessage(leaveMessage string) error {
    r.leaveMessage = leaveMessage
    r.Set("leave_message", leaveMessage)
    return nil
}

func (r TmallExchangeAgreeRequest) GetLeaveMessage() string {
    return r.leaveMessage
}

func (r *TmallExchangeAgreeRequest) SetAddressId(addressId int64) error {
    r.addressId = addressId
    r.Set("address_id", addressId)
    return nil
}

func (r TmallExchangeAgreeRequest) GetAddressId() int64 {
    return r.addressId
}

func (r *TmallExchangeAgreeRequest) SetCompleteAddress(completeAddress string) error {
    r.completeAddress = completeAddress
    r.Set("complete_address", completeAddress)
    return nil
}

func (r TmallExchangeAgreeRequest) GetCompleteAddress() string {
    return r.completeAddress
}

func (r *TmallExchangeAgreeRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

func (r TmallExchangeAgreeRequest) GetDisputeId() int64 {
    return r.disputeId
}

func (r *TmallExchangeAgreeRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TmallExchangeAgreeRequest) GetFields() []string {
    return r.fields
}

func (r *TmallExchangeAgreeRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r TmallExchangeAgreeRequest) GetMobile() string {
    return r.mobile
}

