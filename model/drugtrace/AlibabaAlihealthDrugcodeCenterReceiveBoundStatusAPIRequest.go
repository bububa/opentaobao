package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest 接收中央随机化系统和临床研究机构的绑定确认状态 API请求
// alibaba.alihealth.drugcode.center.receive.bound.status
//
// 临床用药试验-接收中央随机化系统和临床研究机构的绑定确认状态
type AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest struct {
	model.Params
	// 临床研究机构id
	_hospitalRefEntId string
	// 中央随机化系统id
	_centerRandomSysId string
	// 项目id
	_projectId int64
	// 状态 4:绑定成功 5:绑定失败
	_status int64
}

// NewAlibabaalihealthdrugcodecenterreceiveboundstatusRequest 初始化AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest对象
func NewAlibabaalihealthdrugcodecenterreceiveboundstatusRequest() *AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest {
	return &AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.center.receive.bound.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHospitalRefEntId is HospitalRefEntId Setter
// 临床研究机构id
func (r *AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest) SetHospitalRefEntId(_hospitalRefEntId string) error {
	r._hospitalRefEntId = _hospitalRefEntId
	r.Set("hospital_ref_ent_id", _hospitalRefEntId)
	return nil
}

// GetHospitalRefEntId HospitalRefEntId Getter
func (r AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest) GetHospitalRefEntId() string {
	return r._hospitalRefEntId
}

// SetCenterRandomSysId is CenterRandomSysId Setter
// 中央随机化系统id
func (r *AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest) SetCenterRandomSysId(_centerRandomSysId string) error {
	r._centerRandomSysId = _centerRandomSysId
	r.Set("center_random_sys_id", _centerRandomSysId)
	return nil
}

// GetCenterRandomSysId CenterRandomSysId Getter
func (r AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest) GetCenterRandomSysId() string {
	return r._centerRandomSysId
}

// SetProjectId is ProjectId Setter
// 项目id
func (r *AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest) SetProjectId(_projectId int64) error {
	r._projectId = _projectId
	r.Set("project_id", _projectId)
	return nil
}

// GetProjectId ProjectId Getter
func (r AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest) GetProjectId() int64 {
	return r._projectId
}

// SetStatus is Status Setter
// 状态 4:绑定成功 5:绑定失败
func (r *AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest) GetStatus() int64 {
	return r._status
}
