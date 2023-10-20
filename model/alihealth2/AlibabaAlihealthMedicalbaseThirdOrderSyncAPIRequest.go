package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest 第三方订单同步 API请求
// alibaba.alihealth.medicalbase.third.order.sync
//
// 第三方订单同步
type AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest struct {
	model.Params
	// 请求参数
	_orderRequest *MedicalBaseTopRequestDto
}

// NewAlibabaAlihealthMedicalbaseThirdOrderSyncRequest 初始化AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest对象
func NewAlibabaAlihealthMedicalbaseThirdOrderSyncRequest() *AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest {
	return &AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest) Reset() {
	r._orderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.third.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderRequest is OrderRequest Setter
// 请求参数
func (r *AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest) SetOrderRequest(_orderRequest *MedicalBaseTopRequestDto) error {
	r._orderRequest = _orderRequest
	r.Set("order_request", _orderRequest)
	return nil
}

// GetOrderRequest OrderRequest Getter
func (r AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest) GetOrderRequest() *MedicalBaseTopRequestDto {
	return r._orderRequest
}

var poolAlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalbaseThirdOrderSyncRequest()
	},
}

// GetAlibabaAlihealthMedicalbaseThirdOrderSyncRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest
func GetAlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest() *AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest {
	return poolAlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest.Get().(*AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest 将 AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest(v *AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest.Put(v)
}
