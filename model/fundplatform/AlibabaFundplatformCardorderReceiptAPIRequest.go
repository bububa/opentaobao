package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformcardorderreceiptAPIRequest 通知确认收货 API请求
// alibaba.fundplatform.cardorder.receipt
//
// 告知卡商这一批储值卡已经被用户确认收货
type AlibabafundplatformcardorderreceiptAPIRequest struct {
	model.Params
	// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
	_ownSign string
	// 通知制卡成功的子卡子单号
	_cardOrderId int64
}

// NewAlibabafundplatformcardorderreceiptRequest 初始化AlibabafundplatformcardorderreceiptAPIRequest对象
func NewAlibabafundplatformcardorderreceiptRequest() *AlibabafundplatformcardorderreceiptAPIRequest {
	return &AlibabafundplatformcardorderreceiptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformcardorderreceiptAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorder.receipt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformcardorderreceiptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformcardorderreceiptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOwnSign is OwnSign Setter
// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
func (r *AlibabafundplatformcardorderreceiptAPIRequest) SetOwnSign(_ownSign string) error {
	r._ownSign = _ownSign
	r.Set("own_sign", _ownSign)
	return nil
}

// GetOwnSign OwnSign Getter
func (r AlibabafundplatformcardorderreceiptAPIRequest) GetOwnSign() string {
	return r._ownSign
}

// SetCardOrderId is CardOrderId Setter
// 通知制卡成功的子卡子单号
func (r *AlibabafundplatformcardorderreceiptAPIRequest) SetCardOrderId(_cardOrderId int64) error {
	r._cardOrderId = _cardOrderId
	r.Set("card_order_id", _cardOrderId)
	return nil
}

// GetCardOrderId CardOrderId Getter
func (r AlibabafundplatformcardorderreceiptAPIRequest) GetCardOrderId() int64 {
	return r._cardOrderId
}
