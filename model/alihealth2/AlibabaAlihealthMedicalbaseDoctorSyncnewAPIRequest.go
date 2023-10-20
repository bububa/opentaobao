package alihealth2

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest) Reset() {
	r._topChannelDoctorSyncDTO = nil
	r.Params.ToZero()
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

var poolAlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalbaseDoctorSyncnewRequest()
	},
}

// GetAlibabaAlihealthMedicalbaseDoctorSyncnewRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest
func GetAlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest() *AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest {
	return poolAlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest.Get().(*AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest 将 AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest(v *AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest.Put(v)
}
