package fundplatform

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardorderReceiptAPIRequest 通知确认收货 API请求
// alibaba.fundplatform.cardorder.receipt
//
// 告知卡商这一批储值卡已经被用户确认收货
type AlibabaFundplatformCardorderReceiptAPIRequest struct {
	model.Params
	// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
	_ownSign string
	// 通知制卡成功的子卡子单号
	_cardOrderId int64
}

// NewAlibabaFundplatformCardorderReceiptRequest 初始化AlibabaFundplatformCardorderReceiptAPIRequest对象
func NewAlibabaFundplatformCardorderReceiptRequest() *AlibabaFundplatformCardorderReceiptAPIRequest {
	return &AlibabaFundplatformCardorderReceiptAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFundplatformCardorderReceiptAPIRequest) Reset() {
	r._ownSign = ""
	r._cardOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderReceiptAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorder.receipt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderReceiptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFundplatformCardorderReceiptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOwnSign is OwnSign Setter
// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
func (r *AlibabaFundplatformCardorderReceiptAPIRequest) SetOwnSign(_ownSign string) error {
	r._ownSign = _ownSign
	r.Set("own_sign", _ownSign)
	return nil
}

// GetOwnSign OwnSign Getter
func (r AlibabaFundplatformCardorderReceiptAPIRequest) GetOwnSign() string {
	return r._ownSign
}

// SetCardOrderId is CardOrderId Setter
// 通知制卡成功的子卡子单号
func (r *AlibabaFundplatformCardorderReceiptAPIRequest) SetCardOrderId(_cardOrderId int64) error {
	r._cardOrderId = _cardOrderId
	r.Set("card_order_id", _cardOrderId)
	return nil
}

// GetCardOrderId CardOrderId Getter
func (r AlibabaFundplatformCardorderReceiptAPIRequest) GetCardOrderId() int64 {
	return r._cardOrderId
}

var poolAlibabaFundplatformCardorderReceiptAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFundplatformCardorderReceiptRequest()
	},
}

// GetAlibabaFundplatformCardorderReceiptRequest 从 sync.Pool 获取 AlibabaFundplatformCardorderReceiptAPIRequest
func GetAlibabaFundplatformCardorderReceiptAPIRequest() *AlibabaFundplatformCardorderReceiptAPIRequest {
	return poolAlibabaFundplatformCardorderReceiptAPIRequest.Get().(*AlibabaFundplatformCardorderReceiptAPIRequest)
}

// ReleaseAlibabaFundplatformCardorderReceiptAPIRequest 将 AlibabaFundplatformCardorderReceiptAPIRequest 放入 sync.Pool
func ReleaseAlibabaFundplatformCardorderReceiptAPIRequest(v *AlibabaFundplatformCardorderReceiptAPIRequest) {
	v.Reset()
	poolAlibabaFundplatformCardorderReceiptAPIRequest.Put(v)
}
