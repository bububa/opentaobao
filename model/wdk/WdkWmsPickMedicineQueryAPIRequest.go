package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* WdkWmsPickMedicineQueryAPIRequest
查询拣货单中的药品信息 API请求
wdk.wms.pick.medicine.query

联营商药机查询拣货单中的药品信息 */
type WdkWmsPickMedicineQueryAPIRequest struct {
	model.Params
	// shopId
	_shopId int64
	// uuid
	_uuid string
}

// New
