package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalbasedoctorstatussyncAPIRequest 挂号医生上下架 API请求
// alibaba.alihealth.medicalbase.doctor.status.sync
//
// 挂号医院上下线
type AlibabaalihealthmedicalbasedoctorstatussyncAPIRequest struct {
	model.Params
	// 医生下架
	_topChannelDoctorSyncDTO *TopChannelDoctorSyncDto
}

// NewAlibabaalihealthmedicalbasedoctorstatussyncRequest 初始化AlibabaalihealthmedicalbasedoctorstatussyncAPIRequest对象
func NewAlibabaalihealthmedicalbasedoctorstatussyncRequest() *AlibabaalihealthmedicalbasedoctorstatussyncAPIRequest {
	return &AlibabaalihealthmedicalbasedoctorstatussyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalbasedoctorstatussyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.doctor.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalbasedoctorstatussyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalbasedoctorstatussyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopChannelDoctorSyncDTO is TopChannelDoctorSyncDTO Setter
// 医生下架
func (r *AlibabaalihealthmedicalbasedoctorstatussyncAPIRequest) SetTopChannelDoctorSyncDTO(_topChannelDoctorSyncDTO *TopChannelDoctorSyncDto) error {
	r._topChannelDoctorSyncDTO = _topChannelDoctorSyncDTO
	r.Set("top_channel_doctor_sync_d_t_o", _topChannelDoctorSyncDTO)
	return nil
}

// GetTopChannelDoctorSyncDTO TopChannelDoctorSyncDTO Getter
func (r AlibabaalihealthmedicalbasedoctorstatussyncAPIRequest) GetTopChannelDoctorSyncDTO() *TopChannelDoctorSyncDto {
	return r._topChannelDoctorSyncDTO
}
