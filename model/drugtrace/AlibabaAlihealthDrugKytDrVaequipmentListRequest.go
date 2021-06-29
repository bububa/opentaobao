package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取企业冷链设备信息 API请求
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

// 初始化AlibabaAlihealthDrugKytDrVaequipmentListRequest对象
func NewAlibabaAlihealthDrugKytDrVaequipmentListRequest() *AlibabaAlihealthDrugKytDrVaequipmentListRequest{
    return &AlibabaAlihealthDrugKytDrVaequipmentListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.vaequipment.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 操作企业ID （appkey授权）
func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetRefEntId() string {
    return r.refEntId
}
// TargetRefEntId Setter
// 目标企业ID
func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetTargetRefEntId(targetRefEntId string) error {
    r.targetRefEntId = targetRefEntId
    r.Set("target_ref_ent_id", targetRefEntId)
    return nil
}

// TargetRefEntId Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetTargetRefEntId() string {
    return r.targetRefEntId
}
// EquipmentCodeOrName Setter
// 设备编号或名称
func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetEquipmentCodeOrName(equipmentCodeOrName string) error {
    r.equipmentCodeOrName = equipmentCodeOrName
    r.Set("equipment_code_or_name", equipmentCodeOrName)
    return nil
}

// EquipmentCodeOrName Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetEquipmentCodeOrName() string {
    return r.equipmentCodeOrName
}
// EquipmentType Setter
// 设备类型
func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetEquipmentType(equipmentType string) error {
    r.equipmentType = equipmentType
    r.Set("equipment_type", equipmentType)
    return nil
}

// EquipmentType Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetEquipmentType() string {
    return r.equipmentType
}
// Status Setter
// 设备状态，1：正常；0：停止
func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetStatus() string {
    return r.status
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetPage() int64 {
    return r.page
}
// PageSize Setter
// 每页显示数量
func (r *AlibabaAlihealthDrugKytDrVaequipmentListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListRequest) GetPageSize() int64 {
    return r.pageSize
}
