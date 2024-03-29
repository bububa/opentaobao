package lstlogistics

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstShiporderQueryAPIRequest) Reset() {
	r._source = ""
	r._outOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstShiporderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.shiporder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstShiporderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstShiporderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaLstShiporderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstShiporderQueryRequest()
	},
}

// GetAlibabaLstShiporderQueryRequest 从 sync.Pool 获取 AlibabaLstShiporderQueryAPIRequest
func GetAlibabaLstShiporderQueryAPIRequest() *AlibabaLstShiporderQueryAPIRequest {
	return poolAlibabaLstShiporderQueryAPIRequest.Get().(*AlibabaLstShiporderQueryAPIRequest)
}

// ReleaseAlibabaLstShiporderQueryAPIRequest 将 AlibabaLstShiporderQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstShiporderQueryAPIRequest(v *AlibabaLstShiporderQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstShiporderQueryAPIRequest.Put(v)
}
