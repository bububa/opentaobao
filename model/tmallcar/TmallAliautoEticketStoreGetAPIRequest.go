package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoEticketStoreGetAPIRequest 查询电子凭证对应门店信息 API请求
// tmall.aliauto.eticket.store.get
//
// 查询电子凭证对应门店信息
type TmallAliautoEticketStoreGetAPIRequest struct {
	model.Params
	// 订单号列表，最多支持20条
	_orderIds []int64
}

// NewTmallAliautoEticketStoreGetRequest 初始化TmallAliautoEticketStoreGetAPIRequest对象
func NewTmallAliautoEticketStoreGetRequest() *TmallAliautoEticketStoreGetAPIRequest {
	return &TmallAliautoEticketStoreGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoEticketStoreGetAPIRequest) Reset() {
	r._orderIds = r._orderIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoEticketStoreGetAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.eticket.store.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoEticketStoreGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoEticketStoreGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderIds is OrderIds Setter
// 订单号列表，最多支持20条
func (r *TmallAliautoEticketStoreGetAPIRequest) SetOrderIds(_orderIds []int64) error {
	r._orderIds = _orderIds
	r.Set("order_ids", _orderIds)
	return nil
}

// GetOrderIds OrderIds Getter
func (r TmallAliautoEticketStoreGetAPIRequest) GetOrderIds() []int64 {
	return r._orderIds
}

var poolTmallAliautoEticketStoreGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoEticketStoreGetRequest()
	},
}

// GetTmallAliautoEticketStoreGetRequest 从 sync.Pool 获取 TmallAliautoEticketStoreGetAPIRequest
func GetTmallAliautoEticketStoreGetAPIRequest() *TmallAliautoEticketStoreGetAPIRequest {
	return poolTmallAliautoEticketStoreGetAPIRequest.Get().(*TmallAliautoEticketStoreGetAPIRequest)
}

// ReleaseTmallAliautoEticketStoreGetAPIRequest 将 TmallAliautoEticketStoreGetAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoEticketStoreGetAPIRequest(v *TmallAliautoEticketStoreGetAPIRequest) {
	v.Reset()
	poolTmallAliautoEticketStoreGetAPIRequest.Put(v)
}
