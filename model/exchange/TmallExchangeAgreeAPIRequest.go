package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangeagreeAPIRequest 卖家同意换货申请 API请求
// tmall.exchange.agree
//
// 卖家同意换货申请
type TmallexchangeagreeAPIRequest struct {
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

// NewTmallexchangeagreeRequest 初始化TmallexchangeagreeAPIRequest对象
func NewTmallexchangeagreeRequest() *TmallexchangeagreeAPIRequest {
	return &TmallexchangeagreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallexchangeagreeAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallexchangeagreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallexchangeagreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段。当前支持的有 dispute_id, bizorder_id, modified, status
func (r *TmallexchangeagreeAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallexchangeagreeAPIRequest) GetFields() []string {
	return r._fields
}

// SetPost is Post Setter
// 邮政编码
func (r *TmallexchangeagreeAPIRequest) SetPost(_post string) error {
	r._post = _post
	r.Set("post", _post)
	return nil
}

// GetPost Post Getter
func (r TmallexchangeagreeAPIRequest) GetPost() string {
	return r._post
}

// SetLeaveMessage is LeaveMessage Setter
// 卖家留言
func (r *TmallexchangeagreeAPIRequest) SetLeaveMessage(_leaveMessage string) error {
	r._leaveMessage = _leaveMessage
	r.Set("leave_message", _leaveMessage)
	return nil
}

// GetLeaveMessage LeaveMessage Getter
func (r TmallexchangeagreeAPIRequest) GetLeaveMessage() string {
	return r._leaveMessage
}

// SetCompleteAddress is CompleteAddress Setter
// 详细收货地址
func (r *TmallexchangeagreeAPIRequest) SetCompleteAddress(_completeAddress string) error {
	r._completeAddress = _completeAddress
	r.Set("complete_address", _completeAddress)
	return nil
}

// GetCompleteAddress CompleteAddress Getter
func (r TmallexchangeagreeAPIRequest) GetCompleteAddress() string {
	return r._completeAddress
}

// SetMobile is Mobile Setter
// 收货人手机号
func (r *TmallexchangeagreeAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TmallexchangeagreeAPIRequest) GetMobile() string {
	return r._mobile
}

// SetLeaveMessagePics is LeaveMessagePics Setter
// 上传图片举证
func (r *TmallexchangeagreeAPIRequest) SetLeaveMessagePics(_leaveMessagePics *model.File) error {
	r._leaveMessagePics = _leaveMessagePics
	r.Set("leave_message_pics", _leaveMessagePics)
	return nil
}

// GetLeaveMessagePics LeaveMessagePics Getter
func (r TmallexchangeagreeAPIRequest) GetLeaveMessagePics() *model.File {
	return r._leaveMessagePics
}

// SetAddressId is AddressId Setter
// 收货地址id，如需获取请调用该top接口：taobao.logistics.address.search，对应属性为contact_id
func (r *TmallexchangeagreeAPIRequest) SetAddressId(_addressId int64) error {
	r._addressId = _addressId
	r.Set("address_id", _addressId)
	return nil
}

// GetAddressId AddressId Getter
func (r TmallexchangeagreeAPIRequest) GetAddressId() int64 {
	return r._addressId
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallexchangeagreeAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallexchangeagreeAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}
