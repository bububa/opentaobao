package alihealth2

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.dept.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
