package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest
五道口悬挂链信息批量确认 API请求
taobao.wdk.equipment.conveyor.batchconfirm

批量消息确认 */
type TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest struct {
	model.Params
	// 仓库code
	_warehouseCode string
	// 待确认的uuid列表
	_uuids []string
}

// New
