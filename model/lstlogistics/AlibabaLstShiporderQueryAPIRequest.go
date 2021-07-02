package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstShiporderQueryAPIRequest 零售通发货单查询 API请求
// alibaba.lst.shiporder.query
//
// 通过该接口可以查询零售通运保保发货单，并处理相关业务流程。
type AlibabaLstShiporderQueryAPIRequest struct {
	model.Params
	// 零售通
	_source string
	// 订单
	_outOrderId string
}

// NewAlibabaLstShiporderQueryRequest 初始化AlibabaLstShiporderQueryAPIRequest对象
func NewAlibabaLstShiporderQueryRequest() *AlibabaLstShiporderQueryAPIRequest {
	return &AlibabaLstShiporderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstShiporderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.shiporder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstShiporderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSource is Source Setter
// 零售通
func (r *AlibabaLstShiporderQueryAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaLstShiporderQueryAPIRequest) GetSource() string {
	return r._source
}

// SetOutOrderId is OutOrderId Setter
// 订单
func (r *AlibabaLstShiporderQueryAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaLstShiporderQueryAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}
