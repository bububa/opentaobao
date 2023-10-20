package medicalbase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalbaseorderstatussyncAPIRequest 号源直连订单状态同步接口 API请求
// alibaba.alihealth.medicalbase.order.status.sync
//
// 互联网医院isv批量通过接口批量导入
type AlibabaalihealthmedicalbaseorderstatussyncAPIRequest struct {
	model.Params
	// 订单信息
	_orderlSyncDTO *OrderlSyncDto
}

// NewAlibabaalihealthmedicalbaseorderstatussyncRequest 初始化AlibabaalihealthmedicalbaseorderstatussyncAPIRequest对象
func NewAlibabaalihealthmedicalbaseorderstatussyncRequest() *AlibabaalihealthmedicalbaseorderstatussyncAPIRequest {
	return &AlibabaalihealthmedicalbaseorderstatussyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalbaseorderstatussyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.order.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalbaseorderstatussyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalbaseorderstatussyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderlSyncDTO is OrderlSyncDTO Setter
// 订单信息
func (r *AlibabaalihealthmedicalbaseorderstatussyncAPIRequest) SetOrderlSyncDTO(_orderlSyncDTO *OrderlSyncDto) error {
	r._orderlSyncDTO = _orderlSyncDTO
	r.Set("orderl_sync_d_t_o", _orderlSyncDTO)
	return nil
}

// GetOrderlSyncDTO OrderlSyncDTO Getter
func (r AlibabaalihealthmedicalbaseorderstatussyncAPIRequest) GetOrderlSyncDTO() *OrderlSyncDto {
	return r._orderlSyncDTO
}
