package lsttrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeShiporderQueryAPIRequest 供应商数据开放--发货单接口 API请求
// alibaba.lst.trade.shiporder.query
//
// 供应商数据开放--发货单接口
type AlibabaLstTradeShiporderQueryAPIRequest struct {
	model.Params
	// 入参
	_paramLstShipOrderQuery *LstShipOrderQuery
}

// NewAlibabaLstTradeShiporderQueryRequest 初始化AlibabaLstTradeShiporderQueryAPIRequest对象
func NewAlibabaLstTradeShiporderQueryRequest() *AlibabaLstTradeShiporderQueryAPIRequest {
	return &AlibabaLstTradeShiporderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstTradeShiporderQueryAPIRequest) Reset() {
	r._paramLstShipOrderQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeShiporderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.shiporder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeShiporderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeShiporderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamLstShipOrderQuery is ParamLstShipOrderQuery Setter
// 入参
func (r *AlibabaLstTradeShiporderQueryAPIRequest) SetParamLstShipOrderQuery(_paramLstShipOrderQuery *LstShipOrderQuery) error {
	r._paramLstShipOrderQuery = _paramLstShipOrderQuery
	r.Set("param_lst_ship_order_query", _paramLstShipOrderQuery)
	return nil
}

// GetParamLstShipOrderQuery ParamLstShipOrderQuery Getter
func (r AlibabaLstTradeShiporderQueryAPIRequest) GetParamLstShipOrderQuery() *LstShipOrderQuery {
	return r._paramLstShipOrderQuery
}

var poolAlibabaLstTradeShiporderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstTradeShiporderQueryRequest()
	},
}

// GetAlibabaLstTradeShiporderQueryRequest 从 sync.Pool 获取 AlibabaLstTradeShiporderQueryAPIRequest
func GetAlibabaLstTradeShiporderQueryAPIRequest() *AlibabaLstTradeShiporderQueryAPIRequest {
	return poolAlibabaLstTradeShiporderQueryAPIRequest.Get().(*AlibabaLstTradeShiporderQueryAPIRequest)
}

// ReleaseAlibabaLstTradeShiporderQueryAPIRequest 将 AlibabaLstTradeShiporderQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstTradeShiporderQueryAPIRequest(v *AlibabaLstTradeShiporderQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstTradeShiporderQueryAPIRequest.Put(v)
}
