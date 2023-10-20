package aesolution

import (
	"sync"
)

// OrderDto 结构体
type OrderDto struct {
	// product list
	ProductList []OrderProductDto `json:"product_list,omitempty" xml:"product_list>order_product_dto,omitempty"`
	// seller fuller name
	SellerSignerFullname string `json:"seller_signer_fullname,omitempty" xml:"seller_signer_fullname,omitempty"`
	// seller operator login id
	SellerOperatorLoginId string `json:"seller_operator_login_id,omitempty" xml:"seller_operator_login_id,omitempty"`
	// seller login id
	SellerLoginId string `json:"seller_login_id,omitempty" xml:"seller_login_id,omitempty"`
	// pay type: migs: Credit card payments go through the RMB channel; migs: 102MasterCard pays and takes the RMB channel; migs101:Visa Pay and take the RMB channel; pp101: PayPal; mb: MoneyBooker channel; tt101: Bank Transfer payment; wu101: West Union payment; wp101: Visa pays for the US dollar channel; wp102: Mastercard to pay for the US dollar channel; qw101: QIWI payment; cybs101: Visa takes the payment of the CYBS channel; cybs102: Mastercard to pay for CYBS channels; wm101: WebMoney payment; ebanx101: Brazilian Beloto payment;
	PaymentType string `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// order status
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// order detail url
	OrderDetailUrl string `json:"order_detail_url,omitempty" xml:"order_detail_url,omitempty"`
	// logistics status。logistics status。(WAIT_SELLER_SEND_GOODS: Waiting for seller to ship; SELLER_SEND_PART_GOODS: Partial delivery by seller; SELLER_SEND_GOODS: Seller has shipped; BUYER_ACCEPT_GOODS: Buyer has confirmed receipt; NO_LOGISTICS: No logistics transfer)
	LogisticsStatus string `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
	// logistics escrow fee rate
	LogisitcsEscrowFeeRate string `json:"logisitcs_escrow_fee_rate,omitempty" xml:"logisitcs_escrow_fee_rate,omitempty"`
	// Remaining delivery time (minute)
	LeftSendGoodMin string `json:"left_send_good_min,omitempty" xml:"left_send_good_min,omitempty"`
	// Remaining delivery time (hour）
	LeftSendGoodHour string `json:"left_send_good_hour,omitempty" xml:"left_send_good_hour,omitempty"`
	// Remaining delivery time (days)
	LeftSendGoodDay string `json:"left_send_good_day,omitempty" xml:"left_send_good_day,omitempty"`
	// issue status (NO_ISSUE; IN_ISSUE; END_ISSUE)
	IssueStatus string `json:"issue_status,omitempty" xml:"issue_status,omitempty"`
	// Last order update time
	GmtUpdate string `json:"gmt_update,omitempty" xml:"gmt_update,omitempty"`
	// Last order delivery time
	GmtSendGoodsTime string `json:"gmt_send_goods_time,omitempty" xml:"gmt_send_goods_time,omitempty"`
	// order pay time (The gmtPaysuccess field has the same meaning in the order details.)it&#39;s US Pacific time
	GmtPayTime string `json:"gmt_pay_time,omitempty" xml:"gmt_pay_time,omitempty"`
	// order create time,it&#39;s US Pacific time
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// fund status (NOT_PAY; PAY_SUCCESS; WAIT_SELLER_CHECK)
	FundStatus string `json:"fund_status,omitempty" xml:"fund_status,omitempty"`
	// frozen status。(NO_FROZEN:no frozen; IN_FROZEN:in frozen)
	FrozenStatus string `json:"frozen_status,omitempty" xml:"frozen_status,omitempty"`
	// order finished reason
	EndReason string `json:"end_reason,omitempty" xml:"end_reason,omitempty"`
	// buyer full name
	BuyerSignerFullname string `json:"buyer_signer_fullname,omitempty" xml:"buyer_signer_fullname,omitempty"`
	// buyer login id
	BuyerLoginId string `json:"buyer_login_id,omitempty" xml:"buyer_login_id,omitempty"`
	// order type。（AE_COMMON:common type,AE_TRIAL:trial type;AE_RECHARGE:recharge order)
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// pickup type of order
	OfflinePickupType string `json:"offline_pickup_type,omitempty" xml:"offline_pickup_type,omitempty"`
	// The remain time of the current status (negative number indicates the timeout period)
	TimeoutLeftTime int64 `json:"timeout_left_time,omitempty" xml:"timeout_left_time,omitempty"`
	// pay amount
	PayAmount *SimpleMoney `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// order ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// loan amount details
	LoanAmount *SimpleMoney `json:"loan_amount,omitempty" xml:"loan_amount,omitempty"`
	// escrow fee rate
	EscrowFeeRate int64 `json:"escrow_fee_rate,omitempty" xml:"escrow_fee_rate,omitempty"`
	// escrow fee
	EscrowFee *SimpleMoney `json:"escrow_fee,omitempty" xml:"escrow_fee,omitempty"`
	// Whether mobile phone orders
	Phone bool `json:"phone,omitempty" xml:"phone,omitempty"`
	// Have you requested a loan?
	HasRequestLoan bool `json:"has_request_loan,omitempty" xml:"has_request_loan,omitempty"`
}

var poolOrderDto = sync.Pool{
	New: func() any {
		return new(OrderDto)
	},
}

// GetOrderDto() 从对象池中获取OrderDto
func GetOrderDto() *OrderDto {
	return poolOrderDto.Get().(*OrderDto)
}

// ReleaseOrderDto 释放OrderDto
func ReleaseOrderDto(v *OrderDto) {
	v.ProductList = v.ProductList[:0]
	v.SellerSignerFullname = ""
	v.SellerOperatorLoginId = ""
	v.SellerLoginId = ""
	v.PaymentType = ""
	v.OrderStatus = ""
	v.OrderDetailUrl = ""
	v.LogisticsStatus = ""
	v.LogisitcsEscrowFeeRate = ""
	v.LeftSendGoodMin = ""
	v.LeftSendGoodHour = ""
	v.LeftSendGoodDay = ""
	v.IssueStatus = ""
	v.GmtUpdate = ""
	v.GmtSendGoodsTime = ""
	v.GmtPayTime = ""
	v.GmtCreate = ""
	v.FundStatus = ""
	v.FrozenStatus = ""
	v.EndReason = ""
	v.BuyerSignerFullname = ""
	v.BuyerLoginId = ""
	v.BizType = ""
	v.OfflinePickupType = ""
	v.TimeoutLeftTime = 0
	v.PayAmount = nil
	v.OrderId = 0
	v.LoanAmount = nil
	v.EscrowFeeRate = 0
	v.EscrowFee = nil
	v.Phone = false
	v.HasRequestLoan = false
	poolOrderDto.Put(v)
}
