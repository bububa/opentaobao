package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
接收中央随机化系统和临床研究机构的绑定确认状态 APIRequest
alibaba.alihealth.drugcode.center.receive.bound.status

临床用药试验-接收中央随机化系统和临床研究机构的绑定确认状态
*/
type AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest struct {
    model.Params

    // 项目id
    projectId   int64 

    // 临床研究机构id
    hospitalRefEntId   string 

    // 状态 4:绑定成功 5:绑定失败
    status   int64 

    // 中央随机化系统id
    centerRandomSysId   string 

}

func NewAlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest() *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest{
    return &AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.center.receive.bound.status"
}

func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) SetProjectId(projectId int64) error {
    r.projectId = projectId
    r.Set("project_id", projectId)
    return nil
}

func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetProjectId() int64 {
    return r.projectId
}

func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) SetHospitalRefEntId(hospitalRefEntId string) error {
    r.hospitalRefEntId = hospitalRefEntId
    r.Set("hospital_ref_ent_id", hospitalRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetHospitalRefEntId() string {
    return r.hospitalRefEntId
}

func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetStatus() int64 {
    return r.status
}

func (r *AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) SetCenterRandomSysId(centerRandomSysId string) error {
    r.centerRandomSysId = centerRandomSysId
    r.Set("center_random_sys_id", centerRandomSysId)
    return nil
}

func (r AlibabaAlihealthDrugcodeCenterReceiveBoundStatusRequest) GetCenterRandomSysId() string {
    return r.centerRandomSysId
}

