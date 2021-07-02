package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest 接收中央随机化系统和临床研究机构的绑定确认状态 API请求
// alibaba.alihealth.drugcode.center.receive.bound.status
//
// 临床用药试验-接收中央随机化系统和临床研究机构的绑定确认状态
type AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest struct {
	model.Params
	// 项目id
	_projectId int64
	// 临床研究机构id
	_hospitalRefEntId string
	// 状态 4:绑定成功 5:绑定失败
	_status int64
	// 中央随机化系统id
	_centerRandomSysId string
}

// NewAlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest 初始化AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest对象
func NewAlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest() *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest {
	return &AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.center.receive.bound.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProjectId is ProjectId Setter
// 项目id
func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) SetProjectId(_projectId int64) error {
	r._projectId = _projectId
	r.Set("project_id", _projectId)
	return nil
}

// GetProjectId ProjectId Getter
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) GetProjectId() int64 {
	return r._projectId
}

// SetHospitalRefEntId is HospitalRefEntId Setter
// 临床研究机构id
func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) SetHospitalRefEntId(_hospitalRefEntId string) error {
	r._hospitalRefEntId = _hospitalRefEntId
	r.Set("hospital_ref_ent_id", _hospitalRefEntId)
	return nil
}

// GetHospitalRefEntId HospitalRefEntId Getter
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) GetHospitalRefEntId() string {
	return r._hospitalRefEntId
}

// SetStatus is Status Setter
// 状态 4:绑定成功 5:绑定失败
func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) GetStatus() int64 {
	return r._status
}

// SetCenterRandomSysId is CenterRandomSysId Setter
// 中央随机化系统id
func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) SetCenterRandomSysId(_centerRandomSysId string) error {
	r._centerRandomSysId = _centerRandomSysId
	r.Set("center_random_sys_id", _centerRandomSysId)
	return nil
}

// GetCenterRandomSysId CenterRandomSysId Getter
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) GetCenterRandomSysId() string {
	return r._centerRandomSysId
}
