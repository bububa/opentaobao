package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
接收中央随机化系统和临床研究机构的绑定确认状态 API请求
alibaba.alihealth.drugcode.center.receive.bound.status

临床用药试验-接收中央随机化系统和临床研究机构的绑定确认状态
*/
type AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest struct {
    model.Params
    // 项目id
    _projectId   int64
    // 临床研究机构id
    _hospitalRefEntId   string
    // 状态 4:绑定成功 5:绑定失败
    _status   int64
    // 中央随机化系统id
    _centerRandomSysId   string
}

// 初始化AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest对象
func NewAlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest() *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest{
    return &AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.center.receive.bound.status"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProjectId Setter
// 项目id
func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) SetProjectId(_projectId int64) error {
    r._projectId = _projectId
    r.Set("project_id", _projectId)
    return nil
}

// ProjectId Getter
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetProjectId() int64 {
    return r._projectId
}
// HospitalRefEntId Setter
// 临床研究机构id
func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) SetHospitalRefEntId(_hospitalRefEntId string) error {
    r._hospitalRefEntId = _hospitalRefEntId
    r.Set("hospital_ref_ent_id", _hospitalRefEntId)
    return nil
}

// HospitalRefEntId Getter
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetHospitalRefEntId() string {
    return r._hospitalRefEntId
}
// Status Setter
// 状态 4:绑定成功 5:绑定失败
func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetStatus() int64 {
    return r._status
}
// CenterRandomSysId Setter
// 中央随机化系统id
func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) SetCenterRandomSysId(_centerRandomSysId string) error {
    r._centerRandomSysId = _centerRandomSysId
    r.Set("center_random_sys_id", _centerRandomSysId)
    return nil
}

// CenterRandomSysId Getter
func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetCenterRandomSysId() string {
    return r._centerRandomSysId
}
