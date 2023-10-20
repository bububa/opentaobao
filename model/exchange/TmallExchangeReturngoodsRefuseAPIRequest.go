package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangereturngoodsrefuseAPIRequest 卖家拒绝确认收货 API请求
// tmall.exchange.returngoods.refuse
//
// 卖家拒绝买家换货申请
type TmallexchangereturngoodsrefuseAPIRequest struct {
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

// NewTmallexchangereturngoodsrefuseRequest 初始化TmallexchangereturngoodsrefuseAPIRequest对象
func NewTmallexchangereturngoodsrefuseRequest() *TmallexchangereturngoodsrefuseAPIRequest {
	return &TmallexchangereturngoodsrefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallexchangereturngoodsrefuseAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.returngoods.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallexchangereturngoodsrefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallexchangereturngoodsrefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLeaveMessage is LeaveMessage Setter
// 留言说明
func (r *TmallexchangereturngoodsrefuseAPIRequest) SetLeaveMessage(_leaveMessage string) error {
	r._leaveMessage = _leaveMessage
	r.Set("leave_message", _leaveMessage)
	return nil
}

// GetLeaveMessage LeaveMessage Getter
func (r TmallexchangereturngoodsrefuseAPIRequest) GetLeaveMessage() string {
	return r._leaveMessage
}

// SetLeaveMessagePics is LeaveMessagePics Setter
// 凭证图片
func (r *TmallexchangereturngoodsrefuseAPIRequest) SetLeaveMessagePics(_leaveMessagePics *model.File) error {
	r._leaveMessagePics = _leaveMessagePics
	r.Set("leave_message_pics", _leaveMessagePics)
	return nil
}

// GetLeaveMessagePics LeaveMessagePics Getter
func (r TmallexchangereturngoodsrefuseAPIRequest) GetLeaveMessagePics() *model.File {
	return r._leaveMessagePics
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallexchangereturngoodsrefuseAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallexchangereturngoodsrefuseAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetSellerRefuseReasonId is SellerRefuseReasonId Setter
// 拒绝原因ID
func (r *TmallexchangereturngoodsrefuseAPIRequest) SetSellerRefuseReasonId(_sellerRefuseReasonId int64) error {
	r._sellerRefuseReasonId = _sellerRefuseReasonId
	r.Set("seller_refuse_reason_id", _sellerRefuseReasonId)
	return nil
}

// GetSellerRefuseReasonId SellerRefuseReasonId Getter
func (r TmallexchangereturngoodsrefuseAPIRequest) GetSellerRefuseReasonId() int64 {
	return r._sellerRefuseReasonId
}
