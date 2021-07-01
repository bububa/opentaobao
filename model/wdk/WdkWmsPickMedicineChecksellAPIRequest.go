package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* WdkWmsPickMedicineChecksellAPIRequest
联营商药品柜核销 API请求
wdk.wms.pick.medicine.checksell

联营商药品柜核销 */
type WdkWmsPickMedicineChecksellAPIRequest struct {
	model.Params
	// 从二维码扫描出的信息
	_uuid string
	// shopId
	_shopId int64
}

// New
