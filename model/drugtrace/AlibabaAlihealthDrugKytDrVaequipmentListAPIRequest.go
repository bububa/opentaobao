package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest
获取企业冷链设备信息 API请求
alibaba.alihealth.drug.kyt.dr.vaequipment.list

获取企业冷链设备信息 */
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

// New
