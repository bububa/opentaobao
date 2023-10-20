package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest 挂号科室上下线 API请求
// alibaba.alihealth.medicalbase.dept.status.sync
//
// 挂号医院上下线
type AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest struct {
	model.Params
	// 科室状态同步
	_topChannelDeptSyncDTO *TopChannelDeptSyncDto
}

// NewAlibabaAlihealthMedicalbaseDeptStatusSyncRequest 初始化AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest对象
func NewAlibabaAlihealthMedicalbaseDeptStatusSyncRequest() *AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest {
	return &AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest) Reset() {
	r._topChannelDeptSyncDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.dept.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopChannelDeptSyncDTO is TopChannelDeptSyncDTO Setter
// 科室状态同步
func (r *AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest) SetTopChannelDeptSyncDTO(_topChannelDeptSyncDTO *TopChannelDeptSyncDto) error {
	r._topChannelDeptSyncDTO = _topChannelDeptSyncDTO
	r.Set("top_channel_dept_sync_d_t_o", _topChannelDeptSyncDTO)
	return nil
}

// GetTopChannelDeptSyncDTO TopChannelDeptSyncDTO Getter
func (r AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest) GetTopChannelDeptSyncDTO() *TopChannelDeptSyncDto {
	return r._topChannelDeptSyncDTO
}

var poolAlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalbaseDeptStatusSyncRequest()
	},
}

// GetAlibabaAlihealthMedicalbaseDeptStatusSyncRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest
func GetAlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest() *AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest {
	return poolAlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest.Get().(*AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest 将 AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest(v *AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest.Put(v)
}
