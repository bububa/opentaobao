package lsttrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerOrderDetailQueryAPIRequest 订单详情查看(卖家视角) API请求
// alibaba.lst.trade.seller.order.detail.query
//
// 订单详情查看(卖家视角)
type AlibabaLstTradeSellerOrderDetailQueryAPIRequest struct {
	model.Params
	// 入参
	_param *LstTradeGetSellerOrderListParam
}

// NewAlibabaLstTradeSellerOrderDetailQueryRequest 初始化AlibabaLstTradeSellerOrderDetailQueryAPIRequest对象
func NewAlibabaLstTradeSellerOrderDetailQueryRequest() *AlibabaLstTradeSellerOrderDetailQueryAPIRequest {
	return &AlibabaLstTradeSellerOrderDetailQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstTradeSellerOrderDetailQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOrderDetailQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.order.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOrderDetailQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeSellerOrderDetailQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaLstTradeSellerOrderDetailQueryAPIRequest) SetParam(_param *LstTradeGetSellerOrderListParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaLstTradeSellerOrderDetailQueryAPIRequest) GetParam() *LstTradeGetSellerOrderListParam {
	return r._param
}

var poolAlibabaLstTradeSellerOrderDetailQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstTradeSellerOrderDetailQueryRequest()
	},
}

// GetAlibabaLstTradeSellerOrderDetailQueryRequest 从 sync.Pool 获取 AlibabaLstTradeSellerOrderDetailQueryAPIRequest
func GetAlibabaLstTradeSellerOrderDetailQueryAPIRequest() *AlibabaLstTradeSellerOrderDetailQueryAPIRequest {
	return poolAlibabaLstTradeSellerOrderDetailQueryAPIRequest.Get().(*AlibabaLstTradeSellerOrderDetailQueryAPIRequest)
}

// ReleaseAlibabaLstTradeSellerOrderDetailQueryAPIRequest 将 AlibabaLstTradeSellerOrderDetailQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstTradeSellerOrderDetailQueryAPIRequest(v *AlibabaLstTradeSellerOrderDetailQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstTradeSellerOrderDetailQueryAPIRequest.Put(v)
}
