package lsttrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeOrderGetAPIRequest 零售通交易订单查询--品牌商视角 API请求
// alibaba.lst.trade.order.get
//
// 根据订单id查询零售通交易订单
type AlibabaLstTradeOrderGetAPIRequest struct {
	model.Params
	// 主订单id
	_mainOrderId int64
	// 子订单id
	_subOrderId int64
}

// NewAlibabaLstTradeOrderGetRequest 初始化AlibabaLstTradeOrderGetAPIRequest对象
func NewAlibabaLstTradeOrderGetRequest() *AlibabaLstTradeOrderGetAPIRequest {
	return &AlibabaLstTradeOrderGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstTradeOrderGetAPIRequest) Reset() {
	r._mainOrderId = 0
	r._subOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 主订单id
func (r *AlibabaLstTradeOrderGetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r AlibabaLstTradeOrderGetAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetSubOrderId is SubOrderId Setter
// 子订单id
func (r *AlibabaLstTradeOrderGetAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r AlibabaLstTradeOrderGetAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}

var poolAlibabaLstTradeOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstTradeOrderGetRequest()
	},
}

// GetAlibabaLstTradeOrderGetRequest 从 sync.Pool 获取 AlibabaLstTradeOrderGetAPIRequest
func GetAlibabaLstTradeOrderGetAPIRequest() *AlibabaLstTradeOrderGetAPIRequest {
	return poolAlibabaLstTradeOrderGetAPIRequest.Get().(*AlibabaLstTradeOrderGetAPIRequest)
}

// ReleaseAlibabaLstTradeOrderGetAPIRequest 将 AlibabaLstTradeOrderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstTradeOrderGetAPIRequest(v *AlibabaLstTradeOrderGetAPIRequest) {
	v.Reset()
	poolAlibabaLstTradeOrderGetAPIRequest.Put(v)
}
