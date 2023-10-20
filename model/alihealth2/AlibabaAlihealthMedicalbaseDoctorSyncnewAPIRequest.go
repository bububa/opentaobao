package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest 直连医生上传 API请求
// alibaba.alihealth.medicalbase.doctor.syncnew
//
// 直连医生上传
type AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest struct {
	model.Params
	// 医生下架
	_topChannelDoctorSyncDTO *TopChannelDoctorSyncDto
}

// NewAlibabaAlihealthMedicalbaseDoctorSyncnewRequest 初始化AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest对象
func NewAlibabaAlihealthMedicalbaseDoctorSyncnewRequest() *AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest {
	return &AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.doctor.syncnew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopChannelDoctorSyncDTO is TopChannelDoctorSyncDTO Setter
// 医生下架
func (r *AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest) SetTopChannelDoctorSyncDTO(_topChannelDoctorSyncDTO *TopChannelDoctorSyncDto) error {
	r._topChannelDoctorSyncDTO = _topChannelDoctorSyncDTO
	r.Set("top_channel_doctor_sync_d_t_o", _topChannelDoctorSyncDTO)
	return nil
}

// GetTopChannelDoctorSyncDTO TopChannelDoctorSyncDTO Getter
func (r AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest) GetTopChannelDoctorSyncDTO() *TopChannelDoctorSyncDto {
	return r._topChannelDoctorSyncDTO
}
