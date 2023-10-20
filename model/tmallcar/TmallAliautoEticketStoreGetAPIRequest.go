package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoeticketstoregetAPIRequest 查询电子凭证对应门店信息 API请求
// tmall.aliauto.eticket.store.get
//
// 查询电子凭证对应门店信息
type TmallaliautoeticketstoregetAPIRequest struct {
	model.Params
	// 订单号列表，最多支持20条
	_orderIds []int64
}

// NewTmallaliautoeticketstoregetRequest 初始化TmallaliautoeticketstoregetAPIRequest对象
func NewTmallaliautoeticketstoregetRequest() *TmallaliautoeticketstoregetAPIRequest {
	return &TmallaliautoeticketstoregetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautoeticketstoregetAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.eticket.store.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautoeticketstoregetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautoeticketstoregetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderIds is OrderIds Setter
// 订单号列表，最多支持20条
func (r *TmallaliautoeticketstoregetAPIRequest) SetOrderIds(_orderIds []int64) error {
	r._orderIds = _orderIds
	r.Set("order_ids", _orderIds)
	return nil
}

// GetOrderIds OrderIds Getter
func (r TmallaliautoeticketstoregetAPIRequest) GetOrderIds() []int64 {
	return r._orderIds
}
