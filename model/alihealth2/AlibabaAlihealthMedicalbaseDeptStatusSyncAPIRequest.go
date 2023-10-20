package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalbasedeptstatussyncAPIRequest 挂号科室上下线 API请求
// alibaba.alihealth.medicalbase.dept.status.sync
//
// 挂号医院上下线
type AlibabaalihealthmedicalbasedeptstatussyncAPIRequest struct {
	model.Params
	// 科室状态同步
	_topChannelDeptSyncDTO *TopChannelDeptSyncDto
}

// NewAlibabaalihealthmedicalbasedeptstatussyncRequest 初始化AlibabaalihealthmedicalbasedeptstatussyncAPIRequest对象
func NewAlibabaalihealthmedicalbasedeptstatussyncRequest() *AlibabaalihealthmedicalbasedeptstatussyncAPIRequest {
	return &AlibabaalihealthmedicalbasedeptstatussyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalbasedeptstatussyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.dept.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalbasedeptstatussyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalbasedeptstatussyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopChannelDeptSyncDTO is TopChannelDeptSyncDTO Setter
// 科室状态同步
func (r *AlibabaalihealthmedicalbasedeptstatussyncAPIRequest) SetTopChannelDeptSyncDTO(_topChannelDeptSyncDTO *TopChannelDeptSyncDto) error {
	r._topChannelDeptSyncDTO = _topChannelDeptSyncDTO
	r.Set("top_channel_dept_sync_d_t_o", _topChannelDeptSyncDTO)
	return nil
}

// GetTopChannelDeptSyncDTO TopChannelDeptSyncDTO Getter
func (r AlibabaalihealthmedicalbasedeptstatussyncAPIRequest) GetTopChannelDeptSyncDTO() *TopChannelDeptSyncDto {
	return r._topChannelDeptSyncDTO
}
