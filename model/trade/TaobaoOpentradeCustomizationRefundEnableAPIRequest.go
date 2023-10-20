package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradecustomizationrefundenableAPIRequest 定制订单设置允许仅退款 API请求
// taobao.opentrade.customization.refund.enable
//
// 定制订单设置允许仅退款
type TaobaoopentradecustomizationrefundenableAPIRequest struct {
	model.Params
	// 主订单ID
	_tradeId int64
}

// NewTaobaoopentradecustomizationrefundenableRequest 初始化TaobaoopentradecustomizationrefundenableAPIRequest对象
func NewTaobaoopentradecustomizationrefundenableRequest() *TaobaoopentradecustomizationrefundenableAPIRequest {
	return &TaobaoopentradecustomizationrefundenableAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradecustomizationrefundenableAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.customization.refund.enable"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradecustomizationrefundenableAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradecustomizationrefundenableAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeId is TradeId Setter
// 主订单ID
func (r *TaobaoopentradecustomizationrefundenableAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaoopentradecustomizationrefundenableAPIRequest) GetTradeId() int64 {
	return r._tradeId
}
