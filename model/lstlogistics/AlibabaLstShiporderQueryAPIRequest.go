package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstshiporderqueryAPIRequest 零售通发货单查询 API请求
// alibaba.lst.shiporder.query
//
// 通过该接口可以查询零售通运保保发货单，并处理相关业务流程。
type AlibabalstshiporderqueryAPIRequest struct {
	model.Params
	// 零售通
	_source string
	// 订单
	_outOrderId string
}

// NewAlibabalstshiporderqueryRequest 初始化AlibabalstshiporderqueryAPIRequest对象
func NewAlibabalstshiporderqueryRequest() *AlibabalstshiporderqueryAPIRequest {
	return &AlibabalstshiporderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstshiporderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.shiporder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstshiporderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstshiporderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSource is Source Setter
// 零售通
func (r *AlibabalstshiporderqueryAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabalstshiporderqueryAPIRequest) GetSource() string {
	return r._source
}

// SetOutOrderId is OutOrderId Setter
// 订单
func (r *AlibabalstshiporderqueryAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabalstshiporderqueryAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}
