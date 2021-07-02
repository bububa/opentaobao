package medicalbase

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.order.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
