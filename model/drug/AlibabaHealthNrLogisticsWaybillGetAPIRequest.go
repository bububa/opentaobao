package drug

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthNrLogisticsWaybillGetAPIRequest 电子面单查询接口 API请求
// alibaba.health.nr.logistics.waybill.get
//
// 商家登录后根据订单号查询物流单号及电子面单信息
type AlibabaHealthNrLogisticsWaybillGetAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewAlibabaHealthNrLogisticsWaybillGetRequest 初始化AlibabaHealthNrLogisticsWaybillGetAPIRequest对象
func NewAlibabaHealthNrLogisticsWaybillGetRequest() *AlibabaHealthNrLogisticsWaybillGetAPIRequest {
	return &AlibabaHealthNrLogisticsWaybillGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHealthNrLogisticsWaybillGetAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrLogisticsWaybillGetAPIRequest) GetApiMethodName() string {
	return "alibaba.health.nr.logistics.waybill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrLogisticsWaybillGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHealthNrLogisticsWaybillGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaHealthNrLogisticsWaybillGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHealthNrLogisticsWaybillGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolAlibabaHealthNrLogisticsWaybillGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHealthNrLogisticsWaybillGetRequest()
	},
}

// GetAlibabaHealthNrLogisticsWaybillGetRequest 从 sync.Pool 获取 AlibabaHealthNrLogisticsWaybillGetAPIRequest
func GetAlibabaHealthNrLogisticsWaybillGetAPIRequest() *AlibabaHealthNrLogisticsWaybillGetAPIRequest {
	return poolAlibabaHealthNrLogisticsWaybillGetAPIRequest.Get().(*AlibabaHealthNrLogisticsWaybillGetAPIRequest)
}

// ReleaseAlibabaHealthNrLogisticsWaybillGetAPIRequest 将 AlibabaHealthNrLogisticsWaybillGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaHealthNrLogisticsWaybillGetAPIRequest(v *AlibabaHealthNrLogisticsWaybillGetAPIRequest) {
	v.Reset()
	poolAlibabaHealthNrLogisticsWaybillGetAPIRequest.Put(v)
}
