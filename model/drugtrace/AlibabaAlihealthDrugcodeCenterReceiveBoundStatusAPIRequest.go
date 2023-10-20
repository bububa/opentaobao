package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest 接收中央随机化系统和临床研究机构的绑定确认状态 API请求
// alibaba.alihealth.drugcode.center.receive.bound.status
//
// 临床用药试验-接收中央随机化系统和临床研究机构的绑定确认状态
type AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest struct {
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

// NewAlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest 初始化AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest对象
func NewAlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest() *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest {
	return &AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) Reset() {
	r._hospitalRefEntId = ""
	r._centerRandomSysId = ""
	r._projectId = 0
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.center.receive.bound.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest()
	},
}

// GetAlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest 从 sync.Pool 获取 AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest
func GetAlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest() *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest {
	return poolAlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest.Get().(*AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest)
}

// ReleaseAlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest 将 AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest(v *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest.Put(v)
}
