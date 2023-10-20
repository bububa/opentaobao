package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest 获取企业冷链设备信息 API请求
// alibaba.alihealth.drug.kyt.dr.vaequipment.list
//
// 获取企业冷链设备信息
type AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest struct {
	model.Params
	// 操作企业ID （appkey授权）
	_refEntId string
	// 目标企业ID
	_targetRefEntId string
	// 设备编号或名称
	_equipmentCodeOrName string
	// 设备类型
	_equipmentType string
	// 设备状态，1：正常；0：停止
	_status string
	// 页码
	_page int64
	// 每页显示数量
	_pageSize int64
}

// NewAlibabaAlihealthDrugKytDrVaequipmentListRequest 初始化AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest对象
func NewAlibabaAlihealthDrugKytDrVaequipmentListRequest() *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest {
	return &AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) Reset() {
	r._refEntId = ""
	r._targetRefEntId = ""
	r._equipmentCodeOrName = ""
	r._equipmentType = ""
	r._status = ""
	r._page = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.vaequipment.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 操作企业ID （appkey授权）
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetTargetRefEntId is TargetRefEntId Setter
// 目标企业ID
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetTargetRefEntId(_targetRefEntId string) error {
	r._targetRefEntId = _targetRefEntId
	r.Set("target_ref_ent_id", _targetRefEntId)
	return nil
}

// GetTargetRefEntId TargetRefEntId Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetTargetRefEntId() string {
	return r._targetRefEntId
}

// SetEquipmentCodeOrName is EquipmentCodeOrName Setter
// 设备编号或名称
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetEquipmentCodeOrName(_equipmentCodeOrName string) error {
	r._equipmentCodeOrName = _equipmentCodeOrName
	r.Set("equipment_code_or_name", _equipmentCodeOrName)
	return nil
}

// GetEquipmentCodeOrName EquipmentCodeOrName Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetEquipmentCodeOrName() string {
	return r._equipmentCodeOrName
}

// SetEquipmentType is EquipmentType Setter
// 设备类型
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetEquipmentType(_equipmentType string) error {
	r._equipmentType = _equipmentType
	r.Set("equipment_type", _equipmentType)
	return nil
}

// GetEquipmentType EquipmentType Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetEquipmentType() string {
	return r._equipmentType
}

// SetStatus is Status Setter
// 设备状态，1：正常；0：停止
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetStatus() string {
	return r._status
}

// SetPage is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 每页显示数量
func (r *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaAlihealthDrugKytDrVaequipmentListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytDrVaequipmentListRequest()
	},
}

// GetAlibabaAlihealthDrugKytDrVaequipmentListRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest
func GetAlibabaAlihealthDrugKytDrVaequipmentListAPIRequest() *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest {
	return poolAlibabaAlihealthDrugKytDrVaequipmentListAPIRequest.Get().(*AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytDrVaequipmentListAPIRequest 将 AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrVaequipmentListAPIRequest(v *AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrVaequipmentListAPIRequest.Put(v)
}
