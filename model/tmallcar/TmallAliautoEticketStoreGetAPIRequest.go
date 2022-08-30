package tmallcar

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoEticketStoreGetAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.eticket.store.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoEticketStoreGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
