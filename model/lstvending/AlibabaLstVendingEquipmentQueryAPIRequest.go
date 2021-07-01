package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstVendingEquipmentQueryAPIRequest
自动售卖机设备信息查询 API请求
alibaba.lst.vending.equipment.query

为零售通品牌商提供已租赁的设备信息查询。 */
type AlibabaLstVendingEquipmentQueryAPIRequest struct {
	model.Params
	// 设备查询条件
	_openEquipmentQuery *OpenEquipmentQuery
}

// New
