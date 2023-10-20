package lsttrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerOrderListQueryAPIRequest 订单列表查看(卖家视角) API请求
// alibaba.lst.trade.seller.order.list.query
//
// 卖家视角订单查询，查询授权经销商订单列表
type AlibabaLstTradeSellerOrderListQueryAPIRequest struct {
	model.Params
	// 入参
	_param *LstTradeGetSellerOrderListParam
}

// NewAlibabaLstTradeSellerOrderListQueryRequest 初始化AlibabaLstTradeSellerOrderListQueryAPIRequest对象
func NewAlibabaLstTradeSellerOrderListQueryRequest() *AlibabaLstTradeSellerOrderListQueryAPIRequest {
	return &AlibabaLstTradeSellerOrderListQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstTradeSellerOrderListQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOrderListQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.order.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOrderListQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeSellerOrderListQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaLstTradeSellerOrderListQueryAPIRequest) SetParam(_param *LstTradeGetSellerOrderListParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaLstTradeSellerOrderListQueryAPIRequest) GetParam() *LstTradeGetSellerOrderListParam {
	return r._param
}

var poolAlibabaLstTradeSellerOrderListQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstTradeSellerOrderListQueryRequest()
	},
}

// GetAlibabaLstTradeSellerOrderListQueryRequest 从 sync.Pool 获取 AlibabaLstTradeSellerOrderListQueryAPIRequest
func GetAlibabaLstTradeSellerOrderListQueryAPIRequest() *AlibabaLstTradeSellerOrderListQueryAPIRequest {
	return poolAlibabaLstTradeSellerOrderListQueryAPIRequest.Get().(*AlibabaLstTradeSellerOrderListQueryAPIRequest)
}

// ReleaseAlibabaLstTradeSellerOrderListQueryAPIRequest 将 AlibabaLstTradeSellerOrderListQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstTradeSellerOrderListQueryAPIRequest(v *AlibabaLstTradeSellerOrderListQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstTradeSellerOrderListQueryAPIRequest.Put(v)
}
