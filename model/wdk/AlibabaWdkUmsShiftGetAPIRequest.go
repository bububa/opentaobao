package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsShiftGetAPIRequest
移库单获取 API请求
alibaba.wdk.ums.shift.get

移库单获取 */
type AlibabaWdkUmsShiftGetAPIRequest struct {
	model.Params
	// 店仓code，指的是库调对象，对应一个物理店或仓编码
	_warehouseCode string
}

// New
