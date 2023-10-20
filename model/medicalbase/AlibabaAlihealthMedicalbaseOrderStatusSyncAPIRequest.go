package medicalbase

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest 号源直连订单状态同步接口 API请求
// alibaba.alihealth.medicalbase.order.status.sync
//
// 互联网医院isv批量通过接口批量导入
type AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest struct {
	model.Params
	// 订单信息
	_orderlSyncDTO *OrderlSyncDto
}

// NewAlibabaAlihealthMedicalbaseOrderStatusSyncRequest 初始化AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest对象
func NewAlibabaAlihealthMedicalbaseOrderStatusSyncRequest() *AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest {
	return &AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) Reset() {
	r._orderlSyncDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.order.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderlSyncDTO is OrderlSyncDTO Setter
// 订单信息
func (r *AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) SetOrderlSyncDTO(_orderlSyncDTO *OrderlSyncDto) error {
	r._orderlSyncDTO = _orderlSyncDTO
	r.Set("orderl_sync_d_t_o", _orderlSyncDTO)
	return nil
}

// GetOrderlSyncDTO OrderlSyncDTO Getter
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) GetOrderlSyncDTO() *OrderlSyncDto {
	return r._orderlSyncDTO
}

var poolAlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalbaseOrderStatusSyncRequest()
	},
}

// GetAlibabaAlihealthMedicalbaseOrderStatusSyncRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest
func GetAlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest() *AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest {
	return poolAlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest.Get().(*AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest 将 AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest(v *AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest.Put(v)
}
