package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeAgreeAPIRequest 卖家同意换货申请 API请求
// tmall.exchange.agree
//
// 卖家同意换货申请
type TmallExchangeAgreeAPIRequest struct {
	model.Params
	// 返回字段。当前支持的有 dispute_id, bizorder_id, modified, status
	_fields []string
	// 邮政编码
	_post string
	// 卖家留言
	_leaveMessage string
	// 详细收货地址
	_completeAddress string
	// 收货人手机号
	_mobile string
	// 上传图片举证
	_leaveMessagePics *model.File
	// 收货地址id，如需获取请调用该top接口：taobao.logistics.address.search，对应属性为contact_id
	_addressId int64
	// 换货单号ID
	_disputeId int64
}

// NewTmallExchangeAgreeRequest 初始化TmallExchangeAgreeAPIRequest对象
func NewTmallExchangeAgreeRequest() *TmallExchangeAgreeAPIRequest {
	return &TmallExchangeAgreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeAgreeAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeAgreeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 返回字段。当前支持的有 dispute_id, bizorder_id, modified, status
func (r *TmallExchangeAgreeAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallExchangeAgreeAPIRequest) GetFields() []string {
	return r._fields
}

// SetPost is Post Setter
// 邮政编码
func (r *TmallExchangeAgreeAPIRequest) SetPost(_post string) error {
	r._post = _post
	r.Set("post", _post)
	return nil
}

// GetPost Post Getter
func (r TmallExchangeAgreeAPIRequest) GetPost() string {
	return r._post
}

// SetLeaveMessage is LeaveMessage Setter
// 卖家留言
func (r *TmallExchangeAgreeAPIRequest) SetLeaveMessage(_leaveMessage string) error {
	r._leaveMessage = _leaveMessage
	r.Set("leave_message", _leaveMessage)
	return nil
}

// GetLeaveMessage LeaveMessage Getter
func (r TmallExchangeAgreeAPIRequest) GetLeaveMessage() string {
	return r._leaveMessage
}

// SetCompleteAddress is CompleteAddress Setter
// 详细收货地址
func (r *TmallExchangeAgreeAPIRequest) SetCompleteAddress(_completeAddress string) error {
	r._completeAddress = _completeAddress
	r.Set("complete_address", _completeAddress)
	return nil
}

// GetCompleteAddress CompleteAddress Getter
func (r TmallExchangeAgreeAPIRequest) GetCompleteAddress() string {
	return r._completeAddress
}

// SetMobile is Mobile Setter
// 收货人手机号
func (r *TmallExchangeAgreeAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TmallExchangeAgreeAPIRequest) GetMobile() string {
	return r._mobile
}

// SetLeaveMessagePics is LeaveMessagePics Setter
// 上传图片举证
func (r *TmallExchangeAgreeAPIRequest) SetLeaveMessagePics(_leaveMessagePics *model.File) error {
	r._leaveMessagePics = _leaveMessagePics
	r.Set("leave_message_pics", _leaveMessagePics)
	return nil
}

// GetLeaveMessagePics LeaveMessagePics Getter
func (r TmallExchangeAgreeAPIRequest) GetLeaveMessagePics() *model.File {
	return r._leaveMessagePics
}

// SetAddressId is AddressId Setter
// 收货地址id，如需获取请调用该top接口：taobao.logistics.address.search，对应属性为contact_id
func (r *TmallExchangeAgreeAPIRequest) SetAddressId(_addressId int64) error {
	r._addressId = _addressId
	r.Set("address_id", _addressId)
	return nil
}

// GetAddressId AddressId Getter
func (r TmallExchangeAgreeAPIRequest) GetAddressId() int64 {
	return r._addressId
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeAgreeAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallExchangeAgreeAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}
