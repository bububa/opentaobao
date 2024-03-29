package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest 挂号医生上下架 API请求
// alibaba.alihealth.medicalbase.doctor.status.sync
//
// 挂号医院上下线
type AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest struct {
	model.Params
	// 医生下架
	_topChannelDoctorSyncDTO *TopChannelDoctorSyncDto
}

// NewAlibabaAlihealthMedicalbaseDoctorStatusSyncRequest 初始化AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest对象
func NewAlibabaAlihealthMedicalbaseDoctorStatusSyncRequest() *AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest {
	return &AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest) Reset() {
	r._topChannelDoctorSyncDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.doctor.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopChannelDoctorSyncDTO is TopChannelDoctorSyncDTO Setter
// 医生下架
func (r *AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest) SetTopChannelDoctorSyncDTO(_topChannelDoctorSyncDTO *TopChannelDoctorSyncDto) error {
	r._topChannelDoctorSyncDTO = _topChannelDoctorSyncDTO
	r.Set("top_channel_doctor_sync_d_t_o", _topChannelDoctorSyncDTO)
	return nil
}

// GetTopChannelDoctorSyncDTO TopChannelDoctorSyncDTO Getter
func (r AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest) GetTopChannelDoctorSyncDTO() *TopChannelDoctorSyncDto {
	return r._topChannelDoctorSyncDTO
}

var poolAlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalbaseDoctorStatusSyncRequest()
	},
}

// GetAlibabaAlihealthMedicalbaseDoctorStatusSyncRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest
func GetAlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest() *AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest {
	return poolAlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest.Get().(*AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest 将 AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest(v *AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest.Put(v)
}
