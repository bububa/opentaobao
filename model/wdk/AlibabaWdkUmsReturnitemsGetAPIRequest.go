package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsReturnitemsGetAPIRequest
退货库位商品查询（退货出库辅助）-回流单 API请求
alibaba.wdk.ums.returnitems.get

退货库位商品查询（退货出库辅助）-回流单 */
type AlibabaWdkUmsReturnitemsGetAPIRequest struct {
	model.Params
	// 店仓code，指的是作业对象，对应一个物理店或仓编码
	_warehouseCode string
}

// New
