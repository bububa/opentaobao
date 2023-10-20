package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeCustomizationRefundEnableAPIRequest 定制订单设置允许仅退款 API请求
// taobao.opentrade.customization.refund.enable
//
// 定制订单设置允许仅退款
type TaobaoOpentradeCustomizationRefundEnableAPIRequest struct {
	model.Params
	// 主订单ID
	_tradeId int64
}

// NewTaobaoOpentradeCustomizationRefundEnableRequest 初始化TaobaoOpentradeCustomizationRefundEnableAPIRequest对象
func NewTaobaoOpentradeCustomizationRefundEnableRequest() *TaobaoOpentradeCustomizationRefundEnableAPIRequest {
	return &TaobaoOpentradeCustomizationRefundEnableAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpentradeCustomizationRefundEnableAPIRequest) Reset() {
	r._tradeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeCustomizationRefundEnableAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.customization.refund.enable"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeCustomizationRefundEnableAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpentradeCustomizationRefundEnableAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeId is TradeId Setter
// 主订单ID
func (r *TaobaoOpentradeCustomizationRefundEnableAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaoOpentradeCustomizationRefundEnableAPIRequest) GetTradeId() int64 {
	return r._tradeId
}

var poolTaobaoOpentradeCustomizationRefundEnableAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpentradeCustomizationRefundEnableRequest()
	},
}

// GetTaobaoOpentradeCustomizationRefundEnableRequest 从 sync.Pool 获取 TaobaoOpentradeCustomizationRefundEnableAPIRequest
func GetTaobaoOpentradeCustomizationRefundEnableAPIRequest() *TaobaoOpentradeCustomizationRefundEnableAPIRequest {
	return poolTaobaoOpentradeCustomizationRefundEnableAPIRequest.Get().(*TaobaoOpentradeCustomizationRefundEnableAPIRequest)
}

// ReleaseTaobaoOpentradeCustomizationRefundEnableAPIRequest 将 TaobaoOpentradeCustomizationRefundEnableAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpentradeCustomizationRefundEnableAPIRequest(v *TaobaoOpentradeCustomizationRefundEnableAPIRequest) {
	v.Reset()
	poolTaobaoOpentradeCustomizationRefundEnableAPIRequest.Put(v)
}
