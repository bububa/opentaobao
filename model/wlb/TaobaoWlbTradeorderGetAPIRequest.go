package wlb

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbTradeorderGetAPIRequest 根据交易号获取物流宝订单 API请求
// taobao.wlb.tradeorder.get
//
// 根据交易类型和交易id查询物流宝订单详情
type TaobaoWlbTradeorderGetAPIRequest struct {
	model.Params
	// 子交易号
	_subTradeId string
	// 指定交易类型的交易号
	_tradeId string
	// 交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易
	_tradeType string
}

// NewTaobaoWlbTradeorderGetRequest 初始化TaobaoWlbTradeorderGetAPIRequest对象
func NewTaobaoWlbTradeorderGetRequest() *TaobaoWlbTradeorderGetAPIRequest {
	return &TaobaoWlbTradeorderGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbTradeorderGetAPIRequest) Reset() {
	r._subTradeId = ""
	r._tradeId = ""
	r._tradeType = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbTradeorderGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.tradeorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbTradeorderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbTradeorderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubTradeId is SubTradeId Setter
// 子交易号
func (r *TaobaoWlbTradeorderGetAPIRequest) SetSubTradeId(_subTradeId string) error {
	r._subTradeId = _subTradeId
	r.Set("sub_trade_id", _subTradeId)
	return nil
}

// GetSubTradeId SubTradeId Getter
func (r TaobaoWlbTradeorderGetAPIRequest) GetSubTradeId() string {
	return r._subTradeId
}

// SetTradeId is TradeId Setter
// 指定交易类型的交易号
func (r *TaobaoWlbTradeorderGetAPIRequest) SetTradeId(_tradeId string) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaoWlbTradeorderGetAPIRequest) GetTradeId() string {
	return r._tradeId
}

// SetTradeType is TradeType Setter
// 交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易
func (r *TaobaoWlbTradeorderGetAPIRequest) SetTradeType(_tradeType string) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TaobaoWlbTradeorderGetAPIRequest) GetTradeType() string {
	return r._tradeType
}

var poolTaobaoWlbTradeorderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbTradeorderGetRequest()
	},
}

// GetTaobaoWlbTradeorderGetRequest 从 sync.Pool 获取 TaobaoWlbTradeorderGetAPIRequest
func GetTaobaoWlbTradeorderGetAPIRequest() *TaobaoWlbTradeorderGetAPIRequest {
	return poolTaobaoWlbTradeorderGetAPIRequest.Get().(*TaobaoWlbTradeorderGetAPIRequest)
}

// ReleaseTaobaoWlbTradeorderGetAPIRequest 将 TaobaoWlbTradeorderGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbTradeorderGetAPIRequest(v *TaobaoWlbTradeorderGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbTradeorderGetAPIRequest.Put(v)
}
