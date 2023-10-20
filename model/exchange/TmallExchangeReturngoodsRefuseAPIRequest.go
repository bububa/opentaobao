package exchange

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeReturngoodsRefuseAPIRequest 卖家拒绝确认收货 API请求
// tmall.exchange.returngoods.refuse
//
// 卖家拒绝买家换货申请
type TmallExchangeReturngoodsRefuseAPIRequest struct {
	model.Params
	// 留言说明
	_leaveMessage string
	// 凭证图片
	_leaveMessagePics *model.File
	// 换货单号ID
	_disputeId int64
	// 拒绝原因ID
	_sellerRefuseReasonId int64
}

// NewTmallExchangeReturngoodsRefuseRequest 初始化TmallExchangeReturngoodsRefuseAPIRequest对象
func NewTmallExchangeReturngoodsRefuseRequest() *TmallExchangeReturngoodsRefuseAPIRequest {
	return &TmallExchangeReturngoodsRefuseAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallExchangeReturngoodsRefuseAPIRequest) Reset() {
	r._leaveMessage = ""
	r._leaveMessagePics = nil
	r._disputeId = 0
	r._sellerRefuseReasonId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.returngoods.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLeaveMessage is LeaveMessage Setter
// 留言说明
func (r *TmallExchangeReturngoodsRefuseAPIRequest) SetLeaveMessage(_leaveMessage string) error {
	r._leaveMessage = _leaveMessage
	r.Set("leave_message", _leaveMessage)
	return nil
}

// GetLeaveMessage LeaveMessage Getter
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetLeaveMessage() string {
	return r._leaveMessage
}

// SetLeaveMessagePics is LeaveMessagePics Setter
// 凭证图片
func (r *TmallExchangeReturngoodsRefuseAPIRequest) SetLeaveMessagePics(_leaveMessagePics *model.File) error {
	r._leaveMessagePics = _leaveMessagePics
	r.Set("leave_message_pics", _leaveMessagePics)
	return nil
}

// GetLeaveMessagePics LeaveMessagePics Getter
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetLeaveMessagePics() *model.File {
	return r._leaveMessagePics
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeReturngoodsRefuseAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetSellerRefuseReasonId is SellerRefuseReasonId Setter
// 拒绝原因ID
func (r *TmallExchangeReturngoodsRefuseAPIRequest) SetSellerRefuseReasonId(_sellerRefuseReasonId int64) error {
	r._sellerRefuseReasonId = _sellerRefuseReasonId
	r.Set("seller_refuse_reason_id", _sellerRefuseReasonId)
	return nil
}

// GetSellerRefuseReasonId SellerRefuseReasonId Getter
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetSellerRefuseReasonId() int64 {
	return r._sellerRefuseReasonId
}

var poolTmallExchangeReturngoodsRefuseAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallExchangeReturngoodsRefuseRequest()
	},
}

// GetTmallExchangeReturngoodsRefuseRequest 从 sync.Pool 获取 TmallExchangeReturngoodsRefuseAPIRequest
func GetTmallExchangeReturngoodsRefuseAPIRequest() *TmallExchangeReturngoodsRefuseAPIRequest {
	return poolTmallExchangeReturngoodsRefuseAPIRequest.Get().(*TmallExchangeReturngoodsRefuseAPIRequest)
}

// ReleaseTmallExchangeReturngoodsRefuseAPIRequest 将 TmallExchangeReturngoodsRefuseAPIRequest 放入 sync.Pool
func ReleaseTmallExchangeReturngoodsRefuseAPIRequest(v *TmallExchangeReturngoodsRefuseAPIRequest) {
	v.Reset()
	poolTmallExchangeReturngoodsRefuseAPIRequest.Put(v)
}
