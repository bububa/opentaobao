package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundstatusgetAPIRequest 查询退款状态 API请求
// taobao.refund.status.get
//
// 根据主订单或者子订单id查询退款状态,入参传入主订单或者子订单号1.如果当传入子订单时，返回子订单最后一笔退款单的状态,如果子订单申请退款退款返回空list.2.如果传传入主订单，则返回主订单下所有子订单的最后一笔退款单状态，如果对应的子订单没有生成退款单，则对应子订单对应数据返回。
type TaobaorefundstatusgetAPIRequest struct {
	model.Params
	// 入参对象
	_queryParam *RefundQueryByOrderIdRequest
}

// NewTaobaorefundstatusgetRequest 初始化TaobaorefundstatusgetAPIRequest对象
func NewTaobaorefundstatusgetRequest() *TaobaorefundstatusgetAPIRequest {
	return &TaobaorefundstatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorefundstatusgetAPIRequest) GetApiMethodName() string {
	return "taobao.refund.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorefundstatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorefundstatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryParam is QueryParam Setter
// 入参对象
func (r *TaobaorefundstatusgetAPIRequest) SetQueryParam(_queryParam *RefundQueryByOrderIdRequest) error {
	r._queryParam = _queryParam
	r.Set("query_param", _queryParam)
	return nil
}

// GetQueryParam QueryParam Getter
func (r TaobaorefundstatusgetAPIRequest) GetQueryParam() *RefundQueryByOrderIdRequest {
	return r._queryParam
}
