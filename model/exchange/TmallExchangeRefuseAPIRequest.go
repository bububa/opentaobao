package exchange

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeRefuseAPIRequest 卖家拒绝换货申请 API请求
// tmall.exchange.refuse
//
// 卖家拒绝换货申请
type TmallExchangeRefuseAPIRequest struct {
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

// NewTmallExchangeRefuseRequest 初始化TmallExchangeRefuseAPIRequest对象
func NewTmallExchangeRefuseRequest() *TmallExchangeRefuseAPIRequest {
	return &TmallExchangeRefuseAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallExchangeRefuseAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r._leaveMessage = ""
	r._leaveMessagePics = nil
	r._disputeId = 0
	r._sellerRefuseReasonId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeRefuseAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeRefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallExchangeRefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段。目前支持dispute_id, bizorder_id, modified, status
func (r *TmallExchangeRefuseAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallExchangeRefuseAPIRequest) GetFields() []string {
	return r._fields
}

// SetLeaveMessage is LeaveMessage Setter
// 拒绝换货申请时的留言
func (r *TmallExchangeRefuseAPIRequest) SetLeaveMessage(_leaveMessage string) error {
	r._leaveMessage = _leaveMessage
	r.Set("leave_message", _leaveMessage)
	return nil
}

// GetLeaveMessage LeaveMessage Getter
func (r TmallExchangeRefuseAPIRequest) GetLeaveMessage() string {
	return r._leaveMessage
}

// SetLeaveMessagePics is LeaveMessagePics Setter
// 凭证图片
func (r *TmallExchangeRefuseAPIRequest) SetLeaveMessagePics(_leaveMessagePics *model.File) error {
	r._leaveMessagePics = _leaveMessagePics
	r.Set("leave_message_pics", _leaveMessagePics)
	return nil
}

// GetLeaveMessagePics LeaveMessagePics Getter
func (r TmallExchangeRefuseAPIRequest) GetLeaveMessagePics() *model.File {
	return r._leaveMessagePics
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeRefuseAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallExchangeRefuseAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetSellerRefuseReasonId is SellerRefuseReasonId Setter
// 换货原因对应ID
func (r *TmallExchangeRefuseAPIRequest) SetSellerRefuseReasonId(_sellerRefuseReasonId int64) error {
	r._sellerRefuseReasonId = _sellerRefuseReasonId
	r.Set("seller_refuse_reason_id", _sellerRefuseReasonId)
	return nil
}

// GetSellerRefuseReasonId SellerRefuseReasonId Getter
func (r TmallExchangeRefuseAPIRequest) GetSellerRefuseReasonId() int64 {
	return r._sellerRefuseReasonId
}

var poolTmallExchangeRefuseAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallExchangeRefuseRequest()
	},
}

// GetTmallExchangeRefuseRequest 从 sync.Pool 获取 TmallExchangeRefuseAPIRequest
func GetTmallExchangeRefuseAPIRequest() *TmallExchangeRefuseAPIRequest {
	return poolTmallExchangeRefuseAPIRequest.Get().(*TmallExchangeRefuseAPIRequest)
}

// ReleaseTmallExchangeRefuseAPIRequest 将 TmallExchangeRefuseAPIRequest 放入 sync.Pool
func ReleaseTmallExchangeRefuseAPIRequest(v *TmallExchangeRefuseAPIRequest) {
	v.Reset()
	poolTmallExchangeRefuseAPIRequest.Put(v)
}
