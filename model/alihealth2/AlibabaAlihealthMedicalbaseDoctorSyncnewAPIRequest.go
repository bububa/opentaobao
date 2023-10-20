package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalbasedoctorsyncnewAPIRequest 直连医生上传 API请求
// alibaba.alihealth.medicalbase.doctor.syncnew
//
// 直连医生上传
type AlibabaalihealthmedicalbasedoctorsyncnewAPIRequest struct {
	model.Params
	// 医生下架
	_topChannelDoctorSyncDTO *TopChannelDoctorSyncDto
}

// NewAlibabaalihealthmedicalbasedoctorsyncnewRequest 初始化AlibabaalihealthmedicalbasedoctorsyncnewAPIRequest对象
func NewAlibabaalihealthmedicalbasedoctorsyncnewRequest() *AlibabaalihealthmedicalbasedoctorsyncnewAPIRequest {
	return &AlibabaalihealthmedicalbasedoctorsyncnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalbasedoctorsyncnewAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.doctor.syncnew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalbasedoctorsyncnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalbasedoctorsyncnewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopChannelDoctorSyncDTO is TopChannelDoctorSyncDTO Setter
// 医生下架
func (r *AlibabaalihealthmedicalbasedoctorsyncnewAPIRequest) SetTopChannelDoctorSyncDTO(_topChannelDoctorSyncDTO *TopChannelDoctorSyncDto) error {
	r._topChannelDoctorSyncDTO = _topChannelDoctorSyncDTO
	r.Set("top_channel_doctor_sync_d_t_o", _topChannelDoctorSyncDTO)
	return nil
}

// GetTopChannelDoctorSyncDTO TopChannelDoctorSyncDTO Getter
func (r AlibabaalihealthmedicalbasedoctorsyncnewAPIRequest) GetTopChannelDoctorSyncDTO() *TopChannelDoctorSyncDto {
	return r._topChannelDoctorSyncDTO
}
