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
type AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest struct {
    model.Params
    // 操作企业ID （appkey授权）
    _refEntId   string
    // 目标企业ID
    _targetRefEntId   string
    // 设备编号或名称
    _equipmentCodeOrName   string
    // 设备类型
    _equipmentType   string
    // 设备状态，1：正常；0：停止
    _status   string
    // 页码
    _page   int64
    // 每页显示数量
    _pageSize   int64
}

// 初始化AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest对象
func NewAlibabaAlihealthDrugKytDrVaequipmentListRequest() *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest{
    return &AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.vaequipment.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 操作企业ID （appkey授权）
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// TargetRefEntId Setter
// 目标企业ID
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetTargetRefEntId(_targetRefEntId string) error {
    r._targetRefEntId = _targetRefEntId
    r.Set("target_ref_ent_id", _targetRefEntId)
    return nil
}

// TargetRefEntId Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetTargetRefEntId() string {
    return r._targetRefEntId
}
// EquipmentCodeOrName Setter
// 设备编号或名称
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetEquipmentCodeOrName(_equipmentCodeOrName string) error {
    r._equipmentCodeOrName = _equipmentCodeOrName
    r.Set("equipment_code_or_name", _equipmentCodeOrName)
    return nil
}

// EquipmentCodeOrName Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetEquipmentCodeOrName() string {
    return r._equipmentCodeOrName
}
// EquipmentType Setter
// 设备类型
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetEquipmentType(_equipmentType string) error {
    r._equipmentType = _equipmentType
    r.Set("equipment_type", _equipmentType)
    return nil
}

// EquipmentType Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetEquipmentType() string {
    return r._equipmentType
}
// Status Setter
// 设备状态，1：正常；0：停止
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetStatus() string {
    return r._status
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 每页显示数量
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
