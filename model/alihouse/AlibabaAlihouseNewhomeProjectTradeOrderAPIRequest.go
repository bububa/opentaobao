package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest 旺铺楼盘和交易商品排序 API请求
// alibaba.alihouse.newhome.project.trade.order
//
// 旺铺楼盘和交易商品排序
type AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest struct {
	model.Params
	// 参数
	_tradeOrder *ProjectTradeOrderDto
}

// NewAlibabaAlihouseNewhomeProjectTradeOrderRequest 初始化AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectTradeOrderRequest() *AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest {
	return &AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest) Reset() {
	r._tradeOrder = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.trade.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeOrder is TradeOrder Setter
// 参数
func (r *AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest) SetTradeOrder(_tradeOrder *ProjectTradeOrderDto) error {
	r._tradeOrder = _tradeOrder
	r.Set("trade_order", _tradeOrder)
	return nil
}

// GetTradeOrder TradeOrder Getter
func (r AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest) GetTradeOrder() *ProjectTradeOrderDto {
	return r._tradeOrder
}

var poolAlibabaAlihouseNewhomeProjectTradeOrderAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectTradeOrderRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectTradeOrderRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest
func GetAlibabaAlihouseNewhomeProjectTradeOrderAPIRequest() *AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectTradeOrderAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectTradeOrderAPIRequest 将 AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectTradeOrderAPIRequest(v *AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectTradeOrderAPIRequest.Put(v)
}
