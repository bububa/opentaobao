package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradesSoldQueryAPIRequest
根据收件人信息查询交易单号 API请求
taobao.trades.sold.query

根据收件人信息查询交易单号。 */
type TaobaoTradesSoldQueryAPIRequest struct {
	model.Params
	// 查询条件列表，多个条件之间是OR关系，最多支持20个。receiver_name、receiver_mobile、receiver_phone至少有一个值不为空。
	_queryList []OrderQuery
}

// NewTaobaoTradesSoldQueryRequest 初始化TaobaoTradesSoldQueryAPIRequest对象
func NewTaobaoTradesSoldQueryRequest() *TaobaoTradesSoldQueryAPIRequest {
	return &TaobaoTradesSoldQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradesSoldQueryAPIRequest) GetApiMethodName() string {
	return "taobao.trades.sold.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradesSoldQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is QueryList Setter
// 查询条件列表，多个条件之间是OR关系，最多支持20个。receiver_name、receiver_mobile、receiver_phone至少有一个值不为空。
func (r *TaobaoTradesSoldQueryAPIRequest) SetQueryList(_queryList []OrderQuery) error {
	r._queryList = _queryList
	r.Set("query_list", _queryList)
	return nil
}

// Get QueryList Getter
func (r TaobaoTradesSoldQueryAPIRequest) GetQueryList() []OrderQuery {
	return r._queryList
}
