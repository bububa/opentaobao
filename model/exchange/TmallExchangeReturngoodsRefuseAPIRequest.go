package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeReturngoodsRefuseAPIRequest
卖家拒绝确认收货 API请求
tmall.exchange.returngoods.refuse

卖家拒绝买家换货申请 */
type TmallExchangeReturngoodsRefuseAPIRequest struct {
	model.Params
	// 凭证图片
	_leaveMessagePics *model.File
	// 留言说明
	_leaveMessage string
	// 换货单号ID
	_disputeId int64
	// 拒绝原因ID
	_sellerRefuseReasonId int64
}

// NewTmallExchangeReturngoodsRefuseRequest 初始化TmallExchangeReturngoodsRefuseAPIRequest对象
func NewTmallExchangeReturngoodsRefuseRequest() *TmallExchangeReturngoodsRefuseAPIRequest {
	return &TmallExchangeReturngoodsRefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.returngoods.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LeaveMessagePics Setter
// 凭证图片
func (r *TmallExchangeReturngoodsRefuseAPIRequest) SetLeaveMessagePics(_leaveMessagePics *model.File) error {
	r._leaveMessagePics = _leaveMessagePics
	r.Set("leave_message_pics", _leaveMessagePics)
	return nil
}

// Get LeaveMessagePics Getter
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetLeaveMessagePics() *model.File {
	return r._leaveMessagePics
}

// Set is LeaveMessage Setter
// 留言说明
func (r *TmallExchangeReturngoodsRefuseAPIRequest) SetLeaveMessage(_leaveMessage string) error {
	r._leaveMessage = _leaveMessage
	r.Set("leave_message", _leaveMessage)
	return nil
}

// Get LeaveMessage Getter
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetLeaveMessage() string {
	return r._leaveMessage
}

// Set is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeReturngoodsRefuseAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// Get DisputeId Getter
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// Set is SellerRefuseReasonId Setter
// 拒绝原因ID
func (r *TmallExchangeReturngoodsRefuseAPIRequest) SetSellerRefuseReasonId(_sellerRefuseReasonId int64) error {
	r._sellerRefuseReasonId = _sellerRefuseReasonId
	r.Set("seller_refuse_reason_id", _sellerRefuseReasonId)
	return nil
}

// Get SellerRefuseReasonId Getter
func (r TmallExchangeReturngoodsRefuseAPIRequest) GetSellerRefuseReasonId() int64 {
	return r._sellerRefuseReasonId
}
