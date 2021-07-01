package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsRetrieveConfirmAPIRequest
回流单－外部对已拉取到的UMS单据进行确认 API请求
alibaba.wdk.ums.retrieve.confirm

回流单－外部对已拉取到的UMS单据进行确认 */
type AlibabaWdkUmsRetrieveConfirmAPIRequest struct {
	model.Params
	// 店仓code，指的是作业对象，对应一个物理店或仓编码
	_warehouseCode string
	// 唯一识别码
	_uuid string
}

// New
