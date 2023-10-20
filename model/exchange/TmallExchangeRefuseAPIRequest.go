package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangerefuseAPIRequest 卖家拒绝换货申请 API请求
// tmall.exchange.refuse
//
// 卖家拒绝换货申请
type TmallexchangerefuseAPIRequest struct {
	model.Params
	// 返回字段。目前支持dispute_id, bizorder_id, modified, status
	_fields []string
	// 拒绝换货申请时的留言
	_leaveMessage string
	// 凭证图片
	_leaveMessagePics *model.File
	// 换货单号ID
	_disputeId int64
	// 换货原因对应ID
	_sellerRefuseReasonId int64
}

// NewTmallexchangerefuseRequest 初始化TmallexchangerefuseAPIRequest对象
func NewTmallexchangerefuseRequest() *TmallexchangerefuseAPIRequest {
	return &TmallexchangerefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallexchangerefuseAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallexchangerefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallexchangerefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段。目前支持dispute_id, bizorder_id, modified, status
func (r *TmallexchangerefuseAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallexchangerefuseAPIRequest) GetFields() []string {
	return r._fields
}

// SetLeaveMessage is LeaveMessage Setter
// 拒绝换货申请时的留言
func (r *TmallexchangerefuseAPIRequest) SetLeaveMessage(_leaveMessage string) error {
	r._leaveMessage = _leaveMessage
	r.Set("leave_message", _leaveMessage)
	return nil
}

// GetLeaveMessage LeaveMessage Getter
func (r TmallexchangerefuseAPIRequest) GetLeaveMessage() string {
	return r._leaveMessage
}

// SetLeaveMessagePics is LeaveMessagePics Setter
// 凭证图片
func (r *TmallexchangerefuseAPIRequest) SetLeaveMessagePics(_leaveMessagePics *model.File) error {
	r._leaveMessagePics = _leaveMessagePics
	r.Set("leave_message_pics", _leaveMessagePics)
	return nil
}

// GetLeaveMessagePics LeaveMessagePics Getter
func (r TmallexchangerefuseAPIRequest) GetLeaveMessagePics() *model.File {
	return r._leaveMessagePics
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallexchangerefuseAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallexchangerefuseAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetSellerRefuseReasonId is SellerRefuseReasonId Setter
// 换货原因对应ID
func (r *TmallexchangerefuseAPIRequest) SetSellerRefuseReasonId(_sellerRefuseReasonId int64) error {
	r._sellerRefuseReasonId = _sellerRefuseReasonId
	r.Set("seller_refuse_reason_id", _sellerRefuseReasonId)
	return nil
}

// GetSellerRefuseReasonId SellerRefuseReasonId Getter
func (r TmallexchangerefuseAPIRequest) GetSellerRefuseReasonId() int64 {
	return r._sellerRefuseReasonId
}
