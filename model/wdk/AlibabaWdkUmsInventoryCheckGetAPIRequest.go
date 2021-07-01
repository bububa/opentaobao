package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsInventoryCheckGetAPIRequest
盘点结果单-回流单 API请求
alibaba.wdk.ums.inventory.check.get

盘点结果单-回流单 */
type AlibabaWdkUmsInventoryCheckGetAPIRequest struct {
	model.Params
	// 店仓code，指的是库调对象，对应一个物理店或仓编码
	_warehouseCode string
}

// New
