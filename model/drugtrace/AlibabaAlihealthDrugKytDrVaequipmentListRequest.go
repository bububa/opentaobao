package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取企业冷链设备信息 APIRequest
alibaba.alihealth.drug.kyt.dr.vaequipment.list

获取企业冷链设备信息
*/
type AlibabaAlihealthDrugKytDrVaequipmentListRequest struct {
    model.Params

    // 操作企业ID （appkey授权）
    refEntId   string 

    // 目标企业ID
    targetRefEntId   string 

    // 设备编号或名称
    equipmentCodeOrName   string 

    // 设备类型
    equipmentType   string 

    // 设备状态，1：正常；0：停止
    status   string 

    // 页码
    page   int64 

    // 每页显示数量
    pageSize   int64 

}

func NewAlibabaAlihealthDrugKytDrVaequipmentListRequest() *AlibabaAlihealthDrugKytDrVaequipmentListRequest{
    return &AlibabaAlihealthDrugKytDrVaequipmentListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.vaequipment.list"
}

func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetTargetRefEntId(targetRefEntId string) error {
    r.targetRefEntId = targetRefEntId
    r.Set("target_ref_ent_id", targetRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetTargetRefEntId() string {
    return r.targetRefEntId
}

func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetEquipmentCodeOrName(equipmentCodeOrName string) error {
    r.equipmentCodeOrName = equipmentCodeOrName
    r.Set("equipment_code_or_name", equipmentCodeOrName)
    return nil
}

func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetEquipmentCodeOrName() string {
    return r.equipmentCodeOrName
}

func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetEquipmentType(equipmentType string) error {
    r.equipmentType = equipmentType
    r.Set("equipment_type", equipmentType)
    return nil
}

func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetEquipmentType() string {
    return r.equipmentType
}

func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetStatus() string {
    return r.status
}

func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetPageSize() int64 {
    return r.pageSize
}

