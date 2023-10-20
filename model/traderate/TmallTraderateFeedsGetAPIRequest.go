package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTraderateFeedsGetAPIRequest 查询子订单对应的评价、追评以及语义标签 API请求
// tmall.traderate.feeds.get
//
// 通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签
type TmallTraderateFeedsGetAPIRequest struct {
	model.Params
	// 交易子订单ID
	_childTradeId int64
}

// NewTmallTraderateFeedsGetRequest 初始化TmallTraderateFeedsGetAPIRequest对象
func NewTmallTraderateFeedsGetRequest() *TmallTraderateFeedsGetAPIRequest {
	return &TmallTraderateFeedsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraderateFeedsGetAPIRequest) GetApiMethodName() string {
	return "tmall.traderate.feeds.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraderateFeedsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTraderateFeedsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChildTradeId is ChildTradeId Setter
// 交易子订单ID
func (r *TmallTraderateFeedsGetAPIRequest) SetChildTradeId(_childTradeId int64) error {
	r._childTradeId = _childTradeId
	r.Set("child_trade_id", _childTradeId)
	return nil
}

// GetChildTradeId ChildTradeId Getter
func (r TmallTraderateFeedsGetAPIRequest) GetChildTradeId() int64 {
	return r._childTradeId
}
