package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbtradeordergetAPIRequest 根据交易号获取物流宝订单 API请求
// taobao.wlb.tradeorder.get
//
// 根据交易类型和交易id查询物流宝订单详情
type TaobaowlbtradeordergetAPIRequest struct {
	model.Params
	// 子交易号
	_subTradeId string
	// 指定交易类型的交易号
	_tradeId string
	// 交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易
	_tradeType string
}

// NewTaobaowlbtradeordergetRequest 初始化TaobaowlbtradeordergetAPIRequest对象
func NewTaobaowlbtradeordergetRequest() *TaobaowlbtradeordergetAPIRequest {
	return &TaobaowlbtradeordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbtradeordergetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.tradeorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbtradeordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbtradeordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubTradeId is SubTradeId Setter
// 子交易号
func (r *TaobaowlbtradeordergetAPIRequest) SetSubTradeId(_subTradeId string) error {
	r._subTradeId = _subTradeId
	r.Set("sub_trade_id", _subTradeId)
	return nil
}

// GetSubTradeId SubTradeId Getter
func (r TaobaowlbtradeordergetAPIRequest) GetSubTradeId() string {
	return r._subTradeId
}

// SetTradeId is TradeId Setter
// 指定交易类型的交易号
func (r *TaobaowlbtradeordergetAPIRequest) SetTradeId(_tradeId string) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaowlbtradeordergetAPIRequest) GetTradeId() string {
	return r._tradeId
}

// SetTradeType is TradeType Setter
// 交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易
func (r *TaobaowlbtradeordergetAPIRequest) SetTradeType(_tradeType string) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TaobaowlbtradeordergetAPIRequest) GetTradeType() string {
	return r._tradeType
}
