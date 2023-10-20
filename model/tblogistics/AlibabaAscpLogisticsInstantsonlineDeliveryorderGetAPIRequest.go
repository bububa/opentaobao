package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest 同城配送在线下单获取配送单 API请求
// alibaba.ascp.logistics.instantsonline.deliveryorder.get
//
// 同城配送在线下单获取配送单
type AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest struct {
	model.Params
	// ERP单号
	_outOrderId string
}

// NewAlibabaAscpLogisticsInstantsonlineDeliveryorderGetRequest 初始化AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest对象
func NewAlibabaAscpLogisticsInstantsonlineDeliveryorderGetRequest() *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest {
	return &AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest) Reset() {
	r._outOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.instantsonline.deliveryorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutOrderId is OutOrderId Setter
// ERP单号
func (r *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

var poolAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpLogisticsInstantsonlineDeliveryorderGetRequest()
	},
}

// GetAlibabaAscpLogisticsInstantsonlineDeliveryorderGetRequest 从 sync.Pool 获取 AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest
func GetAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest() *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest {
	return poolAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest.Get().(*AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest)
}

// ReleaseAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest 将 AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest(v *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest) {
	v.Reset()
	poolAlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest.Put(v)
}
